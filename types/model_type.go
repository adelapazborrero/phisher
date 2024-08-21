package types

import (
	"time"

	"github.com/google/uuid"
)

type ID string

func NewID() ID {
	id := uuid.New()
	return ID(id.String())
}

func TimePtr(t time.Time) *time.Time {
	return &t
}
