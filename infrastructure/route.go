package infrastructure

import (
	"fmt"
	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	infrastructureconfiguration "panel-subs/infrastructure/configuration"
	infrastructurehandler "panel-subs/infrastructure/handler"
	infrastructurepinger "panel-subs/infrastructure/pinger"
	adminapplicationusecase "panel-subs/pkg/admin/application/usecase"
	admininfrastructurecontainer "panel-subs/pkg/admin/infrastructure/container"
	adminhandler "panel-subs/pkg/admin/infrastructure/handler"
)

var APP *fiber.App

func InitializeRoute() {
	fiberApp := fiber.New(fiber.Config{ReadBufferSize: 15000})

	fiberApp.Use(cors.New())

	fiberApp.Use(
		otelfiber.Middleware(otelfiber.WithSpanNameFormatter(func(ctx *fiber.Ctx) string {
			return fmt.Sprintf("%s %s", ctx.Method(), ctx.Path())
		})))

	APP = fiberApp

	// Pinger
	redisPinger := infrastructurepinger.NewRedisPinger(RDS)

	handler := infrastructurehandler.NewHandler(redisPinger)
	apiGroup := APP.Group("/api")

	// Health Handler
	healthGroupHandler := apiGroup.Group("/health")
	healthGroupHandler.Get("/liveness", handler.Liveness)
	healthGroupHandler.Get("/readiness", handler.Readiness)

	// modul admin
	insertAdminUseCase := adminapplicationusecase.NewInsertAdminUseCase(DB, admininfrastructurecontainer.LOGGER, admininfrastructurecontainer.ADMINREPOSITORY)

	adminAPI := apiGroup.Group("/admin")

	adminHandler := adminhandler.NewAdminHandler(VALIDATOR, insertAdminUseCase)
	adminAPI.Post("/", adminHandler.Insert)
}

func Listen() {
	APP.Listen(":" + infrastructureconfiguration.Port)
}
