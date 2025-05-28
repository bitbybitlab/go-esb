package models

import (
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
)

type ExternalSystem struct {
	BaseDirectoryModel
	Name   string       `db:"name"`
	Ip     nulls.String `db:"ip"`
	Port   nulls.String `db:"port"`
	Path   nulls.String `db:"path"`
	Driver nulls.String `db:"driver"`
}

type ExternalSystems []ExternalSystem

func (object *ExternalSystem) Create(tx *pop.Connection) (*validate.Errors, error) {
	object.CreateTime = time.Now()
	object.UpdateTime = object.CreateTime
	object.Version = 1

	return tx.ValidateAndCreate(object)
}

func (object *ExternalSystem) Update(tx *pop.Connection) (*validate.Errors, error) {
	object.UpdateTime = time.Now()
	object.Version += 1

	return tx.ValidateAndUpdate(object)
}
