package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "hello world"},
	}
	fmt.Println("hainnnnnn")
	json.NewEncoder(w).Encode(articles)
}

func textPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "post page")
}
func now(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GaneshaChaturthy")
}

func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", now)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", textPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {
	handleRequest()
}
