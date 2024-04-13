package controllers

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"markasbali_go_final_project/proto"
)

func InitializeGRPCClient() {
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to initilize client, error : %s", err.Error())
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	proto.NewRestoreServiceClient(conn)
}

func Restore(ctx context.Context, client proto.RestoreServiceClient, process string, module string, dbName string) (string, error) {
	res, err := client.Restore(ctx, &proto.RestoreRequest{Module: module, Process: process, DbName: dbName})
	if err != nil {
		return "", err
	}

	return res.String(), err
}
