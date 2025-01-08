package conversion

import (
	"encoding/binary"
	"fmt"
	"onedb-core/engine/datatype"
	"reflect"
	"unsafe"
)

// SuperCastData casts the data to the specified kind and returns it
// SuperCastData casts the data to the specified kind and returns it
func SuperCastData(data interface{}, kind reflect.Kind) (interface{}, error) {
	// Check if the type is supported
	if !datatype.ValidDataType(reflect.TypeOf(data).Kind()){
		return nil,fmt.Errorf("data: %s of type '%s' unsupported",data,reflect.TypeOf(data).Kind())
	}

	// Use reflection to cast the data to the specified kind
	v := reflect.ValueOf(data)

	// Handle conversion based on target kind
	switch kind {
	case reflect.Int:
		switch v.Kind() {
		case reflect.Int:
			return v.Int(), nil
		case reflect.Int8:
			return int(v.Int()), nil
		case reflect.Int16:
			return int(v.Int()), nil
		case reflect.Int32:
			return int(v.Int()), nil
		case reflect.Int64:
			return int(v.Int()), nil
		}
	case reflect.Int8:
		if v.Kind() == reflect.Int {
			// Casting from int to int8, check for overflow
			val := v.Int()
			if val < -128 || val > 127 {
				return nil, fmt.Errorf("unsupported cast from %T to int8: value out of range", data)
			}
			return int8(val), nil
		}
		if v.Kind() == reflect.Int8 {
			return v.Int(), nil
		}
	case reflect.Int16:
		if v.Kind() == reflect.Int {
			// Casting from int to int16, check for overflow
			val := v.Int()
			if val < -32768 || val > 32767 {
				return nil, fmt.Errorf("unsupported cast from %T to int16: value out of range", data)
			}
			return int16(val), nil
		}
		if v.Kind() == reflect.Int16 {
			return v.Int(), nil
		}
	case reflect.Int32:
		if v.Kind() == reflect.Int {
			// Casting from int to int32, check for overflow
			val := v.Int()
			if val < -2147483648 || val > 2147483647 {
				return nil, fmt.Errorf("unsupported cast from %T to int32: value out of range", data)
			}
			return int32(val), nil
		}
		if v.Kind() == reflect.Int32 {
			return v.Int(), nil
		}
	case reflect.Int64:
		if v.Kind() == reflect.Int {
			// Casting from int to int64
			return int64(v.Int()), nil
		}
		if v.Kind() == reflect.Int64 {
			return v.Int(), nil
		}
	case reflect.Float32:
		if v.Kind() == reflect.Float32 {
			return v.Float(), nil
		} else if v.Kind() == reflect.Float64 {
			return float32(v.Float()), nil
		}
	case reflect.Float64:
		if v.Kind() == reflect.Float64 {
			return v.Float(), nil
		} else if v.Kind() == reflect.Float32 {
			return float64(v.Float()), nil
		}
	case reflect.String:
		if v.Kind() == reflect.String {
			return v.String(), nil
		}
	case reflect.Bool:
		if v.Kind() == reflect.Bool {
			return v.Bool(), nil
		}

	}

	return nil, fmt.Errorf("unsupported cast from %T to %s", data, kind)
}
// SupportedDataTypes checks if the data type is supported


// intToBytes is a helper to convert integers of various sizes to byte slices
func intToBytes(value int, size int) []byte {
	buf := make([]byte, size)
	binary.BigEndian.PutUint64(buf[len(buf)-8:], uint64(value)) // Support for big-endian format
	return buf[len(buf)-size:]                                 // Return the relevant portion
}

// floatToBytes converts float32 and float64 to byte slices
func floatToBytes(value float64, size int) []byte {
	buf := make([]byte, size)
	if size == 4 { // float32
		bits := *(*uint32)(unsafe.Pointer(&value))
		binary.BigEndian.PutUint32(buf, bits)
	} else if size == 8 { // float64
		bits := *(*uint64)(unsafe.Pointer(&value))
		binary.BigEndian.PutUint64(buf, bits)
	}
	return buf
}

// InterfaceToBytes converts supported types to byte slices
func InterfaceToBytes(data interface{}) ([]byte, error) {
	// Check if the type is supported
	if !datatype.ValidDataType(reflect.TypeOf(data).Kind()){
		return nil,fmt.Errorf("data: %s of type '%s' unsupported",data,reflect.TypeOf(data).Kind())
	}

	// Convert the data to bytes based on its type
	switch v := data.(type) {
	case int:
		// Convert int to bytes directly
		return intToBytes(v, 8), nil
	case int8:
		// Convert int8 to bytes
		return []byte{byte(v)}, nil
	case int16:
		// Convert int16 to bytes
		return intToBytes(int(v), 2), nil
	case int32:
		// Convert int32 to bytes
		return intToBytes(int(v), 4), nil
	case int64:
		// Convert int64 to bytes
		return intToBytes(int(v), 8), nil
	case float32:
		// Convert float32 to bytes
		return floatToBytes(float64(v), 4), nil
	case float64:
		// Convert float64 to bytes
		return floatToBytes(v, 8), nil
	case bool:
		// Convert bool to bytes
		if v {
			return []byte{1}, nil
		} else {
			return []byte{0}, nil
		}
	case string:
		// Convert string to bytes directly
		return []byte(v), nil

	default:
		// Fallback case, should not be reached
		return nil, fmt.Errorf("unsupported type: %T", data)
	}
}


