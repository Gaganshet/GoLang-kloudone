package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Item struct {
	UID   string  `json:"UID"`
	Name  string  `json:"Name"`
	Desc  string  `json:"Desc"`
	Price float64 `json:"Price"`
}

var inventory []Item

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Gagan ")
}
func getInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/jon")
	fmt.Println("Function Called:getInventory()")
	json.NewEncoder(w).Encode(inventory)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	inventory = append(inventory, item)
	json.NewEncoder(w).Encode(item)
}

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/inventory", getInventory).Methods("GET")
	router.HandleFunc("/inventory", createItem).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
func main() {
	inventory = append(inventory, Item{
		UID:   "0",
		Name:  "Ganesh",
		Desc:  "Engineer",
		Price: 88,
	})
	inventory = append(inventory, Item{
		UID:   "1",
		Name:  "Gagan",
		Desc:  "software",
		Price: 8800,
	})

	handleRequest()
}
