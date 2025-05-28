package actions

import (
	"fmt"
	"net/http"

	"goesb/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

type ExternalSystemsResource struct {
	buffalo.Resource
}

func (v ExternalSystemsResource) List(c buffalo.Context) error {

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	external_systems := &models.ExternalSystems{}

	q := tx.PaginateFromParams(c.Params())

	if err := q.All(external_systems); err != nil {
		return err
	}

	c.Set("pagination", q.Paginator)

	c.Set("external_systems", external_systems)
	return c.Render(http.StatusOK, r.HTML("external_systems/index.plush.html"))
}

func (v ExternalSystemsResource) Show(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("external_systems/index.plush.html"))
}

func (v ExternalSystemsResource) New(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("external_systems/index.plush.html"))
}

func (v ExternalSystemsResource) Create(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("external_systems/index.plush.html"))
}

func (v ExternalSystemsResource) Edit(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("external_systems/index.plush.html"))
}

func (v ExternalSystemsResource) Update(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("external_systems/index.plush.html"))
}

func (v ExternalSystemsResource) Destroy(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("external_systems/index.plush.html"))
}
