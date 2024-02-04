package schema

import (
	"fmt"
	"onedb-core/config"
	"onedb-core/engine/schema/keys"
	"onedb-core/filesys"
	"onedb-core/libs"
	"onedb-core/structure"
	"os"
	"reflect"
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
	duplicateFields := []string{}
	for _, field := range schema.Fields {
		if _, exists := uniqueNames[field.NAME]; exists {
			duplicateFields = append(duplicateFields, field.NAME)
		} else {
			uniqueNames[field.NAME] = struct{}{}
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

func (schema *Schema) Intitialize() error {
	err := schema.Validate()
	if err != nil {
		return fmt.Errorf("error:couldnot validate schema\n%s", err)
	}
	primaryIDs := []string{}
	for i, field := range schema.Fields {
		//adding indexes
		if field.PKEY.KeyType != 0 {
			schema.Fields[i].DATATYPE = reflect.Int64
			schema.Fields[i].DEFAULT_VALUE = nil
			primaryIDs = append(primaryIDs, field.NAME)
		}
		schema.Fields[i].COLUMN_INDEX = i
	}
	if len(primaryIDs) > 1 {
		return fmt.Errorf("error:There can be only one Primary key got %d:'%s'", len(primaryIDs), primaryIDs)
	}
	if len(primaryIDs) == 0 {
		schema.Fields = append([]Field{
			{NAME: "Id", COLUMN_INDEX: 1, DATATYPE: reflect.Int64, PKEY: keys.PRIMARY_KEY{KeyType: keys.AutoIncreament}}},
			schema.Fields...)
		for i := range schema.Fields {
			schema.Fields[i].COLUMN_INDEX = i
		}
	}
	return nil
}

func CreateSchema(schema Schema) error {
	error := schema.Intitialize()
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
