package models

import (
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

type Routine struct {
	ID      uuid.UUID `json:"id" db:"id"`
	Name    string    `json:"name" db:"name"`
	Type    uuid.UUID `json:"type" db:"type"`
	Code    string    `json:"code" db:"code"`
	Version int       `json:"version" db:"version"`
}

func (d Routine) TableName() string {
	return "routines"
}

func (object *Routine) Create(tx *pop.Connection) (*validate.Errors, error) {
	object.Version = 1

	return tx.ValidateAndCreate(object)
}

func (object *Routine) Update(tx *pop.Connection) (*validate.Errors, error) {
	object.Version++

	return tx.ValidateAndUpdate(object)
}
