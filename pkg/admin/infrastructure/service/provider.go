package admininfrastructureservice

import (
	"github.com/google/wire"
	corerepository "panel-subs/core/repository"
	admininfrastructurecontainer "panel-subs/pkg/admin/infrastructure/container"
	admininfrastructurerepository "panel-subs/pkg/admin/infrastructure/repository"
	admininfrastructurerepositorymodel "panel-subs/pkg/admin/infrastructure/repository/model"

	"sync"
)

var repoOnce sync.Once

var repo *admininfrastructurerepository.AdminRepository

var ProviderSet wire.ProviderSet = wire.NewSet(
	ProvideRepository,
	wire.Bind(new(corerepository.IRepository[admininfrastructurerepositorymodel.Admin]), new(*admininfrastructurerepository.AdminRepository)),
)

func ProvideRepository() *admininfrastructurerepository.AdminRepository {
	repoOnce.Do(func() {
		repo = admininfrastructurerepository.NewAdminRepository(admininfrastructurecontainer.LOGGER)
	})

	return repo
}
