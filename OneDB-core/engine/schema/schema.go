package schema

import (
	"fmt"
	"onedb-core/config"
	"onedb-core/engine/datatype"
	"onedb-core/engine/schema/keys"
	"onedb-core/filesys"
	"onedb-core/libs"
	"onedb-core/structure"
	"os"
	"reflect"
	"strings"
)

type Schema struct {
	Fields           []Field `json:"fields"`
	SchemaName       string  `json:"schema_name"`
	RelativeLocation string  `json:"relative_location"`
	FolderName       string  `json:"folder_name"`
}

func (s *Schema) Encode() ([]byte, error) {
	return filesys.GobEncode(s)
}

// Decode deserializes the byte slice to populate the Person struct.
func (s *Schema) Decode(data []byte) error {
	return filesys.GobDecode(data, s)
}
func (schema *Schema) hasDuplicateFieldName() (bool, []string) {
	uniqueNames := make(map[string]struct{})
	duplicateFields := make([]string, 0)

	for _, field := range schema.Fields {
		// Convert field name to lowercase for case-insensitive comparison
		lowercaseName := strings.ToLower(field.NAME)

		// Check if the lowercase name already exists in the uniqueNames map
		if _, exists := uniqueNames[lowercaseName]; exists {
			// If the lowercase name exists, it's a duplicate, so add it to duplicateFields slice
			duplicateFields = append(duplicateFields, field.NAME)
		} else {
			// If the lowercase name doesn't exist, add it to the uniqueNames map
			uniqueNames[lowercaseName] = struct{}{}
		}
	}

	hasDuplicates := len(duplicateFields) > 0
	return hasDuplicates, duplicateFields
}

func (schema *Schema) Validate() error {
	if schema.SchemaName == "" {
		return fmt.Errorf("error:schema name cannot be empty")
	}
	if libs.ContainsNumber(schema.SchemaName) {
		return fmt.Errorf("error:schema name cannot contain number:'%s'", schema.SchemaName)
	}
	if libs.ContainsSpace(schema.SchemaName) {
		return fmt.Errorf("error:schema name cannot contain spaces:'%s'", schema.SchemaName)
	}
	if libs.ContainsSpecialCharacters(schema.SchemaName) {
		return fmt.Errorf("error:schema name cannot contain special characters other than a hyphen'_':'%s'", schema.SchemaName)
	}
	if len(schema.Fields) == 0 {
		return fmt.Errorf("error:schema requires atleast one field")
	}
	for i := 0; i < len(schema.Fields); i++ {
		error := schema.Fields[i].Validate(i)
		if error != nil {
			return error
		}
	}
	hasDuplicate, duplicateFieldString := schema.hasDuplicateFieldName()
	if hasDuplicate {
		return fmt.Errorf("error:schema cannot have duplicate field names: '%v'", duplicateFieldString)
	}
	return nil

}

func (schema *Schema) ValidateSizeOFFields() error {
	for i, field := range schema.Fields {
		if field.DATATYPE == reflect.String {
			if field.SIZE_IN_BYTE == 0 {
				return fmt.Errorf("error:string fields cannot be empty:'%s'", field.NAME)
			}
		} else {
			size, err := datatype.LoadKindSize(field.DATATYPE)
			if err != nil {
				return err
			}
			schema.Fields[i].SIZE_IN_BYTE = size
		}

	}
	return nil
}

func (schema *Schema) setUpPrimaryKeys() error {
	primaryIDs := []string{}
	//looping throught pkeys and checking if there is more than 1 key
	for i, field := range schema.Fields {
		if field.PKEY.KeyType != 0 {
			schema.Fields[i].DATATYPE = reflect.Int64
			schema.Fields[i].DEFAULT_VALUE = nil
			schema.Fields[i].NOT_NULL = true
			schema.Fields[i].UNIQUE = keys.UNIQUE_KEY{Unique: true}
			primaryIDs = append(primaryIDs, field.NAME)
		}
	}
	if len(primaryIDs) > 1 {
		return fmt.Errorf("error:There can be only one Primary key got %d:'%s'", len(primaryIDs), primaryIDs)
	}
	if len(primaryIDs) == 0 {
		schema.Fields = append([]Field{
			{NAME: "id", COLUMN_INDEX: 1, DATATYPE: reflect.Int64, NOT_NULL: true, PKEY: keys.PRIMARY_KEY{KeyType: keys.AutoIncreament}}},
			schema.Fields...)

	}
	return nil
}

func (schema *Schema) addTimeStamps() {
	schema.Fields = append(schema.Fields, CreateField[int64]("createdAt", true, 0, 0))
	schema.Fields = append(schema.Fields, CreateField[int64]("updatedAt", true, 0, 0))
}

func (s *Schema) createColumnIndex() {
	var idIndex int
	for i, field := range s.Fields {
		if field.NAME == "id" {
			idIndex = i
			break
		}
	}
	if idIndex != 0 {
		s.Fields[0], s.Fields[idIndex] = s.Fields[idIndex], s.Fields[0]
	}
	for i := range s.Fields {
		s.Fields[i].COLUMN_INDEX = i + 1
	}
}
func (schema *Schema) intitialize() error {
	schema.addTimeStamps()
	err := schema.setUpPrimaryKeys()
	if err != nil {
		return err
	}
	err = schema.ValidateSizeOFFields()
	if err != nil {
		return err
	}
	schema.createColumnIndex()
	err = schema.Validate()
	if err != nil {
		return err
	}
	return nil
}

func CreateSchema(schema Schema) error {
	error := schema.intitialize()
	if error != nil {
		return error
	}
	configuration, error := config.ReadConfig()
	if error != nil {
		return error
	}

	return filesys.WriteSchemaToFile(&schema, configuration.DATABASE_STORAGE_ROOT+structure.SCHEMA_PATH+"/"+schema.SchemaName+".bin")
}

func ReadSchema(schemaName string) (Schema, error) {
	schema := &Schema{}
	configuration, error := config.ReadConfig()
	if error != nil {
		return Schema{}, error
	}
	file_loc := configuration.DATABASE_STORAGE_ROOT + structure.SCHEMA_PATH + "/" + schemaName + ".bin"
	_, error = os.Stat(file_loc)
	if error != nil {
		return Schema{}, fmt.Errorf("error:schema with name '%s' doesnot exist", schemaName)
	}
	error = filesys.ReadSchemaToFile(schema, file_loc)
	if error != nil {
		return Schema{}, error
	}
	return *schema, nil
}
