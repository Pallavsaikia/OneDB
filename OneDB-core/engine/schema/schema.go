package schema

import (
	"fmt"
	"onedb-core/libs"
)

type Schema struct {
	Fields             []Field     `json:"fields"`
	SchemaName         string      `json:"schema_name"`
	RelativeLocation   string      `json:"relative_location"`
	PrimaryKeyMetaData PRIMARY_KEY `json:"primary_key_meta"`
	FolderName         string      `json:"folder_name"`
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
		return fmt.Errorf("error:schema cannot contain number:'%s'", schema.SchemaName)
	}
	if libs.ContainsSpace(schema.SchemaName) {
		return fmt.Errorf("error:schema cannot contain spaces:'%s'", schema.SchemaName)
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

// func (schema *Schema) Cleanup() Schema {
// 	if schema.Validate() == nil {

// 	}
// }
