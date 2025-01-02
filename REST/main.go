package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Results []struct {
		Name struct {
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
	} `json:"results"`
}

func main() {
	const endpoint string = "https://randomuser.me/api/?results=10"

	// Realizar la solicitud GET
	request, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Error al realizar la solicitud:", err)
		return
	}
	defer request.Body.Close()

	// Leer la respuesta
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta:", err)
		return
	}

	// Parsear la respuesta JSON
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error al parsear la respuesta JSON:", err)
		return
	}

	// Utilizar un mapa para evitar duplicados
	uniqueNames := make(map[string]bool)

	fmt.Println("Nombres y apellidos únicos:")
	for _, result := range response.Results {
		fullName := result.Name.First + " " + result.Name.Last
		if !uniqueNames[fullName] {
			uniqueNames[fullName] = true
			fmt.Println(fullName)
		}
	}
}
