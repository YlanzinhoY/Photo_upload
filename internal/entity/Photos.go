package entity

import (
	uuidgenerate "github.com/ylanzinhoy/profile_with_photo_upload/pkg/uuidGenerate"
)

type Photo struct {
	ID   uuidgenerate.ID `gorm:"primaryKey"`
	Name string          `gorm:"not null"`
	Data []byte          `gorm:"null"`
}

func NewFile(name string, data []byte) (*Photo, error) {
	return &Photo{
		ID:   uuidgenerate.GenerateUuid(),
		Name: name,
		Data: data,
	}, nil
}
