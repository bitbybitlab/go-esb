package actions

import (
	"fmt"
	"net/http"

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
	return c.Render(http.StatusOK, r.HTML("users/show.plush.html"))
}

func (v UsersResource) New(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("users/new.plush.html"))
}

func (v UsersResource) Create(c buffalo.Context) error {
	c.Set("users", getAllUsers())

	return c.Render(http.StatusOK, r.HTML("users/index.plush.html"))
}

func (v UsersResource) Update(c buffalo.Context) error {
	c.Set("users", getAllUsers())

	return c.Render(http.StatusOK, r.HTML("users/index.plush.html"))
}

func (v UsersResource) Destroy(c buffalo.Context) error {
	c.Set("users", getAllUsers())

	return c.Render(http.StatusOK, r.HTML("users/index.plush.html"))
}

func getAllUsers() []models.User {
	return models.GetAllUsers()
}
