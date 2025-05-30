package models

import (
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

type ValueType struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
}

func (d ValueType) TableName() string {
	return "value_types"
}

func (object *ValueType) Create(tx *pop.Connection) (*validate.Errors, error) {
	return tx.ValidateAndCreate(object)
}

func (object *ValueType) Update(tx *pop.Connection) (*validate.Errors, error) {
	return tx.ValidateAndUpdate(object)
}
