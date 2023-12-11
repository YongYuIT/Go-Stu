package client_test

import (
	"gRPC_Server/pb/apis"
	"gRPC_Server/pb/protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

const (
	address = "localhost:50051"
)

func TestGetFullName(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("conn error: %v", err)
	}
	defer conn.Close()
	nameServiceClient := apis.NewNameServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	fullNameResponse, err := nameServiceClient.GetFullNameFunc(ctx, &protos.GetFullNameRequest{Name: "hello"})
	if err != nil {
		log.Fatalf("GetFullNameRequest error: %v", err)
	}
	log.Printf("GetFullNameRequest: %s", fullNameResponse.FullName)
}
