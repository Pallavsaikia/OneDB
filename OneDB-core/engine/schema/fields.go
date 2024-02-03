package schema

import (
	"fmt"
	"onedb-core/engine/datatype"
	"onedb-core/libs"
	"reflect"
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
	NAME          string       `json:"field_name"`
	DATATYPE      reflect.Kind `json:"data_type"`
	UNIQUE        UNIQUE_KEY   `json:"ukey"`
	FKEY          FOREIGN_KEY  `json:"fkey"`
	NOT_NULL      bool         `json:"notnull"`
	DEFAULT_VALUE any          `json:"default_val"`
	COLUMN_INDEX  int          `json:"column_index"`
	// FIELD_SIZE_IN_BYTE int          `json:"field_size"` //in bytes
}

func (field *Field) Validate(i int) error {
	// field.FIELD_SIZE_IN_BYTE = libs.SizeOfKind(field.DATATYPE)
	if field.NAME == "" {
		return fmt.Errorf("error:fieldname cannot be empty:field index:%d", i)
	}
	if !datatype.ValidDataType(field.DATATYPE) {
		return fmt.Errorf("error:datatpe for '%s' invalid:%s ", field.NAME, field.DATATYPE.String())
	}
	if libs.ContainsSpace(field.NAME) {
		return fmt.Errorf("error:fieldname cannot have spaces:'%s'", field.NAME)
	}
	if libs.ContainsNumber(field.NAME) {
		return fmt.Errorf("error:fieldname cannot have number:'%s'", field.NAME)
	}
	if libs.ContainsSpecialCharacters(field.NAME) {
		return fmt.Errorf("error:fieldname cannot contain special characters other than a hyphen'_':'%s'", field.NAME)
	}
	switch field.DATATYPE {
	case reflect.Array, reflect.Struct, reflect.Map, reflect.Slice:
		return fmt.Errorf("error:datatype cannot be an 'array', 'struct', 'map' or 'slice'")
	}
	if field.DATATYPE == reflect.Invalid {
		return fmt.Errorf("error:choose a valid datatype")
	}
	if field.DEFAULT_VALUE != nil {
		fmt.Print(reflect.TypeOf(field.DEFAULT_VALUE).String())
	}
	return nil
}
