package main

import (
	"fmt"
	"go-rest-api/routes"
)

func main() {
	fmt.Println("Iniciando o serviodor Rest com Go")
	routes.HandleResquets()
}
