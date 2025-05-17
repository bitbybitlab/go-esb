package models

import (
	"time"

	"github.com/gobuffalo/nulls"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	Ref        uuid.UUID
	Username   string
	Password   string
	FirstName  nulls.String `db:"first_name"`
	LastName   nulls.String `db:"last_name"`
	MiddleName nulls.String `db:"middle_name"`
	Phone      nulls.String
	Email      nulls.String
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
	Version    int
}

type Users []User
