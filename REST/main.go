package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	const endpoint string = "https://randomuser.me/api/?results=5"

	request, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("error", err)
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Println("error", err)
	}

	sb := string(body)
	fmt.Println(sb)
}
