package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	fmt.Println("use of json")
	json.NewEncoder(w).Encode(articles)
}

func now(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Happy New Year")
}

func handleRequest() {
	http.HandleFunc("/", now)
	http.HandleFunc("/a", allArticles)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	handleRequest()
}
