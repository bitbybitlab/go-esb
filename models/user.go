package models

import (
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gofrs/uuid"
)

type User struct {
	ID         uuid.UUID    `db:"id"`
	Username   string       `db:"username"`
	Password   string       `db:"password"`
	FirstName  nulls.String `db:"first_name"`
	LastName   nulls.String `db:"last_name"`
	MiddleName nulls.String `db:"middle_name"`
	Phone      nulls.String `db:"phone"`
	Email      nulls.String `db:"email"`
	CreateTime time.Time    `db:"create_time"`
	UpdateTime time.Time    `db:"update_time"`
	Version    int          `db:"version"`
}

type Users []User
