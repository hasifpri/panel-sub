package infrastructurehandler

import (
	"context"
	"errors"
	infrastructurepinger "panel-subs/infrastructure/pinger"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	redisPinger infrastructurepinger.IPinger
}

func NewHandler(redisPinger infrastructurepinger.IPinger) *handler {
	return &handler{redisPinger}
}

func (handler *handler) pingAll(ctx context.Context) error {
	redisStatus := make(chan string)

	//redis
	go func() {
		err := handler.redisPinger.Ping(ctx)
		if err != nil {
			redisStatus <- "Unable to connect to redis : " + err.Error()
		} else {
			redisStatus <- ""
		}
	}()

	redisResult := <-redisStatus
	if redisResult != "" {
		return errors.New(redisResult)
	}

	return nil
}

// Liveness
//
//	@Summary		Show the status of server.
//	@Description	get the liveness status of server.
//	@Tags			Health
//	@Accept			*/*
//	@Produce		plain
//	@Success		200	{string}	string	"OK"
//	@Router			/health/liveness [get]
func (handler *handler) Liveness(fiberCtx *fiber.Ctx) error {
	return fiberCtx.Status(fiber.StatusOK).SendString("OK")
}

// Readiness
//
//	@Summary		Show the readiness status of server.
//	@Description	get the readiness status of server.
//	@Tags			Health
//	@Accept			*/*
//	@Produce		plain
//	@Success		200	{string}	string	"OK"
//	@Router			/health/readiness [get]
func (handler *handler) Readiness(fiberCtx *fiber.Ctx) error {
	statusCode := 200
	statusMessage := "OK"

	err := handler.pingAll(fiberCtx.UserContext())

	if err != nil {
		statusCode = 422
		statusMessage = err.Error()
	}

	return fiberCtx.Status(statusCode).SendString(statusMessage)
}
