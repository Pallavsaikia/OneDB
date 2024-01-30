package main

import (
	"fmt"
	"onedb-core/config"
)

func main() {
	config, err := config.ReadConfig()
	if err != nil {
		return
	}
	fmt.Print(config)
}
