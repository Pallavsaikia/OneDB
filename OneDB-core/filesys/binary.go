package filesys

import (
	"encoding/binary"
	"fmt"
	"os"
	"reflect"
)

// EncodeStructToBinary encodes a struct to binary format.
func EncodeStructToBinary(data interface{}, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	value := reflect.ValueOf(data)

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)

		switch field.Kind() {
		case reflect.Int:
			err = binary.Write(file, binary.LittleEndian, field.Int())
		case reflect.Int8:
			err = binary.Write(file, binary.LittleEndian, int8(field.Int()))
		case reflect.Int16:
			err = binary.Write(file, binary.LittleEndian, int16(field.Int()))
		case reflect.Int32:
			err = binary.Write(file, binary.LittleEndian, int32(field.Int()))
		case reflect.Int64:
			err = binary.Write(file, binary.LittleEndian, field.Int())
		case reflect.Uint:
			err = binary.Write(file, binary.LittleEndian, field.Uint())
		case reflect.Uint8:
			err = binary.Write(file, binary.LittleEndian, uint8(field.Uint()))
		case reflect.Uint16:
			err = binary.Write(file, binary.LittleEndian, uint16(field.Uint()))
		case reflect.Uint32:
			err = binary.Write(file, binary.LittleEndian, uint32(field.Uint()))
		case reflect.Uint64:
			err = binary.Write(file, binary.LittleEndian, field.Uint())
		case reflect.Float32:
			err = binary.Write(file, binary.LittleEndian, float32(field.Float()))
		case reflect.Float64:
			err = binary.Write(file, binary.LittleEndian, field.Float())
		case reflect.String:
			err = binary.Write(file, binary.LittleEndian, []byte(field.String()))
		case reflect.Bool:
			err = binary.Write(file, binary.LittleEndian, field.Bool())
		default:
			return fmt.Errorf("unsupported field type: %v", field.Kind())
		}

		if err != nil {
			return err
		}
	}

	return nil
}

// DecodeBinaryToStruct decodes a binary file to a struct.
func DecodeBinaryToStruct(filePath string, data interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	value := reflect.ValueOf(data)

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)

		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			var fieldValue int64
			err = binary.Read(file, binary.LittleEndian, &fieldValue)
			field.SetInt(fieldValue)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			var fieldValue uint64
			err = binary.Read(file, binary.LittleEndian, &fieldValue)
			field.SetUint(fieldValue)
		case reflect.Float32, reflect.Float64:
			var fieldValue float64
			err = binary.Read(file, binary.LittleEndian, &fieldValue)
			field.SetFloat(fieldValue)
		// Add other cases for different types if needed
		default:
			return fmt.Errorf("unsupported field type: %v", field.Kind())
		}

		if err != nil {
			return err
		}
	}

	return nil
}
