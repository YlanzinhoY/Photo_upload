package entity

import (
	"github.com/go-playground/validator/v10"
	uuidgenerate "github.com/ylanzinhoy/profile_with_photo_upload/pkg/uuidGenerate"
)

type Photo struct {
	ID   uuidgenerate.ID `bson:"id,omitempty"`
	Name string          `bson:"name" validate:"required"`
	Data []byte          `bson:"data" validate:"required"`
}

func NewFile(name string, data []byte) (*Photo, error) {
	return &Photo{
		ID:   uuidgenerate.GenerateUuid(),
		Name: name,
		Data: data,
	}, nil
}

func (p *Photo) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}
