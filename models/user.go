package models

import (
	"strings"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                   uuid.UUID    `db:"id"`
	Username             string       `db:"username"`
	PasswordHash         string       `db:"password"`
	FirstName            nulls.String `db:"first_name"`
	LastName             nulls.String `db:"last_name"`
	MiddleName           nulls.String `db:"middle_name"`
	Phone                nulls.String `db:"phone"`
	Email                nulls.String `db:"email"`
	CreateTime           time.Time    `db:"create_time"`
	UpdateTime           time.Time    `db:"update_time"`
	Version              int          `db:"version"`
	Password             string       `json:"-" db:"-"`
	PasswordConfirmation string       `json:"-" db:"-"`
}

type Users []User

func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Username, Name: "Username"},
		&validators.StringIsPresent{Field: u.PasswordHash, Name: "PasswordHash"},
		&validators.FuncValidator{
			Field:   u.Username,
			Name:    "Username",
			Message: "%s is already taken",
			Fn: func() bool {
				var b bool
				q := tx.Where("username = ?", u.Username)
				if u.ID != uuid.Nil {
					q = q.Where("id != ?", u.ID)
				}
				b, err = q.Exists(u)
				if err != nil {
					return false
				}
				return !b
			},
		},
	), err
}

func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Password, Name: "Password"},
		&validators.StringsMatch{Name: "Password", Field: u.Password, Field2: u.PasswordConfirmation, Message: "PasswordHash does not match confirmation"},
	), err
}

func (u *User) Create(tx *pop.Connection) (*validate.Errors, error) {
	u.Username = strings.ToLower(strings.TrimSpace(u.Username))
	ph, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return validate.NewErrors(), errors.WithStack(err)
	}
	u.PasswordHash = string(ph)
	u.CreateTime = time.Now()
	u.UpdateTime = u.CreateTime
	u.Version = 1

	return tx.ValidateAndCreate(u)
}

func (u *User) Update(tx *pop.Connection) (*validate.Errors, error) {
	u.UpdateTime = time.Now()
	u.Version += 1

	return tx.ValidateAndUpdate(u)
}
