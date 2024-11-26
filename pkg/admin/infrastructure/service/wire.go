//go:build wireinject
// +build wireinject

package admininfrastructureservice

import (
	"github.com/google/wire"
	admininfrastructurerepository "panel-subs/pkg/admin/infrastructure/repository"
)

func WireRepository() *admininfrastructurerepository.AdminRepository {
	panic(wire.Build(ProviderSet))
}
