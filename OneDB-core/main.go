package main

import (
	// "encoding/json"
	"fmt"
	"strconv"
	"time"

	// "onedb-core/config"
	// "onedb-core/engine/database"
	"onedb-core/engine/proto"
	"onedb-core/engine/schema"

	// "onedb-core/engine/schema/keys"
	"reflect"
)

func main() {
	// configuration, err := config.ReadConfig()
	// if err != nil {
	// 	return
	// }
	// fmt.Println(configuration)
	// configs, _ := config.WriteConfig(config.Config{PORT: 3456, DEFAULT_USER: "root"})
	// fmt.Println(configs)
	startTime := time.Now().Local().UnixMilli()
	
	fmt.Println(strconv.Itoa(int(startTime)))
	schema := schema.Schema{
		SchemaName: "Student",
		Fields: []schema.Field{
			schema.CreateField[int8]("test", 1),
			{NAME: "ids", DATATYPE: reflect.String, DEFAULT_VALUE: "sa"},
			{NAME: "asdfd", DATATYPE: reflect.Int32, NOT_NULL: true},
			// {NAME: "asdfd", DATATYPE: reflect.Int16, PKEY: keys.PRIMARY_KEY{KeyType: keys.AutoIncreament}},
		},
	}
	error := schema.GenerateMeta()
	if error != nil {
		fmt.Println(error)
		return
	}
	error = proto.GenerateProtoFile(schema, "")
	endTime := time.Now().Local().UnixMilli()
	fmt.Println(strconv.Itoa(int(endTime)))
	// duration := endTime.Sub(startTime).Milliseconds()
	// fmt.Printf("Time taken: %v,%v,%v\n", startTime,endTime,duration)
	fmt.Println(error)
	// jsonData, _ := json.MarshalIndent(schema, "", "    ")
	// fmt.Println(string(jsonData))

}
