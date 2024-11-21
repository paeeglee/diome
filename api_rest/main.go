package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Client struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var clients []Client = make([]Client, 0)

func HandleFetchClient(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(clients)
}

func HandleCreateClient(w http.ResponseWriter, r *http.Request) {
	var client Client
	_ = json.NewDecoder(r.Body).Decode(&client)

	client.ID = len(clients) + 1
	clients = append(clients, client)

	json.NewEncoder(w).Encode(client)
}

func HandleFetchByIdClient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	idx := -1

	for i, c := range clients {
		if c.ID == id {
			idx = i
		}
	}

	if idx >= 0 {
		json.NewEncoder(w).Encode(clients[idx])
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func HandleUpdateClient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	idx := -1

	for i, c := range clients {
		if c.ID == id {
			idx = i
		}
	}

	if idx >= 0 {
		var client Client
		_ = json.NewDecoder(r.Body).Decode(&client)

		clients[idx].Name = client.Name

		json.NewEncoder(w).Encode(clients[idx])
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func HandleDeleteClient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	idx := -1

	for i, c := range clients {
		if c.ID == id {
			idx = i
		}
	}

	if idx >= 0 {
		clients = append(clients[:idx], clients[idx+1:]...)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/clientes", HandleFetchClient).Methods("GET")
	r.HandleFunc("/clientes/{id}", HandleFetchByIdClient).Methods("GET")
	r.HandleFunc("/clientes", HandleCreateClient).Methods("POST")
	r.HandleFunc("/clientes/{id}", HandleUpdateClient).Methods("PUT")
	r.HandleFunc("/clientes/{id}", HandleDeleteClient).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
