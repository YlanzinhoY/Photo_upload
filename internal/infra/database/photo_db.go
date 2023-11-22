package database

import (
	"github.com/ylanzinhoy/profile_with_photo_upload/internal/entity"
	"gorm.io/gorm"
)


type Photo struct {
	DB *gorm.DB
}

func NewPhoto(db *gorm.DB) *Photo {
	return &Photo{DB: db}
}


func(p *Photo) Upload(photo *entity.Photo) error {
	return p.DB.Create(photo).Error
}

func (p *Photo) FindPhotoByName(name string) (*entity.Photo, error) {
	var photo entity.Photo 
	err := p.DB.First(&photo, "name =?", name).Error
	return &photo, err
}