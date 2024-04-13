package controllers

import (
	"context"
	"markasbali_go_final_project/proto"
)

type RestoreService struct {
	proto.RestoreServiceServer
}

type FileService struct {
	proto.FileServiceServer
}

func GRPCAllDatabasesWithLatestBackup(ctx context.Context, req *proto.RestoreRequest) error {
	return nil
}

func GRPCOneDatabaseWithLatestHistory(ctx context.Context, req *proto.RestoreRequest) error {
	return nil
}

func GRPCDownloadLatestBackedUpByDatabase(ctx context.Context, req *proto.FileRequest) error {
	return nil
}

func (s *RestoreService) Restore(ctx context.Context, req *proto.RestoreRequest) (*proto.RestoreResponses, error) {

	response := ""

	return &proto.RestoreResponses{
		RestoreResponse: response,
	}, nil
}

func (s *FileService) GetFile(ctx context.Context, req *proto.FileRequest) (*proto.FileResponse, error) {

	var fileContent []byte

	response := &proto.FileResponse{
		FileContent: fileContent,
	}

	return response, nil
}
