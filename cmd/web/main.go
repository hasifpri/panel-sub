package main

import (
	"panel-subs/infrastructure"
	infrastructureservice "panel-subs/infrastructure/service"
)

//	@title			PANEL SUB API
//	@version		1.0
//	@description	This is Api for province microservice.
//	@termsOfService	http://swagger.io/terms/

//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization

// @schemes	https
func main() {
	defer infrastructure.CloseAll(infrastructure.CTX)
	infrastructure.Initialize()
	infrastructureservice.Register()
	infrastructure.InitializeRoute()
	infrastructure.Listen()
}
