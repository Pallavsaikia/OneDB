package main

import (
	"fmt"
	"onedb-core/config"
	"onedb-core/engine/proto"
	"onedb-core/engine/schema"
	"onedb-core/filesys"
	"onedb-core/structure"
	"reflect"
	"strconv"
	"time"
)

func main() {
	configuration, err := config.ReadConfig()
	if err != nil {
		return
	}
	if configuration.DATABASE_STORAGE_ROOT == "" {
		pathForRoot := filesys.GetRootDir()
		configuration.DATABASE_STORAGE_ROOT = pathForRoot
		_, err = config.WriteConfig(configuration)
		fmt.Println(err)
	}
	
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
	
	error = proto.GenerateProtoFile(schema, configuration.DATABASE_STORAGE_ROOT+structure.PROTO_PATH)
	endTime := time.Now().Local().UnixMilli()
	fmt.Println(strconv.Itoa(int(endTime)))
	fmt.Println(error)

}
