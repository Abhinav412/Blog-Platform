package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
	ID        string    `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID    string    `json:"userId"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

var (
	client *mongo.Client
	db     *mongo.Database
)

func main() {
	ctx := context.Background()
	uri := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	db = client.Database("blog")

	router := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	handler := c.Handler(router)
	router.HandleFunc("/api/users/signup", signupHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/users/login", loginHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/posts", createPostHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/posts", getPostsByAuthorHandler).Methods(http.MethodGet)

	router.HandleFunc("/api/posts/{id}", deletePostHandler).Methods(http.MethodDelete)

	router.Methods(http.MethodOptions).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.WriteHeader(http.StatusNoContent)
	})

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	col := db.Collection("users")

	var newUser User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding user data: %v", err)
		return
	}

	hashedPassword, err := hashPassword([]byte(newUser.Password))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error hashing password: %v", err)
		return
	}
	newUser.Password = string(hashedPassword)

	newUser.Posts = []Post{}

	_, err = col.InsertOne(ctx, newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error creating user: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

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

	if err := comparePassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	response := struct {
		UserID string `json:"userId"`
	}{
		UserID: user.ID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

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

	newPost.ID, err = getNextPostID(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error generating post ID: %v", err)
		return
	}

	newPost.CreatedAt = time.Now()
	newPost.UpdatedAt = time.Now()

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
	col := db.Collection("posts")

	opts := options.FindOne().SetSort(bson.D{{Key: "createdAt", Value: -1}})
	result := col.FindOne(ctx, bson.D{}, opts)
	var post Post
	if err := result.Decode(&post); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error fetching last created post: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct{ Author string }{post.Author})
}

func getPostsByAuthorHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	col := db.Collection("posts") // Change collection name to "posts"

	author := r.URL.Query().Get("author")

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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func getNextPostID(ctx context.Context) (string, error) {
	col := db.Collection("postsCounters") // Counter collection

	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	update := bson.M{"$inc": bson.M{"value": 1}}
	result := col.FindOneAndUpdate(ctx, bson.M{}, update, opts)
	if result.Err() != nil {
		return "", result.Err()
	}

	var counter struct {
		Value int `bson:"value"`
	}
	err := result.Decode(&counter)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(counter.Value), nil
}

func deletePostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	col := db.Collection("posts")

	postTitle := mux.Vars(r)["title"]

	filter := bson.M{"title": postTitle}

	_, err := col.DeleteOne(ctx, filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error deleting post: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Post deleted successfully")
}
