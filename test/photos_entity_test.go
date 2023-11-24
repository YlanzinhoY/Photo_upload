package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ylanzinhoy/profile_with_photo_upload/internal/entity"
	uuidgenerate "github.com/ylanzinhoy/profile_with_photo_upload/pkg/uuidGenerate"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestPhotoValidationValidPhoto(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NoError(t, db.AutoMigrate(&entity.Photo{}))

	validPhoto := &entity.Photo{
		ID:   uuidgenerate.GenerateUuid(),
		Name: "validPhoto",
		Data: []byte("validByte"),
	}

	err = db.Create(validPhoto).Error
	assert.NoError(t, err)

}

func TestPhotoValidationInvalidPhoto(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NoError(t, db.AutoMigrate(&entity.Photo{}))

	invalidPhoto := &entity.Photo{
		ID:   uuidgenerate.GenerateUuid(),
		Name: "InvalidPhoto",
		Data: nil,
	}

	err = db.Create(invalidPhoto).Error
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "data cannot be null")
}
