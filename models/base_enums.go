package models

import (
	"github.com/gofrs/uuid"
)

type BaseEnumModel struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
}
