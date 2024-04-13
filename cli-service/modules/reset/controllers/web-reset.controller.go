package controllers

import (
	"github.com/gofiber/fiber/v2"
	"markasbali_go_final_project/cli-service/middlewares"
	"markasbali_go_final_project/cli-service/modules/common/services"
	reset "markasbali_go_final_project/cli-service/modules/reset/services"
)

func RouteReset(app *fiber.App) {
	resetGroup := app.Group("/cli/reset", middlewares.CheckAuth)
	resetGroup.Get("/", GetResetDatabases)
	resetGroup.Get("/seed", GetSeedDatabases)
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
		err = reset.ResetDb(config)
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

func GetSeedDatabases(ctx *fiber.Ctx) error {
	configs, _ := services.ReadDatabasesJson()

	// loop for every connection send it to resetDb function
	for _, config := range configs {
		reset.OpenDB(false, config)
		err := reset.AutoMigrate(reset.Mysql.DB)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
				"message": "cannot do auto migration",
				"error":   err,
			})
		}

		err = reset.ImportBook()
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
				"message": "cannot do import books",
				"error":   err,
			})
		}
		err = reset.ImportCar()
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
				"message": "cannot do import books",
				"error":   err,
			})
		}

		err = reset.CloseDb(reset.Mysql.DB)
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
