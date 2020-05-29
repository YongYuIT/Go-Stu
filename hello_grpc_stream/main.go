package main

import (
	"fmt"
	"google.golang.org/grpc"
	pb "hello_grpc_stream/proto"
	"hello_grpc_stream/service"
	"net"
)

//go:generate protoc -I ./proto --go_out=plugins=grpc:./proto ./proto/HelloServerStream.proto

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8989")
	if err != nil {
		fmt.Println("get port err-->", err)
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterSchoolStuInfoServiceServer(grpcServer, &service.SchoolStuInfoServiceServerImp{})
	grpcServer.Serve(listener)
}
