package actions

import (
	"goesb/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/pkg/errors"
)

type UsersResource struct {
	buffalo.Resource
}

func (v UsersResource) List(c buffalo.Context) error {
	return ListHandler(c, &models.Users{}, "users")
}

func (v UsersResource) Show(c buffalo.Context) error {
	return ShowHandler(c, &models.User{}, "users", "user_id")
}

func (v UsersResource) New(c buffalo.Context) error {
	return NewHandler(c, &models.User{}, "users")
}

func (v UsersResource) Create(c buffalo.Context) error {
	return CreateHandler(c, &models.User{}, "users", "user")
}

func (v UsersResource) Edit(c buffalo.Context) error {
	return EditHandler(c, &models.User{}, "users", "user_id")
}

func (v UsersResource) Update(c buffalo.Context) error {
	return UpdateHandler(c, &models.User{}, "users", "user_id", "user")
}

func (v UsersResource) Destroy(c buffalo.Context) error {
	return DestroyHandler(c, &models.User{}, "users", "user_id", "user")
}

func SetCurrentUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid != nil {
			u := &models.User{}
			tx := c.Value("tx").(*pop.Connection)
			err := tx.Find(u, uid)
			if err != nil {
				return errors.WithStack(err)
			}
			c.Set("current_user", u)
		}
		return next(c)
	}
}

func Authorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid == nil {
			c.Flash().Add("danger", "You must be authorized to see that page")
			return c.Redirect(302, "/signin")
		}
		return next(c)
	}
}
