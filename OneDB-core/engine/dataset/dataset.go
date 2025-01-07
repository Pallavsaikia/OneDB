package dataset

import (
	"encoding/binary"
	"fmt"
	"onedb-core/engine/schema"
	"os"
	"reflect"
)

/**
*data is array of any type data
*data a hasmap of field schema.Field.Name  and value is the value of the field
*loop through schema.Field[] and extract the value from data and insert into file
*file name is schema.SchemaName
*file location is schema.DataFileLocation
*file format is bin
*check the size of value if its more than schema.Field[i].Size
*add next  column value after the size of the field
*new line for new row
 */

func GetDataFilePath(schema *schema.Schema) string {
	return schema.DataFileLocation
}
func validateData(schema *schema.Schema, data map[string]interface{}) error {
	// Implement validation logic based on the schema
	// For example, check data types, field presence, etc.
	return nil
}

// serializeData serializes data into a byte slice based on the schema
func serializeData(schema *schema.Schema, data map[string]interface{}) ([]byte, error) {
	// Initialize byte slice for serialized data
	serializedData := make([]byte, 0)

	// Iterate over fields in the schema
	for _, field := range schema.Fields {
		// Get value from data map based on field name
		value, ok := data[field.NAME]
		if !ok {
			return nil, fmt.Errorf("field '%s' not found in data", field.NAME)
		}

		// Serialize value based on data type
		switch field.DATATYPE {
		case reflect.String:
			strValue, ok := value.(string)
			if !ok {
				return nil, fmt.Errorf("expected string for field '%s'", field.NAME)
			}
			serializedData = append(serializedData, []byte(strValue)...)
		case reflect.Int:
			intValue, ok := value.(int)
			if !ok {
				return nil, fmt.Errorf("expected int for field '%s'", field.NAME)
			}
			// Convert int to bytes (assuming int is 4 bytes)
			intBytes := make([]byte, 4)
			binary.BigEndian.PutUint32(intBytes, uint32(intValue))
			serializedData = append(serializedData, intBytes...)
		// Add cases for other data types as needed
		default:
			return nil, fmt.Errorf("unsupported data type for field '%s'", field.NAME)
		}
	}

	return serializedData, nil
}

func Insert(schema *schema.Schema, data map[string]interface{}) error {
	// loop through schema.Field[] and extract the value from data and insert into file
	file, err := os.OpenFile(GetDataFilePath(schema), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, field := range schema.Fields {
		value, ok := data[field.NAME]
		if !ok && field.NOT_NULL {
			return fmt.Errorf("error:field '%s' is required", field.NAME)
		}
		if reflect.TypeOf(value).Kind() != field.DATATYPE {
			return fmt.Errorf("error:field '%s' is of type '%s' but given '%s'", field.NAME,field.DATATYPE,reflect.TypeOf(value).Kind())
		}
		fmt.Print(value)
	}
	// new line for new row
	_, err = file.WriteString("\n")
	return err
}
