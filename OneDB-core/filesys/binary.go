package filesys

import (
	"encoding/binary"
	"fmt"
	"os"
	"reflect"
)

func writeBinaryFile(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Slice {
		return fmt.Errorf("data must be a slice")
	}
	for i := 0; i < val.Len(); i++ {
		elem := val.Index(i)
		for j := 0; j < elem.NumField(); j++ {
			field := elem.Field(j)
			if err := binary.Write(file, binary.LittleEndian, field.Interface()); err != nil {
				return err
			}
		}
	}
	return nil
}

func readBinaryFile(filename string, data interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("data must be a pointer to a slice")
	}

	sliceType := val.Elem().Type().Elem()
	slice := reflect.MakeSlice(val.Elem().Type(), 0, 0)

	for {
		elem := reflect.New(sliceType).Elem()
		for j := 0; j < elem.NumField(); j++ {
			field := elem.Field(j)
			if err := binary.Read(file, binary.LittleEndian, field.Addr().Interface()); err != nil {
				break // end of file
			}
		}
		// Break out of the loop if we reached the end of the file
		if reflect.DeepEqual(elem.Interface(), reflect.Zero(sliceType).Interface()) {
			break
		}
		slice = reflect.Append(slice, elem)
	}

	val.Elem().Set(slice)
	return nil
}

// // Person struct defined later in the code
// type Person struct {
// 	ID   uint32
// 	Name string
// 	Age  uint8
// }

// func main() {
// 	dataToWrite := []Person{
// 		{1, "Alice", 25},
// 		{2, "Bob", 30},
// 		{3, "Charlie", 22},
// 	}

// 	filename := "people.dat"

// 	// Write data to binary file
// 	err := writeBinaryFile(filename, dataToWrite)
// 	if err != nil {
// 		fmt.Println("Error writing binary file:", err)
// 		return
// 	}
// 	fmt.Println("Data written to", filename)

// 	// Read data from binary file
// 	var dataToRead []Person
// 	err = readBinaryFile(filename, &dataToRead)
// 	if err != nil {
// 		fmt.Println("Error reading binary file:", err)
// 		return
// 	}

// 	// Display loaded data
// 	fmt.Println("Loaded data:")
// 	for _, person := range dataToRead {
// 		fmt.Printf("ID: %d, Name: %s, Age: %d\n", person.ID, person.Name, person.Age)
// 	}
// }
