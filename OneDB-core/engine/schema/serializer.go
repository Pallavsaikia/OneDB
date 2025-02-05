package schema

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"onedb-core/engine/conversion"
	"reflect"
	"sort"
)

func (s *Schema) Serialize(data map[string]interface{}) ([]byte, error) {
	var buf bytes.Buffer

	// Sort fields by COLUMN_INDEX
	sort.Slice(s.Fields, func(i, j int) bool {
		return s.Fields[i].COLUMN_INDEX < s.Fields[j].COLUMN_INDEX
	})

	for _, field := range s.Fields {
		val, exists := data[field.NAME]

		// Handle missing values
		if !exists {
			if field.NOT_NULL {
				return nil, fmt.Errorf("missing required field: %s", field.NAME)
			}
			val = field.DEFAULT_VALUE
		}

		// Write data based on SIZE_IN_BYTE
		start := buf.Len()
		switch v := val.(type) {
		case int:
			if err := binary.Write(&buf, binary.LittleEndian, int32(v)); err != nil {
				return nil, err
			}
		case int8:
			if err := binary.Write(&buf, binary.LittleEndian, v); err != nil {
				return nil, err
			}
		case int16:
			if err := binary.Write(&buf, binary.LittleEndian, v); err != nil {
				return nil, err
			}
		case int32:
			if err := binary.Write(&buf, binary.LittleEndian, v); err != nil {
				return nil, err
			}
		case int64:
			if err := binary.Write(&buf, binary.LittleEndian, v); err != nil {
				return nil, err
			}
		case float32:
			if err := binary.Write(&buf, binary.LittleEndian, v); err != nil {
				return nil, err
			}
		case float64:
			if err := binary.Write(&buf, binary.LittleEndian, v); err != nil {
				return nil, err
			}
		case string:
			strBytes := []byte(v)
			if err := binary.Write(&buf, binary.LittleEndian, int32(len(strBytes))); err != nil {
				return nil, err
			}
			if err := binary.Write(&buf, binary.LittleEndian, strBytes); err != nil {
				return nil, err
			}
		case bool:
			if err := binary.Write(&buf, binary.LittleEndian, v); err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("unsupported data type for field %s", field.NAME)
		}

		// Ensure padding based on SIZE_IN_BYTE
		size := field.SIZE_IN_BYTE
		padding := int(size) - (buf.Len() - start)
		if padding > 0 {
			if err := binary.Write(&buf, binary.LittleEndian, make([]byte, padding)); err != nil {
				return nil, err
			}
		}
	}

	return buf.Bytes(), nil
}


func (s *Schema) Deserialize(data []byte) (map[string]interface{}, error) {
	deserializedData := make(map[string]interface{})
	offset := 0

	for _, field := range s.Fields {
		// Determine the number of bytes the field occupies
		fieldSize := field.SIZE_IN_BYTE
		if fieldSize == 0 {
			return nil, fmt.Errorf("invalid field size for field: %s", field.NAME)
		}

		// Extract the field data from the byte array
		fieldData := data[offset : offset+int(fieldSize)]
		offset += int(fieldSize)

		// Deserialize based on the type of the field
		var value interface{}
		var err error

		// Handle type casting according to the field's data type
		switch field.DATATYPE {
		case reflect.Int:
			value, err = conversion.SuperCastData(int64(binary.LittleEndian.Uint32(fieldData)), field.DATATYPE)
		case reflect.Int8:
			value, err =  conversion.SuperCastData(int8(fieldData[0]), field.DATATYPE)
		case reflect.Int16:
			value, err =  conversion.SuperCastData(int16(binary.LittleEndian.Uint16(fieldData)), field.DATATYPE)
		case reflect.Int32:
			value, err =  conversion.SuperCastData(int32(binary.LittleEndian.Uint32(fieldData)), field.DATATYPE)
		case reflect.Int64:
			value, err =  conversion.SuperCastData(int64(binary.LittleEndian.Uint64(fieldData)), field.DATATYPE)
		case reflect.Float32:
			value, err =  conversion.SuperCastData(math.Float32frombits(binary.LittleEndian.Uint32(fieldData)), field.DATATYPE)
		case reflect.Float64:
			value, err =  conversion.SuperCastData(math.Float64frombits(binary.LittleEndian.Uint64(fieldData)), field.DATATYPE)
		case reflect.String:
			value = string(fieldData)
		case reflect.Bool:
			value = fieldData[0] != 0
		default:
			err = fmt.Errorf("unsupported data type for field: %s", field.NAME)
		}

		if err != nil {
			return nil, fmt.Errorf("error deserializing field '%s': %v", field.NAME, err)
		}

		// Add the value to the deserialized data map
		deserializedData[field.NAME] = value
	}

	// Return the deserialized data
	return deserializedData, nil
}
