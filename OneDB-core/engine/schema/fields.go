package schema

import (
	"fmt"
	"onedb-core/engine/datatype"
	"onedb-core/engine/schema/keys"
	"onedb-core/libs"
	"reflect"
	"strings"
)

// Field is a column in the schema.
//
// @param NAME(string): name of the field
//
// @param DATATYPE(reflect.Kind): data type for the column/field.It is same as the golangs inbuilt datatype
//
// @param UNIQUE(bool): is the field unique
//
// @param FKEY(FOREIGN_KEY): foreign key details
//
// @param NOT_NULL(bool): is field allowed t	`o be null
//
// @param DEFAULT_VALUE(any): default value if nothing is given
//
// @param COLUMN_INDEX(int): position of the column in the table
type Field struct {
	NAME          string           `json:"field_name"`
	DATATYPE      reflect.Kind     `json:"data_type"`
	UNIQUE        keys.UNIQUE_KEY  `json:"ukey"`
	FKEY          keys.FOREIGN_KEY `json:"fkey"`
	PKEY          keys.PRIMARY_KEY `json:"pkey"`
	NOT_NULL      bool             `json:"notnull"`
	DEFAULT_VALUE any              `json:"default_val"`
	COLUMN_INDEX  int              `json:"column_index"`
	// FIELD_SIZE_IN_BYTE int          `json:"field_size"` //in bytes
}

func (field *Field) Validate(i int) error {
	// field.FIELD_SIZE_IN_BYTE = libs.SizeOfKind(field.DATATYPE)
	
	if field.NAME == "" {
		return fmt.Errorf("error:fieldname cannot be empty:field index:%d", i)
	}
	if strings.EqualFold(field.NAME, "id") && field.PKEY.KeyType==0{
		return fmt.Errorf("error:fieldname cannot be named 'id' if its not a primary key")
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
		dType := reflect.TypeOf(field.DEFAULT_VALUE)
		if dType.Kind() != field.DATATYPE {
			return fmt.Errorf("error:default value for field '%s' should be of type:'%s' but was found '%s'", field.NAME, field.DATATYPE.String(), dType.Kind().String())
		}

	}
	return nil
}
