package main

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
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

func getComments() ([]Comment, error) {
	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var comments []Comment
	for cur.Next(context.TODO()) {
		var c Comment
		err := cur.Decode(&c)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}

	cur.Close(context.TODO())
	return comments, nil
}

func displayCommentsHandler(w http.ResponseWriter, r *http.Request) {
	comments, err := getComments()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comments)
}

func main() {
	initDB()
	http.HandleFunc("/comments", displayCommentsHandler)
	http.ListenAndServe(":8081", nil)
}
