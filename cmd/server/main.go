package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/ylanzinhoy/profile_with_photo_upload/internal/handlers"
	"github.com/ylanzinhoy/profile_with_photo_upload/internal/infra/database"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading .env file")
	} else {
		fmt.Println("Success loading.env file")
	}
}

func main() {

	mongodbUri := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("MONGO_DB_NAME")
	collectionName := os.Getenv("MONGO_DB_COLLECTION")

	clientOptions := options.Client().ApplyURI(mongodbUri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(context.TODO())

	photoDB := database.NewPhoto(client.Database(dbName).Collection(collectionName))
	photoHandler := handlers.NewPhotoHandler(photoDB)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.SetHeader("Access-Control-Allow-Origin", "*"))
	r.Use(middleware.SetHeader("Access-Control-Allow-Methods", "GET, POST"))
	r.Use(middleware.SetHeader("Access-Control-Allow-Headers", "Content-Type"))

	r.Post("/photo", photoHandler.PhotoUpload)
	r.Get("/photo/{name}", photoHandler.GetPhotoByName)

	http.ListenAndServe(fmt.Sprintf(":%s",os.Getenv("PORT")), r)

}
