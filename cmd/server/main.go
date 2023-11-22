package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ylanzinhoy/profile_with_photo_upload/internal/entity"
	"github.com/ylanzinhoy/profile_with_photo_upload/internal/handlers"
	"github.com/ylanzinhoy/profile_with_photo_upload/internal/infra/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Photo{})

	photoDB := database.NewPhoto(db)
	photoHandler := handlers.NewPhotoHandler(photoDB)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.SetHeader("Access-Control-Allow-Origin", "*"))
	r.Use(middleware.SetHeader("Access-Control-Allow-Methods", "GET, POST"))
	r.Use(middleware.SetHeader("Access-Control-Allow-Headers", "Content-Type"))

	r.Post("/photo", photoHandler.PhotoUpload)
	r.Get("/photo/{name}", photoHandler.GetPhotoById)

	http.ListenAndServe(":3333", r)

}
