package controllers

import (
	"context"
	"encoding/json"
	"markasbali_go_final_project/proto"
	"markasbali_go_final_project/web-service/services"
)

type RestoreService struct {
	proto.RestoreServiceServer
}

type FileService struct {
	proto.FileServiceServer
}

func (s *RestoreService) Restore(ctx context.Context, req *proto.RestoreRequest) (*proto.RestoreResponses, error) {

	var response string
	switch req.Module {
	case "restore":
		{
			switch req.Process {
			case "one-db-latest-history":
				{
					data, err := services.GetOneDatabaseWithHistory(req.DbName)
					if err != nil {
						return &proto.RestoreResponses{
							RestoreResponse: "",
						}, err
					}
					marshal, err := json.Marshal(data)
					if err != nil {
						return &proto.RestoreResponses{
							RestoreResponse: "",
						}, err
					}
					response = string(marshal)

					break
				}
			case "all-db-latest-history":
				{
					data, err := services.GetLatestBackedUpDatabaseList()
					if err != nil {
						return &proto.RestoreResponses{
							RestoreResponse: "",
						}, err
					}
					marshal, err := json.Marshal(data)
					if err != nil {
						return &proto.RestoreResponses{
							RestoreResponse: "",
						}, err
					}
					response = string(marshal)

					break
				}

			}

		}
	}

	return &proto.RestoreResponses{
		RestoreResponse: response,
	}, nil
}

func (s *FileService) GetFile(ctx context.Context, req *proto.FileRequest) (*proto.FileResponse, error) {

	fileId := req.GetFileId()
	file, err := services.GetDownloadLatestBackedUpByDatabase(uint(fileId))
	if err != nil {
		return nil, err
	}

	response := &proto.FileResponse{
		FileContent: file,
	}

	return response, nil
}
