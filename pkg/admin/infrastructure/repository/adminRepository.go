package admininfrastructurerepository

import (
	"context"
	"fmt"
	"github.com/bytesaddict/dancok"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
	applicationservice "panel-subs/application/service"
	corerepository "panel-subs/core/repository"
	infrastructureutilities "panel-subs/infrastructure/utilities"
	admininfrastructurerepositorymodel "panel-subs/pkg/admin/infrastructure/repository/model"
)

type AdminRepository struct {
	logger         applicationservice.ILogger
	queryGenerator *infrastructureutilities.SqlGenerator
}

func NewAdminRepository(logger applicationservice.ILogger) *AdminRepository {
	return &AdminRepository{
		logger:         logger,
		queryGenerator: infrastructureutilities.NewSqlGenerator("admin", "created_at"),
	}
}

func (rep *AdminRepository) Select(db *gorm.DB, ctx context.Context, param corerepository.QueryInfo) (result []admininfrastructurerepositorymodel.Admin, totalItems int32, page int32, pageSize int32, errres error) {

	// Global Tracer
	tr := otel.Tracer("infrastructurerepository-adminRepository")
	ctx, span := tr.Start(ctx, "Select")
	defer span.End()

	correlationID := ""
	if v := ctx.Value("CorrelationID"); v != nil {
		correlationID = v.(string)
	}

	defer func() {
		if err := recover(); err != nil {
			rep.logger.Info("ProductRepository.Select()", correlationID, "Panic Recover()", "Error", err)

			result = make([]admininfrastructurerepositorymodel.Admin, 0)
			errres = fmt.Errorf("panic attact with reason %v ", err)
		}
	}()

	var sortDefault dancok.SortDescriptor
	sortDefault.FieldName = rep.queryGenerator.DefaultFieldForSort
	sortDefault.SortDirection = dancok.Descending

	if len(param.SelectParameter.SortDescriptors) == 0 {
		param.SelectParameter.SortDescriptors = append(param.SelectParameter.SortDescriptors, sortDefault)
	}

	sqlQuery := rep.queryGenerator.Generate(param.SelectParameter)

	rep.logger.Info("AdminRepository.Select()", correlationID, "Retrieve From DB Query", "Query", sqlQuery)

	queryResult := []admininfrastructurerepositorymodel.Admin{}

	err := db.WithContext(ctx).Raw(sqlQuery).Scan(&queryResult).Scan(&queryResult).Error
	if err != nil {
		rep.logger.Error("AdminRepository.Select()", correlationID, "Query Error", "Error", err)
	}

	result = queryResult

	rep.logger.Info("AdminRepository.Select()", correlationID, "Retrieve From DB Query", "Result", queryResult)

	return
}

func (rep *AdminRepository) Find(db *gorm.DB, ctx context.Context, ID int64) (result admininfrastructurerepositorymodel.Admin, errres error) {

	// Use Global Tracer
	tr := otel.Tracer("infrastructurerepository-adminRepository")
	ctx, span := tr.Start(ctx, "Find")
	defer span.End()

	correlationID := ""
	if v := ctx.Value("CorrelationID"); v != nil {
		correlationID = v.(string)
	}

	defer func() {
		if err := recover(); err != nil {
			rep.logger.Info("ProductRepository.Find()", correlationID, "Panic Recover()", "Error", err)

			errres = fmt.Errorf("panic attact with reason %v ", err)
		}
	}()

	rep.logger.Info("AdminRepository.Find()", correlationID, "Find Param", "Param", ID)

	queryResults := admininfrastructurerepositorymodel.Admin{}
	query := db.WithContext(ctx).Where("admin_id=?", ID)

	if err := query.First(&queryResults).Error; err != nil {
		rep.logger.Error("AdminRepository.Find()", correlationID, "Query Error", "Error", err)
		errres = err
		return
	}

	result = queryResults

	rep.logger.Info("AdminRepository.Find()", correlationID, "Find Result", "Result", result)

	return
}

func (rep *AdminRepository) Create(tx *gorm.DB, ctx context.Context, entity admininfrastructurerepositorymodel.Admin) (result admininfrastructurerepositorymodel.Admin, errres error) {

	// Global Tracer
	tr := otel.Tracer("infrastructurerepository-AdminRepository")
	ctx, span := tr.Start(ctx, "Insert")
	defer span.End()

	correlationID := ""
	if v := ctx.Value("CorrelationID"); v != nil {
		correlationID = v.(string)
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			rep.logger.Info("ProductRepository.Create()", correlationID, "Panic Recover()", "Error", err)

			errres = fmt.Errorf("panic attact with reason %v ", err)
		}
	}()

	rep.logger.Info("AdminRepository.Create()", correlationID, "Create Param", "Param", entity)

	if err := tx.Create(&entity).Error; err != nil {
		tx.Rollback()
		rep.logger.Error("AdminRepository.Create()", correlationID, "Save()", "Error", err)
		errres = err
	}

	rep.logger.Info("AdminRepository.Create()", correlationID, "Save()", "Result", entity)

	result = entity
	return
}

func (rep *AdminRepository) Update(tx *gorm.DB, ctx context.Context, entity admininfrastructurerepositorymodel.Admin) (result admininfrastructurerepositorymodel.Admin, errres error) {
	// Use the global TracerProvider.
	tr := otel.Tracer("infrastructurerepository-AdminRepository")
	ctx, span := tr.Start(ctx, "Update")
	defer span.End()

	correlationID := ""
	if v := ctx.Value("CorrelationID"); v != nil {
		correlationID = v.(string)
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			rep.logger.Info("AdminRepository.Update()", correlationID, "Panic Recover()", "Error", err)

			errres = fmt.Errorf("panic attact with reason %v ", err)
		}
	}()

	rep.logger.Info("AdminRepository.Update()", correlationID, "Update Param", "Param", entity)

	if err := tx.WithContext(ctx).Session(&gorm.Session{FullSaveAssociations: true}).Save(&entity).Error; err != nil {
		tx.Rollback()
		rep.logger.Error("ProductRepository.Update()", correlationID, "Save()", "Error", err)
		errres = err
	}

	rep.logger.Info("AdminRepository.Update()", correlationID, "Save()", "Result", "OK")

	result = entity

	return
}

func (rep *AdminRepository) Delete(tx *gorm.DB, ctx context.Context, ID int64) (errres error) {

	// Use the global TracerProvider.
	tr := otel.Tracer("infrastructurerepository-AdminRepository")
	ctx, span := tr.Start(ctx, "Delete")
	defer span.End()

	correlationID := ""
	if v := ctx.Value("CorrelationID"); v != nil {
		correlationID = v.(string)
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			rep.logger.Info("AdminRepository.Delete()", correlationID, "Panic Recover()", "Error", err)

			errres = fmt.Errorf("panic attact with reason %v ", err)
		}
	}()

	rep.logger.Info("AdminRepository.Delete()", correlationID, "Delete Param", "Param", ID)

	model := new(admininfrastructurerepositorymodel.Admin)

	if err := tx.WithContext(ctx).Where("admin_id = ?", ID).Delete(&model).Error; err != nil {
		tx.Rollback()
		rep.logger.Error("AdminRepository.Delete()", correlationID, "Delete()", "Error", err)
		errres = err
	}

	rep.logger.Info("AdminRepository.Delete()", correlationID, "Delete()", "Result", "OK")

	return
}
