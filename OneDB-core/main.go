package main

import (
	"fmt"
	"onedb-core/config"
	"onedb-core/engine/schema"
	"reflect"
)

func main() {
	configuration, err := config.ReadConfig()
	if err != nil {
		return
	}
	fmt.Println(configuration)
	configs, _ := config.WriteConfig(config.Config{PORT: 3456, DEFAULT_USER: "root"})
	fmt.Println(configs)
	// var a int16 = 2
	schema := schema.Schema{
		SchemaName: "Student",
		Fields: []schema.Field{
			{NAME: "asdf", DATATYPE: reflect.Int16, DEFAULT_VALUE: "sa"},
		},
	}
	err = schema.Validate()
	fmt.Println(err)
	// fmt.Println(schema.Fields[0].FIELD_SIZE_IN_BYTE)
}
