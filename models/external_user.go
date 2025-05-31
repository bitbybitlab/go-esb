package models

import (
	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

type ExternalUser struct {
	BaseDirectoryModel
	Name     string             `db:"name"`
	SystemID uuid.UUID          `db:"system"`
	System   ExternalSystem     `belongs_to:"system"`
	TypeID   uuid.UUID          `db:"type"`
	Type     AuthenticationType `belongs_to:"type"`
	Username nulls.String       `db:"username"`
	Password nulls.String       `db:"password"`
	Token    nulls.String       `db:"token"`
}

type ExternalUsers []ExternalUser

func (object *ExternalUser) Create(tx *pop.Connection) (*validate.Errors, error) {
	object.Version = 1

	return tx.ValidateAndCreate(object)
}

func (object *ExternalUser) Update(tx *pop.Connection) (*validate.Errors, error) {
	object.Version += 1

	return tx.ValidateAndUpdate(object)
}
