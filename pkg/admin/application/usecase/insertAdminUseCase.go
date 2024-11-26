package adminapplicationusecase

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
	applicationexception "panel-subs/application/exception"
	applicationservice "panel-subs/application/service"
	corerepository "panel-subs/core/repository"
	infrastructureutilities "panel-subs/infrastructure/utilities"
	adminapplicationrequest "panel-subs/pkg/admin/application/request"
	adminapplicationresponse "panel-subs/pkg/admin/application/response"
	admincoreentities "panel-subs/pkg/admin/core/entities"
	admininfrastructurerepositorymodel "panel-subs/pkg/admin/infrastructure/repository/model"
	admininfrastructureutilities "panel-subs/pkg/admin/infrastructure/utilities"
)

type insertAdminUseCase struct {
	db         *gorm.DB
	logger     applicationservice.ILogger
	repository corerepository.IRepository[admininfrastructurerepositorymodel.Admin]
}

func NewInsertAdminUseCase(
	db *gorm.DB,
	logger applicationservice.ILogger,
	repository corerepository.IRepository[admininfrastructurerepositorymodel.Admin],
) *insertAdminUseCase {
	return &insertAdminUseCase{
		db:         db,
		logger:     logger,
		repository: repository,
	}
}

func (usecase *insertAdminUseCase) Execute(ctx context.Context, input adminapplicationrequest.CreateAdminInfo) (output adminapplicationresponse.CreateAdminResponse, exc *applicationexception.Exception) {
	// Use the global TracerProvider.
	tr := otel.Tracer("applicationusecase-insertAdminUseCase")
	ctx, span := tr.Start(ctx, "Execute")
	defer span.End()

	tx := usecase.db.WithContext(ctx).Begin()

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			exc = applicationexception.Internal("nill pointer", fmt.Errorf(" failed panic nil pointer %v ", err))
			return
		}
	}()

	// Convert to entities
	var adminEntities admincoreentities.Admin
	adminEntities.Name = input.Name
	adminEntities.Email = input.Email
	adminEntities.Password = infrastructureutilities.HashPassword(input.Password)

	// Convert to Model
	model := admininfrastructureutilities.ConvertFromEntitiesToModel(adminEntities)

	entity, err := usecase.repository.Create(tx, ctx, model)
	if err != nil {
		tx.Rollback()
		usecase.logger.Error("insertAdminUseCase.Execute()", ctx.Value("CorrelationID").(string), "pendingRepository.Create()", "Error", err)
		exc = applicationexception.Internal("failed create", err)
		return output, exc
	}

	usecase.logger.Info("insertAdminUseCase.Execute()", ctx.Value("CorrelationID").(string), "repository.Create()", "Data", model)

	if tx.Commit().Error != nil {
		tx.Rollback()
		exc = applicationexception.Internal("failed commit", err)
		usecase.logger.Error("insertProductUseCase.Execute()", ctx.Value("CorrelationID").(string), "pendingRepository.Create()", "Error", err)
		return output, exc
	}

	// convert to response
	output.ID = infrastructureutilities.EncodeID(entity.AdminID)
	output.Name = entity.Name
	output.Email = entity.Email

	return output, nil
}
