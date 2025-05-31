package actions

import (
	"fmt"
	"goesb/models"
	"net/http"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
)

var tmpl_index = "index.plush.html"
var tmpl_show = "show.plush.html"
var tmpl_new = "new.plush.html"
var tmpl_edit = "edit.plush.html"

type AnyType interface {
	Create(*pop.Connection) (*validate.Errors, error)
	Update(*pop.Connection) error
}

func ListHandler(c buffalo.Context, data interface{}, tmpl_path string) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	q := tx.PaginateFromParams(c.Params())
	if err := q.Eager().All(data); err != nil {
		return err
	}

	c.Set("pagination", q.Paginator)
	c.Set("data", data)

	return c.Render(http.StatusOK, r.HTML(tmpl_path+"/"+tmpl_index))
}

func ShowHandler(c buffalo.Context, data AnyType, tmpl_path string, id_name string) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	if err := tx.Eager().Find(data, c.Param(id_name)); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("data", data)

	return c.Render(http.StatusOK, r.HTML(tmpl_path+"/"+tmpl_show))
}

func NewHandler(c buffalo.Context, data AnyType, tmpl_path string) error {
	c.Set("data", data)

	additionalActions(c, data)

	return c.Render(http.StatusOK, r.HTML(tmpl_path+"/"+tmpl_new))
}

func CreateHandler(c buffalo.Context, data AnyType, tmpl_path string, msg_name string) error {
	if err := c.Bind(data); err != nil {
		return err
	}

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// password confirmation by user can't be blank
	if v, ok := data.(*models.User); ok {
		v.PasswordConfirmation = v.Password
	}

	verrs, err := data.Create(tx)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		c.Set("errors", verrs)

		c.Set("data", data)

		return c.Render(http.StatusUnprocessableEntity, r.HTML(tmpl_path+"/"+tmpl_new))
	}

	c.Flash().Add("success", T.Translate(c, msg_name+".created.success"))

	r_path := strings.ReplaceAll(tmpl_path, "_", "-")
	return c.Redirect(http.StatusSeeOther, "/"+r_path)
}

func EditHandler(c buffalo.Context, data AnyType, tmpl_path string, id_name string) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	if err := tx.Eager().Find(data, c.Param(id_name)); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("data", data)

	additionalActions(c, data)

	return c.Render(http.StatusOK, r.HTML(tmpl_path+"/"+tmpl_edit))
}

func UpdateHandler(c buffalo.Context, data AnyType, tmpl_path string, id_name string, msg_name string) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	if err := tx.Find(data, c.Param(id_name)); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := c.Bind(data); err != nil {
		return err
	}

	err := data.Update(tx)
	if err != nil {
		c.Set("errors", err)
		c.Set("data", data)

		return c.Render(http.StatusUnprocessableEntity, r.HTML(tmpl_path+"/"+tmpl_edit))
	}

	c.Flash().Add("success", T.Translate(c, msg_name+".updated.success"))

	r_path := strings.ReplaceAll(tmpl_path, "_", "-")
	return c.Redirect(http.StatusSeeOther, "/"+r_path)
}

func DestroyHandler(c buffalo.Context, data AnyType, tmpl_path string, id_name string, msg_name string) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	if err := tx.Find(data, c.Param(id_name)); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(data); err != nil {
		return err
	}

	c.Flash().Add("success", T.Translate(c, msg_name+".destroyed.success"))

	r_path := strings.ReplaceAll(tmpl_path, "_", "-")
	return c.Redirect(http.StatusSeeOther, "/"+r_path)
}

func additionalActions(c buffalo.Context, data AnyType) {
	if v, ok := data.(*models.ExternalSystem); ok {
		actionsForExternalSystems(c, v)
	}
}

func actionsForExternalSystems(c buffalo.Context, data AnyType) {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return
	}

	ct, err := models.AllConnectionTypes(tx)
	if err != nil {
		return
	}

	c.Set("enums", ct)
}
