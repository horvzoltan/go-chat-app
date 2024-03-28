package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("uri")))
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	defer client.Disconnect(context.Background())

	fmt.Println("Connected to MongoDB Atlas!")

	fs := http.FileServer(http.Dir("web"))
	http.Handle("/", fs)

	http.HandleFunc("/api/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
			return
		}

		var data struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Error decoding request body", http.StatusBadRequest)
			return
		}

		err = registerUser(client, data.Username, data.Email, data.Password)
		if err != nil {
			log.Printf("Error registering user: %v\n", err)
			http.Error(w, "Error registering user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "User registered successfully")
	})

	fmt.Println("Starting server at port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func registerUser(client *mongo.Client, username, email, password string) error {
	collection := client.Database("chat-app").Collection("users")

	hashedPassword, err := HashPassword(password)

	if err != nil {
		return err
	}

	document := bson.D{
		{"username", username},
		{"email", email},
		{"password", hashedPassword},
	}

	_, err = collection.InsertOne(context.TODO(), document)
	return err
}

func loginUser(client *mongo.Client, username, password string) (bool, error) {
	collection := client.Database("chat-app").Collection("users")

	var user struct {
		Username string
		Password string
	}

	err := collection.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&user)
	if err != nil {
		return false, err // User not found
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, nil // Invalid password
	}

	return true, nil // Success
}
