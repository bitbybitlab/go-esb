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
	BaseDirectoryModel
	Username             string       `db:"username"`
	PasswordHash         string       `db:"password"`
	FirstName            nulls.String `db:"first_name"`
	LastName             nulls.String `db:"last_name"`
	MiddleName           nulls.String `db:"middle_name"`
	Phone                nulls.String `db:"phone"`
	Email                nulls.String `db:"email"`
	Password             string       `json:"-" db:"-"`
	PasswordConfirmation string       `json:"-" db:"-"`
}

type Users []User

func (object *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.StringIsPresent{Field: object.Username, Name: "Username"},
		&validators.StringIsPresent{Field: object.PasswordHash, Name: "PasswordHash"},
		&validators.FuncValidator{
			Field:   object.Username,
			Name:    "Username",
			Message: "%s is already taken",
			Fn: func() bool {
				var b bool
				q := tx.Where("username = ?", object.Username)
				if object.ID != uuid.Nil {
					q = q.Where("id != ?", object.ID)
				}
				b, err = q.Exists(object)
				if err != nil {
					return false
				}
				return !b
			},
		},
	), err
}

func (object *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.StringIsPresent{Field: object.Password, Name: "Password"},
		&validators.StringsMatch{Name: "Password", Field: object.Password, Field2: object.PasswordConfirmation,
			Message: "PasswordHash does not match confirmation"},
	), err
}

func (object *User) Create(tx *pop.Connection) (*validate.Errors, error) {
	object.Username = strings.ToLower(strings.TrimSpace(object.Username))
	ph, err := bcrypt.GenerateFromPassword([]byte(object.Password), bcrypt.DefaultCost)
	if err != nil {
		return validate.NewErrors(), errors.WithStack(err)
	}
	object.PasswordHash = string(ph)
	object.CreateTime = time.Now()
	object.UpdateTime = object.CreateTime
	object.Version = 1

	return tx.ValidateAndCreate(object)
}

func (object *User) Update(tx *pop.Connection) (*validate.Errors, error) {
	object.UpdateTime = time.Now()
	object.Version += 1

	return tx.ValidateAndUpdate(object)
}
