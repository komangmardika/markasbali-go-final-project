package main

import (
	"final-project/kelas-beta-golang/config"
	"final-project/kelas-beta-golang/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func InitEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Warn("Cannot load env file, using system env")
	}
}

func main() {
	InitEnv()
	config.OpenDB()

	app := fiber.New()

	controllers.Route(app)

	err := app.Listen(":3000")
	if err != nil {
		logrus.Fatal(
			"Error on running fiber, ",
			err.Error())
	}

}