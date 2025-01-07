package conversion

import (
	"fmt"
	"unsafe"
)

func InterfaceToBytes(data interface{}) ([]byte, error) {
	switch v := data.(type) {
	case int:
		// Convert int to bytes directly using unsafe pointer
		return intToBytes(int(v), 8), nil
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
		bits := *(*uint32)(unsafe.Pointer(&v))
		return intToBytes(int(bits), 4), nil
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
		return nil, fmt.Errorf("unsupported type: %T", data)
	}
}

func intToBytes(num int, size int) []byte {
	bytes := make([]byte, size)
	for i := 0; i < size; i++ {
		bytes[i] = byte(num >> uint(i*8) & 0xFF)
	}
	return bytes
}

