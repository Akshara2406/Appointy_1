package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoField struct {
	ID       string `json: "ID"`
	Name     string `json: "Name"`
	Email    string `json: "Email"`
	Password string `json: "Password"`
}

type MongoFields []MongoField

func allArticles(w http.ResponseWriter, r *http.Request) {
	mongofields := MongoFields{
		MongoField{ID: "akshu_2406", Name: "Akshara", Email: "mageshakshara24@gmail.com", Password: "abc123tch"},
		MongoField{ID: "akshith_2341", Name: "Akshith", Email: "mageshakshith29@gmail.com", Password: "abc1245tch"},
	}

	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(mongofields)
}

type Post struct {
	ID        string `json: "ID"`
	Caption   string `json: "Caption"`
	URL       string `json: "URL"`
	Timestamp string `json: "Timestamp"`
}

type Posts []Post

func allPosts(w http.ResponseWriter, r *http.Request) {
	posts := Posts{
		Post{ID: "akshu_2406", Caption: "Great smile", URL: "abc@instagram.com", Timestamp: "10/03/2017 07:29:46"},
		Post{ID: "akshith_2341", Caption: "We stand together", URL: "acd@instagram.com", Timestamp: "10/07/2017 07:29:46"},
	}

	fmt.Println("Endpoint Hit: 	All posts Endpoint")
	json.NewEncoder(w).Encode(posts)
}

type Auser struct {
	ID        string `json: "ID"`
	Name      string `json: "Name"`
	Email     string `json: "Email"`
	Password  string `json: "Password"`
	Caption   string `json: "Caption"`
	URL       string `json: "URL"`
	Timestamp string `json: "Timestamp"`
}

type Ausers []Auser

func allUsers(w http.ResponseWriter, r *http.Request) {
	ausers := Ausers{
		Auser{ID: "akshu_2406", Name: "Akshara", Email: "mageshakshara24@gmail.com", Password: "abc123tch", Caption: "Great smile", URL: "abc@instagram.com", Timestamp: "10/03/2017 07:29:46"},
		Auser{ID: "akshith_2341", Name: "Akshith", Email: "mageshakshith29@gmail.com", Password: "abc1245tch", Caption: "We stand together", URL: "acd@instagram.com", Timestamp: "10/07/2017 07:29:46"},
	}

	fmt.Println("Endpoint Hit: 	All posts Endpoint")
	json.NewEncoder(w).Encode(ausers)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")

}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", allArticles)
	http.HandleFunc("/posts", allPosts)
	http.HandleFunc("/posts/users", allUsers)

	log.Fatal(http.ListenAndServe(":8081", nil))

}

func main() {
	handleRequests()
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

}
