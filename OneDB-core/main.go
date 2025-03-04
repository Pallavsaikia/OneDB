package main

import (
	"encoding/json"
	"fmt"

	// "onedb-core/config"
	"onedb-core/engine/cache"
	"onedb-core/engine/dataset"
	"onedb-core/engine/schema"

	// "onedb-core/filesys"
	"reflect"
	"strconv"
	"time"
)

func main() {

	c := cache.NewCache(20)
	// _, err :=config.InstallConfig(3456,"root","root",filesys.GetRootDir(),c)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	startTime := time.Now().Local().UnixMilli()

	fmt.Println(strconv.Itoa(int(startTime)))
	schemas := schema.Schema{
		SchemaName: "Student",
		Fields: []schema.Field{
			schema.CreateField[int8]("test", false, 1, 0),
			{NAME: "ids", DATATYPE: reflect.String, DEFAULT_VALUE: "sa", SIZE_IN_BYTE: 24},
		},
	}
	error := schema.CreateSchema(schemas, c)
	if error != nil {
		fmt.Println(error)
	}

	s, error := schema.ReadSchema("Student", c)
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

	error = dataset.Insert(&s, map[string]interface{}{"id": 1, "test": 1, "ids": "s2"})
	fmt.Println(error)
	// error = proto.GenerateProtoFile(schema, configuration.DATABASE_STORAGE_ROOT+structure.PROTO_PATH)
	// endTime := time.Now().Local().UnixMilli()
	// fmt.Println(strconv.Itoa(int(endTime)))
	// fmt.Println(error)

}
