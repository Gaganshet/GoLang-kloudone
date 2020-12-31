package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"the title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "enter title", Desc: "test description", Content: "the content"},
	}
	fmt.Println("Hello")
	json.NewEncoder(w).Encode(articles)
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Happy new year")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Gagan")
}

func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/a", allArticles).Methods("GET")
	myRouter.HandleFunc("/a", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {
	handleRequest()
}
