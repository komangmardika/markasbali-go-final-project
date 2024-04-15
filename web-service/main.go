package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"markasbali_go_final_project/proto"
	"markasbali_go_final_project/web-service/configs"
	"markasbali_go_final_project/web-service/controllers"
	"net"
	"os"
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("cannot find .env file")
	}
	configs.OpenDB(false)
}

func RunGRPCServer() {
	grpcServer := grpc.NewServer()

	proto.RegisterFileServiceServer(grpcServer, &controllers.FileService{})
	proto.RegisterRestoreServiceServer(grpcServer, &controllers.RestoreService{})

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	Init()
	app := fiber.New()
	app.Use(cors.New())
	controllers.RouteWeb(app)
	baseAddress := fmt.Sprintf("%s:%s", os.Getenv("APP_WEB_SERVICE_URL"), os.Getenv("APP_WEB_SERVICE_PORT"))
	go RunGRPCServer()
	err := app.Listen(baseAddress)

	if err != nil {
		panic(err)
	}
}
