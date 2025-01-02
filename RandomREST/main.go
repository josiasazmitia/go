package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	UUID string `json:"uuid"`
	Name struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"name"`
}

type Response struct {
	Results []struct {
		Login User `json:"login"`
	} `json:"results"`
}

func fetchUsers(endpoint string) ([]string, error) {
	fmt.Println("Solicitar usuarios a:", endpoint)

	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("Error al realizar la solicitud: %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error al leer la respuesta: %v", err)
	}

	fmt.Println("Respuesta:", string(body))

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("Error al parsear la respuesta JSON: %v", err)
	}

	return nil, nil
}

func main() {
	const endpoint = "https://randomuser.me/api/?results=10"
	const target = 15000

	users, err := fetchUsers(endpoint)

	if err != nil {
		fmt.Println("Error al obtener usuarios:", err)
	}

	fmt.Println("Usuarios obtenidos:", len(users))

}
