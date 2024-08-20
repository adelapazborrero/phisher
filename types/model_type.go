package types

import "github.com/google/uuid"

type ID string

func NewID() ID {
	id := uuid.New()
	return ID(id.String())
}
