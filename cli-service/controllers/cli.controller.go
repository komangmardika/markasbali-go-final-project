package controllers

import (
	"github.com/gofiber/fiber/v2"
	"markasbali_go_final_project/cli-service/middlewares"
	"markasbali_go_final_project/cli-service/services"
)

func RouteCli(app *fiber.App) {
	cliGroup := app.Group("/cli", middlewares.CheckAuth)
	cliGroup.Get("/backup", GetBackupDatabases)
	cliGroup.Get("/restore", GetRestoreDatabases)
	cliGroup.Get("/common/db-list", GetListDatabasesFromJson)
	cliGroup.Get("/reset", GetResetDatabases)
	cliGroup.Get("/reset/seed", GetSeedDatabases)
}

func GetBackupDatabases(ctx *fiber.Ctx) error {
	err := services.BackupDb()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"message": "cannot do import books",
			"error":   err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(map[string]any{
		"message": "databases has been backed up successfully",
	})
}

func GetListDatabasesFromJson(ctx *fiber.Ctx) error {
	configs, _ := services.ReadDatabasesJson()
	return ctx.Status(fiber.StatusOK).JSON(map[string]any{
		"data":    configs,
		"message": "Success",
	})
}

func GetResetDatabases(ctx *fiber.Ctx) error {

	configs, err := services.ReadDatabasesJson()

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"message": "error opening json file",
			"error":   err,
		})
	}

	// loop for every connection send it to resetDb function
	for _, config := range configs {
		err = services.ResetDb(config)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
				"message": "error reset database",
				"error":   err,
			})
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(map[string]any{
		"data":    configs,
		"message": "databases has been reset successfully",
	})
}

func GetRestoreDatabases(ctx *fiber.Ctx) error {

	err := services.RestoreDb()

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"message": "error restoring database",
			"error":   err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(map[string]any{
		"message": "databases has been restored successfully",
	})
}

func GetSeedDatabases(ctx *fiber.Ctx) error {
	configs, _ := services.ReadDatabasesJson()

	// loop for every connection send it to resetDb function
	for _, config := range configs {
		services.OpenDB(false, config)

		err := services.AutoMigrate(services.Mysql.DB)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
				"message": "cannot do auto migration",
				"error":   err,
			})
		}
		_ = services.SendSeedingProcessToWebSocketServer(config.DatabaseName, "Books")
		err = services.ImportBook()
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
				"message": "cannot do import books",
				"error":   err,
			})
		}
		_ = services.SendSeedingProcessToWebSocketServer(config.DatabaseName, "Cars")
		err = services.ImportCar()
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
				"message": "cannot do import books",
				"error":   err,
			})
		}

		err = services.CloseDb(services.Mysql.DB)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
				"message": "cannot close database",
				"error":   err,
			})
		}

	}
	return ctx.Status(fiber.StatusOK).JSON(map[string]any{
		"data":    configs,
		"message": "databases has been seeded successfully",
	})
}
