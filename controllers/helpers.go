package controllers

import (
	pacrdv1alpha1 "github.com/armory-io/pacrd/api/v1alpha1"
	"github.com/armory/plank"
)

func toSpinApplication(app pacrdv1alpha1.Application) plank.Application {
	return plank.Application{
		Name:        app.Name,
		Email:       app.Spec.Email,
		Description: app.Spec.Description,
		User:        app.Spec.User,
		DataSources: plank.DataSourcesType{
			Enabled:  app.Spec.DataSources.Enabled,
			Disabled: app.Spec.DataSources.Disabled,
		},
		Permissions: plank.PermissionsType{
			Read:    app.Spec.Permissions.Read,
			Write:   app.Spec.Permissions.Write,
			Execute: app.Spec.Permissions.Execute,
		},
	}
}
