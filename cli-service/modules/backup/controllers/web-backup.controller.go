package controllers

import (
	"github.com/gofiber/fiber/v2"
	"markasbali_go_final_project/cli-service/middlewares"
	"markasbali_go_final_project/cli-service/modules/backup/services"
)

func RouteBackup(app *fiber.App) {
	backupGroup := app.Group("/cli/backup", middlewares.CheckAuth)
	backupGroup.Get("/", GetBackupDatabases)
}

func GetBackupDatabases(ctx *fiber.Ctx) error {
	err := services.BackupDb()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"message": "cannot do import books",
			"error":   err,
		})
	}

	// err = common.RemoveFilesInTmpFolder()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"message": "cannot delete files inside tmp folder",
			"error":   err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(map[string]any{
		"message": "Backup done",
	})
}
