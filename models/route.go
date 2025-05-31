package models

import (
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

type Route struct {
	ID      uuid.UUID `json:"id" db:"id"`
	Name    string    `json:"name" db:"name"`
	Path    string    `json:"path" db:"path"`
	System  uuid.UUID `json:"system" db:"system"`
	Method  uuid.UUID `json:"method" db:"method"`
	Version int       `json:"version" db:"version"`
}

func (d Route) TableName() string {
	return "routes"
}

func (object *Route) Create(tx *pop.Connection) (*validate.Errors, error) {
	object.Version = 1

	return tx.ValidateAndCreate(object)
}

func (object *Route) Update(tx *pop.Connection) (*validate.Errors, error) {
	object.Version++

	return tx.ValidateAndUpdate(object)
}
