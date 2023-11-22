package uuidgenerate

import (
	"github.com/google/uuid"
)

type ID = uuid.UUID

func GenerateUuid() ID {
	return ID(uuid.New())
}

func ParseID(s string) (ID, error) {
	id, err := uuid.Parse(s)

	return ID(id), err
}
