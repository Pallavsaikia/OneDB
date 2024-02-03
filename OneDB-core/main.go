package main

import (
	"encoding/json"
	"fmt"
	// "onedb-core/config"
	// "onedb-core/engine/database"
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

	schema := schema.Schema{
		SchemaName: "Student",
		Fields: []schema.Field{
			{NAME: "ids", DATATYPE: reflect.String, DEFAULT_VALUE: "sa"},
			{NAME: "asdfd", DATATYPE: reflect.Int32},
			// {NAME: "asdfd", DATATYPE: reflect.Int16, PKEY: keys.PRIMARY_KEY{KeyType: keys.AutoIncreament}},
		},
	}
	error := schema.GenerateMeta()
	if error != nil {
		fmt.Println(error)
		return
	}
	jsonData, _ := json.MarshalIndent(schema, "", "    ")
	fmt.Println(string(jsonData))

}
