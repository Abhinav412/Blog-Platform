package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux" // Routing library
	"github.com/rs/cors"     // CORS library
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password,omitempty" bson:"password"`
	Posts    []Post `json:"posts,omitempty" bson:"posts"`
}

type Post struct {
	ID        string    `json:"_id,omitempty" bson:"_id"`
	UserID    string    `json:"userId"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

// Function to connect to MongoDB (moved outside main for reusability)
var (
	client *mongo.Client
	db     *mongo.Database
)

func main() {
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

	// CORS configuration
	c := cors.New(cors.Options{
<<<<<<< HEAD
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum age for caching
	})
	handler := c.Handler(router)
	router.HandleFunc("/api/users/signup", signupHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/users/login", loginHandler).Methods(http.MethodPost)
=======
		AllowedOrigins: []string{"http://localhost:3000"}, // Replace with your frontend origin
		AllowedMethods: []string{"POST", "GET"},           // Allow POST and GET methods
		AllowedHeaders: []string{"Content-Type"},          // Allow Content-Type header
	})
	handler := c.Handler(router)

	// Route handlers
>>>>>>> ed02bd0251a08728dfcdf00c6f7ad330df7451b2
	router.HandleFunc("/api/posts", createPostHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/posts", getPostsByAuthorHandler).Methods(http.MethodGet)

	// Handle preflight requests explicitly
	router.Methods(http.MethodOptions).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.WriteHeader(http.StatusNoContent)
	})

	// Start server
	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	col := db.Collection("users") // Replace "users" with your collection name

	var newUser User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding user data: %v", err)
		return
	}

	// Hash password before storing (convert to []byte)
	hashedPassword, err := hashPassword([]byte(newUser.Password))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error hashing password: %v", err)
		return
	}
	newUser.Password = string(hashedPassword)

	// Initialize user document with empty posts array
	newUser.Posts = []Post{}

	// Insert new user with initialized posts field
	_, err = col.InsertOne(ctx, newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error creating user: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
<<<<<<< HEAD
=======
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Post created successfully", "id": post.ID})
>>>>>>> ed02bd0251a08728dfcdf00c6f7ad330df7451b2
}

// hashPassword hashes the given password using bcrypt
func hashPassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}

func comparePassword(hashedPassword []byte, plainPassword []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, plainPassword)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	col := db.Collection("users")

	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&credentials); err != nil {
		http.Error(w, "Error decoding credentials", http.StatusBadRequest)
		return
	}

	// Find user by username
	filter := bson.M{"username": credentials.Username}
	result := col.FindOne(ctx, filter)
	var user User
	if err := result.Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Error fetching user", http.StatusInternalServerError)
		return
	}

	// Compare hashed passwords
	if err := comparePassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Login successful, send back user ID in response
	// You may also include other user details if needed
	response := struct {
		UserID string `json:"userId"`
	}{
		UserID: user.ID,
	}

	// Encode response as JSON and send it back
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Create post handler
// Create post handler
// Create post handler
func createPostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	col := db.Collection("posts") // Change collection name to "posts"

	var newPost Post
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newPost)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding post data: %v", err)
		return
	}

	// Assign a new unique ID for the post
	newPost.ID, err = getNextPostID(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error generating post ID: %v", err)
		return
	}

	// Set timestamps
	newPost.CreatedAt = time.Now()
	newPost.UpdatedAt = time.Now()

	// Insert new post into posts collection
	_, err = col.InsertOne(ctx, newPost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error creating post: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
func getLastCreatedPostAuthorHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	col := db.Collection("posts") // Change collection name to "posts"

	// Find the last created post
	opts := options.FindOne().SetSort(bson.D{{Key: "createdAt", Value: -1}})
	result := col.FindOne(ctx, bson.D{}, opts)
	var post Post
	if err := result.Decode(&post); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error fetching last created post: %v", err)
		return
	}

	// Encode the author of the last created post as JSON and write to response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct{ Author string }{post.Author})
}

func getPostsByAuthorHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	col := db.Collection("posts") // Change collection name to "posts"

	// Extract author from query parameters
	author := r.URL.Query().Get("author")

	// Find posts by author
	filter := bson.M{"author": author}
	cursor, err := col.Find(ctx, filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error fetching posts by author: %v", err)
		return
	}
	defer cursor.Close(ctx)

	var posts []Post
	for cursor.Next(ctx) {
		var post Post
		if err := cursor.Decode(&post); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error decoding post: %v", err)
			return
		}
		posts = append(posts, post)
	}

	// Encode posts as JSON and write to response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
<<<<<<< HEAD
}
func getPostByIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	col := db.Collection("posts") // Change collection name to "posts"

	// Parse post ID from URL path
	postID := mux.Vars(r)["postID"]

	// Find post by ID
	filter := bson.M{"_id": postID}
	result := col.FindOne(ctx, filter)
	var post Post
	if err := result.Decode(&post); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error fetching post: %v", err)
		return
	}

	// Encode post as JSON and write to response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func getCurrentUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	col := db.Collection("users") // No change here since it's still user-related

	// Parse user ID from request headers (assuming it's stored in the "userID" header)
	userID := r.Header.Get("userID")
	if userID == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "User ID not found in request headers")
		return
	}

	// Find user by ID
	filter := bson.M{"_id": userID}
	result := col.FindOne(ctx, filter)
	var user User
	if err := result.Decode(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error fetching user: %v", err)
		return
	}

	// Encode user as JSON and write to response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func getNextPostID(ctx context.Context) (string, error) {
	col := db.Collection("postsCounters") // Counter collection

	// Find the document for posts counter
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	update := bson.M{"$inc": bson.M{"value": 1}}
	result := col.FindOneAndUpdate(ctx, bson.M{}, update, opts)
	if result.Err() != nil {
		return "", result.Err()
	}

	// Decode the counter value
	var counter struct {
		Value int `bson:"value"`
	}
	err := result.Decode(&counter)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(counter.Value), nil
=======
>>>>>>> ed02bd0251a08728dfcdf00c6f7ad330df7451b2
}
