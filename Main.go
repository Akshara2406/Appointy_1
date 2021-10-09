package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	ID      string `json:"ID"`
	DOB     string `json:"DOB"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{ID: "akshu_2406", DOB: "24-06-2002", Content: "Smile to make life matter"},
	}

	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")

}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func main() {
	handleRequests()
}
