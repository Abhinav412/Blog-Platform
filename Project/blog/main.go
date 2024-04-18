package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type BlogPost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	// Initialize Gin-gonic router
	router := gin.Default()

	// MongoDB connection setup
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	// Define MongoDB collection
	collection := client.Database("myblog").Collection("posts")

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	// API endpoint for creating a new post
	router.POST("/api/posts", func(c *gin.Context) {
		var post BlogPost
		if err := c.ShouldBindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Insert the post into MongoDB
		_, err := collection.InsertOne(context.TODO(), post)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Post created successfully"})
	})

	// Run the server
	log.Fatal(http.ListenAndServe(":3001", router))
	router.Run(":3001")
}
