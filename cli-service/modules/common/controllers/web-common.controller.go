package controllers

import (
	"github.com/gofiber/fiber/v2"
	"markasbali_go_final_project/cli-service/middlewares"
	"markasbali_go_final_project/cli-service/modules/common/services"
)

func RouteCommon(app *fiber.App) {
	resetGroup := app.Group("/cli/backup-restore", middlewares.CheckAuth)
	resetGroup.Get("/db-list", GetListDatabasesFromJson)
}

func GetListDatabasesFromJson(ctx *fiber.Ctx) error {
	configs, _ := services.ReadDatabasesJson()
	return ctx.Status(fiber.StatusOK).JSON(map[string]any{
		"data":    configs,
		"message": "Success",
	})
}
