package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"markasbali_go_final_project/cli-service/controllers"
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
	app.Use(cors.New())
	baseAddress := fmt.Sprintf("%s:%s", os.Getenv("APP_CLI_SERVICE_URL"), os.Getenv("APP_CLI_SERVICE_PORT"))

	controllers.RouteCli(app)

	err := app.Listen(baseAddress)

	if err != nil {
		panic(err)
	}
}
