package adminhandler

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
	applicationservice "panel-subs/application/service"
	infrastructureutilities "panel-subs/infrastructure/utilities"
	adminapplicationrequest "panel-subs/pkg/admin/application/request"
	adminapplicationresponse "panel-subs/pkg/admin/application/response"
	admininfrastructurecontainer "panel-subs/pkg/admin/infrastructure/container"
	"time"
)

type adminHandler struct {
	validator          *validator.Validate
	insertAdminUseCase applicationservice.IUseCase[adminapplicationrequest.CreateAdminInfo, adminapplicationresponse.CreateAdminResponse]
}

func NewAdminHandler(
	validator *validator.Validate,
	insertAdminUseCase applicationservice.IUseCase[adminapplicationrequest.CreateAdminInfo, adminapplicationresponse.CreateAdminResponse],
) *adminHandler {
	return &adminHandler{
		validator:          validator,
		insertAdminUseCase: insertAdminUseCase,
	}
}

// Insert
//
//	@Summary		Insert a product.
//	@Description	Insert a product.
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			data	body		adminapplicationrequest.CreateAdminInfo		true	"Insert Admin Request Parameter"
//	@Success		201		{object}	adminapplicationresponse.ApiResponseInsert	"Result"
//	@Failure		400		{object}	adminapplicationresponse.ApiResponseInsert	"Result"
//	@Failure		422		{object}	adminapplicationresponse.ApiResponseInsert	"Result"
//	@Router			/admin [post]
//	@Security		Bearer
func (handler *adminHandler) Insert(fiberCtx *fiber.Ctx) error {

	var timeIn = time.Now()
	// Use the global TracerProvider.
	tr := otel.Tracer("handler-productHandler")
	correlationID := uuid.NewString()
	ctx, span := tr.Start(fiberCtx.UserContext(), "Insert", oteltrace.WithAttributes(attribute.String("CorrelationID", correlationID)))
	defer span.End()

	admininfrastructurecontainer.LOGGER.Info("Insert()", "00000000-0000-0000-0000-000000000000", "Start", "Info", "")

	requestData := adminapplicationrequest.CreateAdminInfo{}
	ctx = context.WithValue(ctx, "CorrelationID", correlationID)

	if err := fiberCtx.BodyParser(&requestData); err != nil {
		admininfrastructurecontainer.LOGGER.Error("Insert()", correlationID, "BodyParser()", "Error", err)
		response := adminapplicationresponse.ApiResponseInsert{}
		response.CorrelationID = correlationID
		response.Success = false
		response.Error = err.Error()
		response.Tin = timeIn
		response.Tout = time.Now()
		response.Data = nil

		return fiberCtx.Status(fiber.StatusBadRequest).JSON(&response)
	}

	if err := handler.validator.Struct(requestData); err != nil {
		admininfrastructurecontainer.LOGGER.Error("Insert()", correlationID, "validator.Struct()", "Error", err)
		response := adminapplicationresponse.ApiResponseInsert{}
		response.CorrelationID = correlationID
		response.Success = false
		response.Error = err.Error()
		response.Tin = timeIn
		response.Tout = time.Now()
		response.Latency = infrastructureutilities.GetLatency(timeIn)
		response.Data = nil

		return fiberCtx.Status(fiber.StatusBadRequest).JSON(&response)
	}

	useCase := handler.insertAdminUseCase
	output, exception := useCase.Execute(ctx, requestData)
	if exception != nil {
		admininfrastructurecontainer.LOGGER.Error("Insert()", correlationID, "Finish", "Result", exception.Error)
		errResult := exception.GetError()

		response := adminapplicationresponse.ApiResponseInsert{}
		response.CorrelationID = correlationID
		response.Success = false
		response.Error = *errResult
		response.Tin = timeIn
		response.Tout = time.Now()
		response.Latency = infrastructureutilities.GetLatency(timeIn)
		response.Data = nil

		return fiberCtx.Status(exception.GetHttpCode()).JSON(&response)

	} else {

		admininfrastructurecontainer.LOGGER.Info("Insert()", correlationID, "Finish", "Result", output)

		response := adminapplicationresponse.ApiResponseInsert{}
		response.CorrelationID = correlationID
		response.Success = true
		response.Error = ""
		response.Tin = timeIn
		response.Tout = time.Now()
		response.Latency = infrastructureutilities.GetLatency(timeIn)
		response.Data = &output

		return fiberCtx.Status(fiber.StatusCreated).JSON(&response)
	}
}
