package models

import (
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

type Thread struct {
	ID                   uuid.UUID `json:"id" db:"id"`
	Name                 string    `json:"name" db:"name"`
	Group                uuid.UUID `json:"group" db:"group"`
	MessageConverterType uuid.UUID `json:"messageConverterType" db:"message_converter_type"`
	Version              int       `json:"version" db:"version"`
}

func (d Thread) TableName() string {
	return "threads"
}

func (object *Thread) Create(tx *pop.Connection) (*validate.Errors, error) {
	object.Version = 1

	return tx.ValidateAndCreate(object)
}

func (object *Thread) Update(tx *pop.Connection) (*validate.Errors, error) {
	object.Version++

	return tx.ValidateAndUpdate(object)
}
