package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux" // Routing library
	"github.com/rs/cors"     // CORS library
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Post struct to represent a blog post
type Post struct {
	ID      int    `json:"_id,omitempty" bson:"_id"` // Use int for ID (consider potential limitations)
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Global variables (replace with your connection details)
var (
	client *mongo.Client
	db     *mongo.Database
	nextID int // To track the next available ID (simple approach, not production-ready)
)

func main() {
	// Connect to MongoDB
	ctx := context.Background()
	uri := "mongodb://localhost:27017" // Replace with your MongoDB connection URI
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Get database reference
	db = client.Database("blog") // Replace with your database name

	// Routing setup with CORS
	router := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // Replace with your frontend origin
		AllowedMethods: []string{"POST", "GET"},           // Allow POST and GET methods
		AllowedHeaders: []string{"Content-Type"},          // Allow Content-Type header
		if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}})
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
	handler := c.Handler(router)

	router.HandleFunc("/api/posts", createPostHandler).Methods(http.MethodPost)
	// ... (other routes if needed)

	// Start server
	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

// Function to handle creating a new post
func createPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Decode request body
	var post Post
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding request body: %v", err)
		return
	}

	// Assign next available ID
	post.ID = nextID
	nextID++ // Increment for next post

	// Insert post into MongoDB
	ctx := context.Background()
	col := db.Collection("posts")     // Replace "posts" with your collection name
	_, err = col.InsertOne(ctx, post) // Ignoring the returned insertion result (consider handling it)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error inserting post: %v", err)
		return
	}

	// Respond with success message and the assigned ID
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Post created successfully", "id": post.ID})
	// Run the server
	log.Fatal(http.ListenAndServe(":3001", router))
	router.Run(":3001")
}
