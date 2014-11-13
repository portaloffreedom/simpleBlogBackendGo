package main

import (
	"fmt"

	"github.com/portaloffreedom/simpleBlogBackendGo/database"
	"github.com/portaloffreedom/simpleBlogBackendGo/network"
)

func main() {
	database.Connect("localhost")
	fmt.Println("Hello, 世界")
	network.StartServer()

	database.Disconnect()
}
