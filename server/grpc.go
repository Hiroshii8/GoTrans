package server

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/Hiroshii8/GoTrans/grpc/proto"
	"github.com/Hiroshii8/GoTrans/translate"
)

type grpcServer struct {
	port      string
	srv       *grpc.Server
	translate *translate.Translate
}

func InitGrpc(port string) *grpcServer {
	ts := translate.New()
	return &grpcServer{
		port:      port,
		translate: ts,
	}
}

func (s *grpcServer) Start() {
	address := "0.0.0.0" + s.port
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Something is wrong : ", err.Error())
	}
	log.Println("listen to ", address)
	newServer := grpc.NewServer()
	pb.RegisterTranslateServer(newServer, s)
	newServer.Serve(lis)
}

func (s *grpcServer) RequestGrpc(ctx context.Context, req *pb.GoTransReq) (resp *pb.GoTransResp, err error) {
	result, error := s.translate.Request(req.From, req.To, req.Text)
	success := true
	errorMessage := ""
	if error != nil {
		success = false
		errorMessage = error.Error()
	}
	return &pb.GoTransResp{
		IsSuccess: success,
		TranslateData: &pb.TranslateData{
			ErrorMessage:  errorMessage,
			TranslateText: result,
		},
	}, error
}
