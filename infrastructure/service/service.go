package infrastructureservice

import (
	"panel-subs/infrastructure"
	infrastructurecontainer "panel-subs/infrastructure/container"
	infrastructurelogging "panel-subs/infrastructure/logging"
	admininfrastructureservice "panel-subs/pkg/admin/infrastructure/service"
)

func Register() {
	infrastructurecontainer.LOGGER = infrastructurelogging.NewLogger(infrastructure.ZAPLOGGER)

	admininfrastructureservice.Register()
}
