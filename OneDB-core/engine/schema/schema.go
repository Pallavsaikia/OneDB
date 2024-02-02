package schema

import (
	"fmt"
)

type Schema struct {
	Fields             []Field     `json:"fields"`
	SchemaName         string      `json:"schema_name"`
	RelativeLocation   string      `json:"relative_location"`
	PrimaryKeyMetaData PRIMARY_KEY `json:"primary_key_meta"`
	FolderName         string      `json:"folder_name"`
}

func (schema *Schema) Validate() error {
	if schema.SchemaName == "" {
		return fmt.Errorf("schema name cannot be empty")
	}
	if len(schema.Fields) == 0 {
		return fmt.Errorf("schema requires atleast one field")
	}

	for i := 0; i < len(schema.Fields); i++ {
		error := schema.Fields[i].Validate(i)
		if error != nil {
			return error
		}
	}
	return nil
}
