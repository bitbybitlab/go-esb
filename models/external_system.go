package models

import (
	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

type ExternalSystem struct {
	BaseDirectoryModel
	Name   string         `db:"name"`
	TypeID uuid.UUID      `db:"type"`
	Type   ConnectionType `belongs_to:"type"`
	Ip     nulls.String   `db:"ip"`
	Port   nulls.String   `db:"port"`
	Path   nulls.String   `db:"path"`
	Driver nulls.String   `db:"driver"`
}

type ExternalSystems []ExternalSystem

func (object *ExternalSystem) Create(tx *pop.Connection) (*validate.Errors, error) {
	object.Version = 1

	return tx.ValidateAndCreate(object)
}

func (object *ExternalSystem) Update(tx *pop.Connection) error {
	object.Version += 1

	return tx.Update(object)
}
