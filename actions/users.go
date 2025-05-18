package actions

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"

	"goesb/models"
)

type UsersResource struct {
	buffalo.Resource
}

func (v UsersResource) List(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	users := &models.Users{}

	q := tx.PaginateFromParams(c.Params())

	if err := q.All(users); err != nil {
		return err
	}

	c.Set("users", users)

	return c.Render(http.StatusOK, r.HTML("users/index.plush.html"))
}

func (v UsersResource) Show(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	user := &models.User{}

	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("user", user)

	return c.Render(http.StatusOK, r.HTML("users/show.plush.html"))
}

func (v UsersResource) New(c buffalo.Context) error {
	c.Set("user", &models.User{})

	return c.Render(http.StatusOK, r.HTML("users/new.plush.html"))
}

func (v UsersResource) Create(c buffalo.Context) error {
	user := &models.User{}

	if err := c.Bind(user); err != nil {
		return err
	}

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	user.CreateTime = time.Now()
	user.UpdateTime = user.CreateTime
	user.Version = 1

	verrs, err := tx.ValidateAndCreate(user)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		c.Set("errors", verrs)

		c.Set("user", user)

		return c.Render(http.StatusUnprocessableEntity, r.HTML("users/new.plush.html"))
	}

	c.Flash().Add("success", T.Translate(c, "user.created.success"))

	return c.Redirect(http.StatusSeeOther, "/users/%v", user.ID)
}

func (v UsersResource) Edit(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	user := &models.User{}

	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("user", user)

	return c.Render(http.StatusOK, r.HTML("users/edit.plush.html"))
}

func (v UsersResource) Update(c buffalo.Context) error {

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	user := &models.User{}

	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := c.Bind(user); err != nil {
		return err
	}

	user.UpdateTime = time.Now()
	user.Version += 1

	verrs, err := tx.ValidateAndUpdate(user)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		c.Set("errors", verrs)
		c.Set("user", user)

		return c.Render(http.StatusUnprocessableEntity, r.HTML("users/edit.plush.html"))
	}

	c.Flash().Add("success", T.Translate(c, "user.updated.success"))

	return c.Redirect(http.StatusSeeOther, "/users/%v", user.ID)
}

func (v UsersResource) Destroy(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	user := &models.User{}

	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(user); err != nil {
		return err
	}

	c.Flash().Add("success", T.Translate(c, "user.destroyed.success"))

	return c.Redirect(http.StatusSeeOther, "/users")
}
