package connection

import (
	"fmt"
)

type ConnectionResponse struct {
}

func EstablishConnection(username string, database string, port string, host string, password string) ConnectionResponse {
	fmt.Print("Trying to connect to " + host + ":" + port + " with user :" + username)
	connectionString := ConnectionResponse{}
	return connectionString
}
