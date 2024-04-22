package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// BlogPost represents a blog post structure
type bp struct {
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

	// Ping the MongoDB server to check the connection
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}

	// Define MongoDB collection
	collection := client.Database("myblog").Collection("posts")

	// CORS middleware
	router.Use(cors.Default())

	// API endpoint for fetching all posts
	router.GET("/api/posts", func(c *gin.Context) {
		var posts []bp

		// Query MongoDB to get all posts
		cursor, err := collection.Find(context.TODO(), nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
			return
		}
		defer cursor.Close(context.TODO())

		// Iterate through the cursor and decode each document into a BlogPost
		for cursor.Next(context.TODO()) {
			var post bp
			if err := cursor.Decode(&post); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode post"})
				return
			}
			posts = append(posts, post) // This line appends each post to the posts slice
		}

		// Check for errors during cursor iteration
		if err := cursor.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cursor error"})
			return
		}

		// Return the posts as JSON
		c.JSON(http.StatusOK, posts)
	})

	// Run the server
	if err := router.Run(":3000"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
