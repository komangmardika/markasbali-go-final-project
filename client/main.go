package main

import (
	"context"
	"encoding/json"
	"final-project/kelas-beta-golang/proto"
	"fmt"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(":50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		logrus.Fatalf("Gagal menginisialisasi client, error : %s", err.Error())
	}
	defer conn.Close()

	greetService := proto.NewGreetServiceClient(conn)

	greetResponse, err := Greet(context.Background(),
		greetService)
	if err != nil {
		logrus.Fatalf("Terjadi error saat melakukan greetings, err:%s",
			err.Error())
	}
	fmt.Println(" ====================== ")
	
	// Mendeklarasikan variabel untuk menyimpan hasil unmarshaling
    var responseData []map[string]interface{}

    // Mengonversi string JSON kembali ke tipe []map[string]interface{}
    if err := json.Unmarshal([]byte(greetResponse), &responseData); err != nil {
        fmt.Println("Error:", err)
        return
    }

	fmt.Println(responseData)
}

func Greet(ctx context.Context,
	client proto.GreetServiceClient) (string, error) {
	res, err := client.Greet(ctx, &proto.GreetRequest{})

	if err != nil {
		return "", err
	}

	return res.Greetresponse, nil
}
