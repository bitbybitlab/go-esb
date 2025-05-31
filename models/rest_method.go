package models

import (
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

type RestMethod struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
}

func (d RestMethod) TableName() string {
	return "rest_methods"
}

func (object *RestMethod) Create(tx *pop.Connection) (*validate.Errors, error) {
	return tx.ValidateAndCreate(object)
}

func (object *RestMethod) Update(tx *pop.Connection) (*validate.Errors, error) {
	return tx.ValidateAndUpdate(object)
}
