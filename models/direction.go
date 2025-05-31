package models

import (
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

type Direction struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
}

func (d Direction) TableName() string {
	return "directions"
}

func (object *Direction) Create(tx *pop.Connection) (*validate.Errors, error) {
	return tx.ValidateAndCreate(object)
}

func (object *Direction) Update(tx *pop.Connection) (*validate.Errors, error) {
	return tx.ValidateAndUpdate(object)
}
