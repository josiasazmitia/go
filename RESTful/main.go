package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type TODO struct {
	ID        string `json:"Id"`
	Title     string `json:"Title"`
	Completed bool   `json:"Completed"`
}

var todos []TODO

func main() {
	http.HandleFunc("/todos", getTodos)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	todos = append(todos, TODO{ID: "1", Title: "Comprar pan", Completed: false})
	json.NewEncoder(w).Encode(todos)
}
