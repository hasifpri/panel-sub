package admininfrastructureservice

import (
	infrastructurecontainer "panel-subs/infrastructure/container"
	admininfrastructurecontainer "panel-subs/pkg/admin/infrastructure/container"
)

func Register() {
	admininfrastructurecontainer.LOGGER = infrastructurecontainer.LOGGER
	admininfrastructurecontainer.ADMINREPOSITORY = WireRepository()
}
