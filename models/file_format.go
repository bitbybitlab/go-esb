package models

import (
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

type FileFormat struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
}

func (d FileFormat) TableName() string {
	return "file_formats"
}

func (object *FileFormat) Create(tx *pop.Connection) (*validate.Errors, error) {
	return tx.ValidateAndCreate(object)
}

func (object *FileFormat) Update(tx *pop.Connection) (*validate.Errors, error) {
	return tx.ValidateAndUpdate(object)
}
