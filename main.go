package main

import (
	"fmt"

	"github.com/maitecr/go-rest-api/database"
	"github.com/maitecr/go-rest-api/models"
	"github.com/maitecr/go-rest-api/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{1, "Nome 1", "História 1"},
		{2, "Nome 2", "Historia 2"},
	}

	database.ConectarBD()

	fmt.Println("Iniciando o servidor rest")

	routes.HandleRequest()
}
