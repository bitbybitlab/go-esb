package models

import (
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

type ThreadGroup struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	Protocol uuid.UUID `db:"protocol" json:"protocol"`
	Parent   uuid.UUID `db:"parent" json:"parent"`
	Version  int       `json:"version" db:"version"`
}

func (d ThreadGroup) TableName() string {
	return "thread_groups"
}

func (object *ThreadGroup) Create(tx *pop.Connection) (*validate.Errors, error) {
	object.Version = 1

	return tx.ValidateAndCreate(object)
}

func (object *ThreadGroup) Update(tx *pop.Connection) (*validate.Errors, error) {
	object.Version++

	return tx.ValidateAndUpdate(object)
}
