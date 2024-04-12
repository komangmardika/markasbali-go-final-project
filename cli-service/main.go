package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	backup "markasbali_go_final_project/cli-service/modules/backup/controllers"
	common "markasbali_go_final_project/cli-service/modules/common/controllers"
	reset "markasbali_go_final_project/cli-service/modules/reset/controllers"
	"os"
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("cannot find .env file")
	}
}

func main() {
	Init()
	app := fiber.New()
	baseAddress := fmt.Sprintf("%s:%s", os.Getenv("APP_CLI_SERVICE_URL"), os.Getenv("APP_CLI_SERVICE_PORT"))

	reset.RouteReset(app)
	backup.RouteBackup(app)
	common.RouteCommon(app)

	err := app.Listen(baseAddress)

	if err != nil {
		panic(err)
	}
}
