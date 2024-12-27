package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	jsonData := map[string]string{
		"query": `
		{
			countries{
			code,
			name,
			currency
			}
		}
		`,
	}

	jsonValues, err := json.Marshal(jsonData)

	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest("POST", "https://countries.trevorblades.com/", bytes.NewBuffer(jsonValues))
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

}
