package main

import (
	"encoding/json"
	"fmt"
	"onedb-core/engine/schema"
	"strconv"
	"time"
)

func main() {

	// configuration, err :=config.InstallConfig(3456,"root","root",filesys.GetRootDir())
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	startTime := time.Now().Local().UnixMilli()

	fmt.Println(strconv.Itoa(int(startTime)))
	// schemas := schema.Schema{
	// 	SchemaName: "Student",
	// 	Fields: []schema.Field{
	// 		schema.CreateField[int8]("test", 1),
	// 		{NAME: "ids", DATATYPE: reflect.String, DEFAULT_VALUE: "sa"},
	// 		{NAME: "asdfd", DATATYPE: reflect.Int32, NOT_NULL: true},
	// 	},
	// }
	// // error := schema.CreateSchema(schemas)
	// // if error != nil {
	// // 	fmt.Println(error)
	// // 	return
	// // }
	s, error := schema.ReadSchema("Student")
	readtime := time.Now().Local().UnixMilli()
	fmt.Println(strconv.Itoa(int(readtime)))
	if error != nil {
		fmt.Println(error)
		return
	}
	jsonData, error := json.MarshalIndent(s, "", "  ")
	if error != nil {
		fmt.Println(error)
		return
	}
	endtime := time.Now().Local().UnixMilli()
	fmt.Println(strconv.Itoa(int(endtime)))
	fmt.Println(string(jsonData))
	// error = proto.GenerateProtoFile(schema, configuration.DATABASE_STORAGE_ROOT+structure.PROTO_PATH)
	// endTime := time.Now().Local().UnixMilli()
	// fmt.Println(strconv.Itoa(int(endTime)))
	// fmt.Println(error)

}
