package main

import (
	"context"
	"encoding/json"
	"final-project/kelas-beta-golang/config"
	"final-project/kelas-beta-golang/proto"
	"final-project/kelas-beta-golang/utils"
	"fmt"
	"net"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type GreetService struct {
	proto.GreetServiceServer
}

func Init()  {
	err := godotenv.Load(".env")
	if err != nil{
		fmt.Println("env not found, using global .env")
	}
	config.OpenDB()
}

func (s *GreetService) Greet(ctx context.Context, req *proto.GreetRequest) (*proto.GreetResponses, error) {
	Init()

	autoResData, err := utils.GetDistinct()
	if err != nil {
		logrus.Error("Error on get database distinct list: ", err.Error())
	}

	responseData := make([]map[string]interface{}, len(autoResData))

	for i, row := range autoResData {

		dataFetch, err := utils.GetLatestByDBName(row)

		if err != nil {
			logrus.Error("Error on get database latest backup name list: ", err.Error())
		}

		responseData[i] = map[string]interface{}{
			"database_name": dataFetch.Nama_Database,
			"latest_backup": map[string]interface{}{
				"id":         dataFetch.Id,
				"file_name":  dataFetch.Nama_File_Backup,
				"timestamp":  dataFetch.CreatedAt,
			},
		}
    }

	// Mengonversi responseData menjadi string JSON
    jsonString, err := json.MarshalIndent(responseData, "", "  ")
    if err != nil {
        fmt.Println("Error:", err)
    }


	return &proto.GreetResponses{
		Greetresponse: string(jsonString),
	}, nil

}

func main() {
	fmt.Println("Memulai gRPC Server ...")
	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		logrus.Fatalf("Gagal menginisialisasi server, error : %s", err.Error())
	}

	grpcServer := grpc.NewServer()
	proto.RegisterGreetServiceServer(grpcServer, &GreetService{})
	grpcServer.Serve(grpcListener)
}
