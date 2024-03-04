package dataset

import (
	// "fmt"
	"onedb-core/engine/schema"
	// "os"
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

func Insert(schema *schema.Schema, data map[string]interface{}) error {
	// loop through schema.Field[] and extract the value from data and insert into file
	// file, err := os.OpenFile(GetDataFilePath(schema), os.O_APPEND|os.O_WRONLY, 0644)
	// if err != nil {
	// 	return err
	// }
	// defer file.Close()
	// for _, field := range schema.Fields {
	// 	value, ok := data[field.NAME]
	// 	if !ok && field.NOT_NULL {
	// 		return fmt.Errorf("error:field '%s' is required", field.NAME)
	// 	}
	// 	if value != nil {
	// 		// check the size of value if its more than schema.Field[i].Size
	// 		if len(value.(string)) > int(field.SIZE_IN_BYTE) {
	// 			return fmt.Errorf("error:field '%s' is too long", field.NAME)
	// 		}
	// 		_, err = file.WriteString(value.(string))
	// 		if err != nil {
	// 			return err
	// 		}
	// 	} else {
	// 		//leave gap of size
	// 		_, err = file.WriteString(fmt.Sprintf("%*s", int(field.SIZE_IN_BYTE), ""))
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}
	// }
	// // new line for new row
	// _, err = file.WriteString("\n")
	// return err
	return nil
}
