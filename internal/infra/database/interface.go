package database

import "github.com/ylanzinhoy/profile_with_photo_upload/internal/entity"


type PhotoInterface interface {
	Upload(photo *entity.Photo) error
	FindPhotoByName(name string) (*entity.Photo, error)
}