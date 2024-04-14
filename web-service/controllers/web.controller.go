package controllers

import (
	"github.com/gofiber/fiber/v2"
	"markasbali_go_final_project/web-service/middlewares"
	"markasbali_go_final_project/web-service/services"
	"strconv"
)

func RouteWeb(app *fiber.App) {
	webGroup := app.Group("/web", middlewares.CheckAuth)
	webGroup.Get("/", GetAllDatabasesWithLatestBackup)
	webGroup.Get("/:db_name", GetOneDatabaseWithLatestHistory)
	webGroup.Post("/:db_name", PostUploadZipFile)
	webGroup.Get("/:id_file/download", GetDownloadLatestBackedUpByDatabase)
}

func GetAllDatabasesWithLatestBackup(ctx *fiber.Ctx) error {
	list, err := services.GetLatestBackedUpDatabaseList()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"message": "cannot get list",
			"error":   err,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(map[string]any{
		"data":    list,
		"message": "Success",
	})
}

func GetOneDatabaseWithLatestHistory(ctx *fiber.Ctx) error {
	dbName := ctx.Params("db_name")
	history, err := services.GetOneDatabaseWithHistory(dbName)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"message": "cannot get data",
			"error":   err,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(map[string]any{
		"data":    history,
		"message": "Success",
	})
}

func GetDownloadLatestBackedUpByDatabase(ctx *fiber.Ctx) error {

	fileId := ctx.Params("id_file")
	f, _ := strconv.ParseUint(fileId, 10, 64)
	fileContent, err := services.GetDownloadLatestBackedUpByDatabase(uint(f))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"message": "cannot get data",
			"error":   err,
		})
	}

	ctx.Set(fiber.HeaderContentType, "application/octet-stream")
	ctx.Set(fiber.HeaderContentDisposition, "attachment; filename=example.txt")

	return ctx.Send(fileContent)
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

	db, dbBackup, err := services.BackupSqlFile(ctx, file, dbName)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"message": "failed to save file",
			"error":   err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(map[string]any{
		"data": map[string]any{
			"id":            dbBackup.ID,
			"database_name": db.DatabaseName,
			"file_name":     dbBackup.FileName,
			"timestamp":     dbBackup.UpdatedAt,
		},
		"message": "Success",
	})
}
