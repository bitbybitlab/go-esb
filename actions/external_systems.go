package actions

import (
	"goesb/models"

	"github.com/gobuffalo/buffalo"
)

type ExternalSystemsResource struct {
	buffalo.Resource
}

func (v ExternalSystemsResource) List(c buffalo.Context) error {
	return ListHandler(c, &models.ExternalSystems{}, "external_systems")
}

func (v ExternalSystemsResource) Show(c buffalo.Context) error {
	return ShowHandler(c, &models.ExternalSystem{}, "external_systems", "external_system_id")
}

func (v ExternalSystemsResource) New(c buffalo.Context) error {
	return NewHandler(c, &models.ExternalSystem{}, "external_systems")
}

func (v ExternalSystemsResource) Create(c buffalo.Context) error {
	return CreateHandler(c, &models.ExternalSystem{}, "external_systems", "external_system")
}

func (v ExternalSystemsResource) Edit(c buffalo.Context) error {
	return EditHandler(c, &models.ExternalSystem{}, "external_systems", "external_system_id")
}

func (v ExternalSystemsResource) Update(c buffalo.Context) error {
	return UpdateHandler(c, &models.ExternalSystem{}, "external_systems", "external_system_id", "external_system")
}

func (v ExternalSystemsResource) Destroy(c buffalo.Context) error {
	return DestroyHandler(c, &models.ExternalSystem{}, "external_systems", "external_system_id", "external_system")
}
