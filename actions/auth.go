package actions

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"goesb/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func AuthNew(c buffalo.Context) error {
	c.Set("user", models.User{})

	return c.Render(200, r.HTML("auth/signin.html"))
}

func AuthCreate(c buffalo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return errors.WithStack(err)
	}

	tx := c.Value("tx").(*pop.Connection)

	err := tx.Where("username = ?", strings.ToLower(u.Username)).First(u)

	bad := func() error {
		c.Set("user", u)

		textError := T.Translate(c, "user.auth.failed")
		verrs := validate.NewErrors()
		verrs.Add("username", textError)
		c.Set("errors", verrs)

		c.Flash().Add("danger", textError)

		return c.Render(422, r.HTML("auth/signin.html"))
	}

	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return bad()
		}
		return errors.WithStack(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(u.Password))
	if err != nil {
		return bad()
	}
	c.Session().Set("current_user_id", u.ID)
	c.Flash().Add("success", "Welcome Back to Buffalo!")

	return c.Redirect(302, "/")
}

func AuthDestroy(c buffalo.Context) error {
	c.Session().Clear()
	c.Flash().Add("success", "You have been logged out!")
	return c.Redirect(302, "/")
}

func AuthNewSignup(c buffalo.Context) error {
	c.Set("user", models.User{})

	return c.Render(200, r.HTML("auth/signup.html"))
}

func AuthCreateSignup(c buffalo.Context) error {
	user := &models.User{}

	if err := c.Bind(user); err != nil {
		return err
	}

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	verrs, err := user.Create(tx)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		c.Set("errors", verrs)

		c.Set("user", user)

		return c.Render(http.StatusUnprocessableEntity, r.HTML("auth/signup.plush.html"))
	}

	c.Session().Set("current_user_id", user.ID)
	c.Flash().Add("success", "Welcome Back to Buffalo!")

	return c.Redirect(302, "/")
}
