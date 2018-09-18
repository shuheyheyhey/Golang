package service

import (
	"context"
	"errors"
	"os"

	pb "../protocol"
)

type MyCatService struct {
}

func (s *MyCatService) GetMyCat(ctx context.Context, message *pb.GetMyCatMessage) (*pb.MyCatResponse, error) {
	switch message.TargetCat {
	case "tama":
		//たまはメインクーン
		return &pb.MyCatResponse{
			Name: "tama",
			Kind: "mainecoon",
		}, nil
	case "mike":
		//ミケはノルウェージャンフォレストキャット
		return &pb.MyCatResponse{
			Name: "mike",
			Kind: "Norwegian Forest Cat",
		}, nil
	}

	return nil, errors.New("Not Found YourCat")
}

type ServerGRPC struct {
}

func (s *ServerGRPC) Upload(ctx context.Context, message *pb.UploadFileMessage) (*pb.UploadFileResponce, error) {
	file, err := os.Create("/Users/shuhei.yukawa/Desktop/Test.png")
	if err != nil {
		return nil, errors.New("Error Could not create file.")
	}

	file.Write(message.Content)
	return &pb.UploadFileResponce{
		Code:    1,
		Message: "OK",
	}, nil
}
