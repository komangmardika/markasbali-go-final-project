package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"markasbali_go_final_project/proto"
	"markasbali_go_final_project/web-service/configs"
	common "markasbali_go_final_project/web-service/modules/backup-restore/controllers"
	"net"
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

	proto.RegisterFileServiceServer(grpcServer, &common.FileService{})
	proto.RegisterRestoreServiceServer(grpcServer, &common.RestoreService{})

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
	app := fiber.New(fiber.Config{BodyLimit: 10 * 1024 * 1024})
	common.RouteWeb(app)

	baseAddress := fmt.Sprintf("%s:%s", configs.GetFiberHttpHost(), configs.GetFiberHttpPort())

	go RunGRPCServer()

	err := app.Listen(baseAddress)
	if err != nil {
		panic(err)
	}
}
