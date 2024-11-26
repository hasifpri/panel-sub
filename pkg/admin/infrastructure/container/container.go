package admininfrastructurecontainer

import (
	applicationservice "panel-subs/application/service"
	corerepository "panel-subs/core/repository"
	admininfrastructurerepositorymodel "panel-subs/pkg/admin/infrastructure/repository/model"
)

var LOGGER applicationservice.ILogger

var ADMINREPOSITORY corerepository.IRepository[admininfrastructurerepositorymodel.Admin]
