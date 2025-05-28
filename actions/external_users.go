package actions

import (
	"fmt"
	"net/http"

	"goesb/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

type ExternalUsersResource struct {
	buffalo.Resource
}

func (v ExternalUsersResource) List(c buffalo.Context) error {

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	ex_usrs := &models.ExternalUsers{}

	q := tx.PaginateFromParams(c.Params())

	if err := q.Eager().All(ex_usrs); err != nil {
		return err
	}

	c.Set("pagination", q.Paginator)

	c.Set("ex_usrs", ex_usrs)
	return c.Render(http.StatusOK, r.HTML("external_users/index.plush.html"))
}

func (v ExternalUsersResource) Show(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("external_users/index.plush.html"))
}

func (v ExternalUsersResource) New(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("external_users/index.plush.html"))
}

func (v ExternalUsersResource) Create(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("external_users/index.plush.html"))
}

func (v ExternalUsersResource) Edit(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("external_users/index.plush.html"))
}

func (v ExternalUsersResource) Update(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("external_users/index.plush.html"))
}

func (v ExternalUsersResource) Destroy(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("external_users/index.plush.html"))
}
