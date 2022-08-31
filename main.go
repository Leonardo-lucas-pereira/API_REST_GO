package main

import (
	"fmt"
	"go-rest-api/database"
	"go-rest-api/routes"
)

func main() {
	database.ConectaComBancoDeDados()

	fmt.Println("Iniciando o serviodor Rest com Go")
	routes.HandleResquets()
}
