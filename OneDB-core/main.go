package main

import (
	"fmt"
	"onedb-core/config"
)

func main() {
	configuration, err := config.ReadConfig()
	if err != nil {
		return
	}
	fmt.Print(configuration)
	configs, _ := config.WriteConfig(config.Config{PORT: 3456,DEFAULT_USER: "root"})
	fmt.Print(configs)
	
}
