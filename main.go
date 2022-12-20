package main

import (
	"context"
	"fmt"

	"microtwo/routes"
	"microtwo/servers"
	"microtwo/utils"
)

func init() {
	fmt.Println("Initializing...")
	utils.LoadEnvs()
}

func main() {

	// Create a new server
	httpServer, err := servers.NewHttpServer(context.Background())

	if err != nil {
		panic(err.Error())
	}

	// Close the server connection
	defer httpServer.Close()

	// Start the server
	httpServer.Start(routes.SetRoutes)
}
