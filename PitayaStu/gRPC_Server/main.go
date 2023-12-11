package main

import (
	"context"
	"gRPC_Server/pb/apis"
	"gRPC_Server/pb/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":50051"
)

type MyUserInfoServer struct {
	apis.NameServiceServer
}

func (MyUserInfoServer) GetFullNameFunc(ctx context.Context, request *protos.GetFullNameRequest) (*protos.GetFullNameResponse, error) {
	return &protos.GetFullNameResponse{FullName: "full-" + request.Name}, nil
}
func (MyUserInfoServer) mustEmbedUnimplementedNameServiceServer() {}

func main() {
	tcp_listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	newServer := grpc.NewServer()
	apis.RegisterNameServiceServer(newServer, MyUserInfoServer{})
	reflection.Register(newServer)
	server_err := newServer.Serve(tcp_listener)
	if server_err != nil {
		log.Fatalf("failed to serve: %v", server_err)
	}
	log.Println("stop main proc")
}
