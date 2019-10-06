package main

import (
	"api"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func allRestaurants(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(writer, "coming up")
}

func main() {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ := mongo.Connect(ctx, clientOptions)

	r := api.NewRouter()

	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Fatal(err)
	}
}
