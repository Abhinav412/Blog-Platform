package main

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Comment struct {
	ID      string json:"id" bson:"_id,omitempty"
	Content string json:"content" bson:"content"
}

var collection *mongo.Collection

func initDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	collection = client.Database("test").Collection("comments")
}

func insertComment(comment Comment) error {
	_, err := collection.InsertOne(context.TODO(), comment)
	return err
}

func saveCommentHandler(w http.ResponseWriter, r *http.Request) {
	var c Comment
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = insertComment(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func main() {
	initDB()
	http.HandleFunc("/comments", saveCommentHandler)
	http.ListenAndServe(":8080", nil)
}