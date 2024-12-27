package main

import (
	"fmt"
	"log"

	"github.com/josiasazmitia/go/GraphQL/service"
)

func main() {
	fmt.Println("Ejemplo de GraphQL en Go")

	countries, err := service.FetchCountries()
	if err != nil {
		log.Fatalf("Error al obtener países: %v", err)
	}

	for _, country := range countries {
		fmt.Printf("Código: %s, Nombre: %s, Moneda: %s\n", country.Code, country.Name, country.Currency)
	}
}
