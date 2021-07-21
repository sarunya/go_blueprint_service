package main

import (
	"bp_service/lib/base"
	"bp_service/src/routes"
	"fmt"
	// "net/http"
	// "github.com/lib/pq"
)

func main() {
	dependencies := base.BootUp()
	routes := routes.GetRoutes()

	base.StartServer(*dependencies, routes)

	fmt.Println("Im going to start a server here!")
}
