package main

import (
	"fmt"
	"onedb-core/config"
	"onedb-core/engine/datatype"
	"onedb-core/engine/schema"
)

func main() {
	configuration, err := config.ReadConfig()
	if err != nil {
		return
	}
	fmt.Println(configuration)
	configs, _ := config.WriteConfig(config.Config{PORT: 3456, DEFAULT_USER: "root"})
	fmt.Println(configs)
	schema := schema.Schema{
		SchemaName: "Student",
		Fields: []schema.Field{
			{NAME: "asdf ", DATATYPE: datatype.Str},
			{NAME: "asdf", DATATYPE: datatype.Str},
		},
	}
	err = schema.Validate()
	fmt.Println(err)
	// fmt.Println(schema)
}
