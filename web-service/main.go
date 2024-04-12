package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"markasbali_go_final_project/web-service/configs"
	common "markasbali_go_final_project/web-service/modules/backup-restore/controllers"
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("cannot find .env file")
	}
	configs.OpenDB(false)
}
func main() {
	Init()
	app := fiber.New(fiber.Config{BodyLimit: 10 * 1024 * 1024})
	common.RouteWeb(app)

	baseAddress := fmt.Sprintf("%s:%s", configs.GetFiberHttpHost(), configs.GetFiberHttpPort())

	err := app.Listen(baseAddress)
	if err != nil {
		panic(err)
	}
}
