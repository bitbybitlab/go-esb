package models

import (
	"github.com/gofrs/uuid"
)

type BaseDirectoryModel struct {
	ID      uuid.UUID `json:"id" db:"id"`
	Version int       `json:"version" db:"version"`
}
