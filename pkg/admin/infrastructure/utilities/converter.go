package admininfrastructureutilities

import (
	admincoreentities "panel-subs/pkg/admin/core/entities"
	admininfrastructurerepositorymodel "panel-subs/pkg/admin/infrastructure/repository/model"
)

func ConvertFromEntitiesToModel(request admincoreentities.Admin) (model admininfrastructurerepositorymodel.Admin) {

	model.AdminID = request.ID
	model.Name = request.Name
	model.Email = request.Email

	return model
}
