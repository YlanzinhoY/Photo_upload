package entity

import (
	"errors"

	uuidgenerate "github.com/ylanzinhoy/profile_with_photo_upload/pkg/uuidGenerate"
	"gorm.io/gorm"
)

type Photo struct {
	ID   uuidgenerate.ID `gorm:"primaryKey"`
	Name string          `gorm:"not null"`
	Data []byte          `gorm:"type:blob;not null"`
}

func NewFile(name string, data []byte) (*Photo, error) {
	return &Photo{
		ID:   uuidgenerate.GenerateUuid(),
		Name: name,
		Data: data,
	}, nil
}

func (p *Photo) BeforeSave(tx *gorm.DB) (err error) {
	if len(p.Data) == 0 {
		return errors.New("data cannot be null")
	}

	return nil
}