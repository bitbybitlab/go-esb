package models

import (
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

type ThreadObject struct {
	ID         uuid.UUID `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	NameObject string    `json:"nameObject" db:"name_object"`
	Type       uuid.UUID `json:"type" db:"type"`
	Parent     uuid.UUID `json:"parent" db:"parent"`
	Version    int       `json:"version" db:"version"`
}

func (d ThreadObject) TableName() string {
	return "thread_objects"
}

func (object *ThreadObject) Create(tx *pop.Connection) (*validate.Errors, error) {
	object.Version = 1

	return tx.ValidateAndCreate(object)
}

func (object *ThreadObject) Update(tx *pop.Connection) (*validate.Errors, error) {
	object.Version++

	return tx.ValidateAndUpdate(object)
}
