package main

import (
	"ar/routes"
	"fmt"
)

func main() {

	fmt.Print("Starting server...\n")

	routes.SetupRouter().Run(":8080")
}
