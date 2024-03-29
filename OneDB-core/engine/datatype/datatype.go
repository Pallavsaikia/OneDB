package datatype

import (
	"fmt"
	"reflect"
	"unsafe"
)

func SizeOfDataType(k reflect.Kind) int {

	switch k {
	case reflect.Int:
		var x int
		return int(unsafe.Sizeof(x))
	case reflect.Int8:
		var x int8
		return int(unsafe.Sizeof(x))
	case reflect.Int16:
		var x int16
		return int(unsafe.Sizeof(x))
	case reflect.Int32:
		var x int32
		return int(unsafe.Sizeof(x))
	case reflect.Int64:
		var x int64
		return int(unsafe.Sizeof(x))
	case reflect.Uint:
		var x uint
		return int(unsafe.Sizeof(x))
	case reflect.Uint8:
		var x uint8
		return int(unsafe.Sizeof(x))
	case reflect.Uint16:
		var x uint16
		return int(unsafe.Sizeof(x))
	case reflect.Uint32:
		var x uint32
		return int(unsafe.Sizeof(x))
	case reflect.Uint64:
		var x uint64
		return int(unsafe.Sizeof(x))
	case reflect.Float32:
		var x float32
		return int(unsafe.Sizeof(x))
	case reflect.Float64:
		var x float64
		return int(unsafe.Sizeof(x))
	case reflect.Complex64:
		var x complex64
		return int(unsafe.Sizeof(x))
	case reflect.Complex128:
		var x complex128
		return int(unsafe.Sizeof(x))
	case reflect.Bool:
		var x bool
		return int(unsafe.Sizeof(x))
	case reflect.Uintptr:
		var x uintptr
		return int(unsafe.Sizeof(x))
	case reflect.String:
		var x string
		return int(unsafe.Sizeof(x))
	default:
		return -1
	}
}

func ValidDataType(k reflect.Kind) bool {
	switch k {
	case reflect.Bool,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Float32,
		reflect.Float64,
		reflect.String:
		return true
	default:
		return false
	}
}


func LoadKindSize(k reflect.Kind) (int16, error) {
	switch k {
	case reflect.Bool:
		return 1, nil
	case reflect.Int8:
		return 1, nil
	case reflect.Int16:
		return 2, nil
	case reflect.Int32:
		return 4, nil
	case reflect.Int64:
		return 8, nil
	case reflect.Int:
		return 8, nil
	case reflect.Float32:
		return 4, nil
	case reflect.Float64:
		return 8, nil
	default:
		return 0, fmt.Errorf("error:unsupported kind")
	}
}