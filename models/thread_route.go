package models

import (
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

type ThreadRoute struct {
	Thread     uuid.UUID `json:"thread" db:"thread"`
	Route      uuid.UUID `json:"route" db:"route"`
	Direction  uuid.UUID `json:"direction" db:"direction"`
	FileFormat uuid.UUID `json:"fileFormat" db:"file_format"`
	Object     uuid.UUID `json:"object" db:"object"`
	Routine    uuid.UUID `json:"routine" db:"routine"`
}

func (d ThreadRoute) TableName() string {
	return "thread_routes"
}

func (object *ThreadRoute) Create(tx *pop.Connection) (*validate.Errors, error) {
	return tx.ValidateAndCreate(object)
}

func (object *ThreadRoute) Update(tx *pop.Connection) (*validate.Errors, error) {
	return tx.ValidateAndUpdate(object)
}
