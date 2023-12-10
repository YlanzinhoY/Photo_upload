package database

import (
	"context"

	"github.com/ylanzinhoy/profile_with_photo_upload/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Photo struct {
	DB *mongo.Collection
}

func NewPhoto(db *mongo.Collection) *Photo {
	return &Photo{DB: db}
}

func (p *Photo) Upload(photo *entity.Photo) error {
	_, err := p.DB.InsertOne(context.TODO(), photo)
	if err != nil {
		return err
	}
	return nil
}

func (p *Photo) FindPhotoByName(name string) (*entity.Photo, error) {
	var photo entity.Photo
	filter := bson.D{
		{Key: "name", Value: name},
	}
	err := p.DB.FindOne(context.TODO(), filter).Decode(&photo)
	return &photo, err
}
