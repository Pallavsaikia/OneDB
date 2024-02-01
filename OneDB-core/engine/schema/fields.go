package schema

import (
	"fmt"
	"onedb-core/engine/datatype"
)

type PrimaryKeyGenerateType int

const (
	TimeStamp PrimaryKeyGenerateType = iota + 1
	AutoIncreament
)

type PRIMARY_KEY struct {
	Name    string                 `json:"pkey_name"`
	KeyType PrimaryKeyGenerateType `json:"primary_key_gen_type"`
}
type FOREIGN_KEY struct {
	RelatedSchemaName string `json:"relation_schema_name"`
}
type UNIQUE_KEY struct {
	Unique bool `json:"unique"`
}
type Field struct {
	NAME     string            `json:"field_name"`
	DATATYPE datatype.DataType `json:"data_type"`
	UNIQUE   UNIQUE_KEY        `json:"ukey"`
	FKEY     FOREIGN_KEY       `json:"fkey"`
}

func (f Field) IsValid() (bool, error) {
	if f.NAME != "" {
		return false, fmt.Errorf("error:fieldname cannot be empty")
	}
	if !f.DATATYPE.IsValidDataType() {
		return false, fmt.Errorf("error:invalid datatype for %s", f.NAME)
	}
	return true, nil
}
