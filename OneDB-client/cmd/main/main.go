package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/Pallavsaikia/OneDb/OneDB-client/config"
	"github.com/Pallavsaikia/OneDb/OneDB-client/internal/connection"
)


func main() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	reader := bufio.NewReader(os.Stdin)
	username := flag.String("U", config.DEFAULT_USERNAME, "Username")
	database := flag.String("D", "", "Database name")
	port := flag.String("P", config.DEFAULT_PORT, "Port")
	host := flag.String("H", config.DEFAULT_HOST, "host")
	flag.Parse()
	fmt.Print("Password:")
	password, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Exiting the code...", password)
		return
	}
	connection.EstablishConnection(*username, *database, *port, *host,password)
	fmt.Print("\n\n#########################################\n")
	fmt.Print("Welcome to OneDb v1.0.0.\n")
	fmt.Print("#########################################\n")
	for {
		fmt.Print("OneDb>")
		userInput, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Println("You entered:", userInput)
	}
	fmt.Println("\nCtrl+C pressed. Exiting.")
}
