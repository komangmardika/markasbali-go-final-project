package controllers

import (
	"github.com/gofiber/fiber/v2"
	"markasbali_go_final_project/web-service/middlewares"
	"markasbali_go_final_project/web-service/modules/backup-restore/services"
)

func RouteWeb(app *fiber.App) {
	webGroup := app.Group("/web", middlewares.CheckAuth)
	webGroup.Get("/", GetLatestBackedUpDatabaseList)
	webGroup.Get("/:db_name", GetOneDatabaseHistory)
	webGroup.Post("/:db_name", PostUploadZipFile)
	webGroup.Get("/:id_file/download", GetDownloadLatestBackedUpByDatabase)
}

func GetLatestBackedUpDatabaseList(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(map[string]any{
		"data":    nil,
		"message": "Success",
	})
}

func GetOneDatabaseHistory(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(map[string]any{
		"data":    nil,
		"message": "Success",
	})
}

func GetDownloadLatestBackedUpByDatabase(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(map[string]any{
		"data":    nil,
		"message": "Success",
	})
}

func PostUploadZipFile(ctx *fiber.Ctx) error {

	dbName := ctx.Params("db_name")

	if dbName == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"message": "parameter cannot be null",
		})
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"message": "no file attachment found",
			"error":   err,
		})
	}

	err = services.BackupSqlFile(ctx, file, dbName)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"message": "failed to save file",
			"error":   err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(map[string]any{
		"data":    nil,
		"message": "Success",
	})
}
