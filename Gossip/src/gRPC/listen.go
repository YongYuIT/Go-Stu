package gRPC

import (
	"fmt"
	"net"
	"google.golang.org/grpc"
	pb "./proto-src/update-msg"
	"google.golang.org/grpc/reflection"
	"./proto-src"
	"time"
	"log"
	"golang.org/x/net/context"
)

var IsOut bool = false
var Self_add string = ""
var Anchor_add string = ""

const (
	port = ":33333"
)

func Listen(interval int) {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterUpdateServer(s, proto_src.NewServer(Self_add))
	reflection.Register(s)

	go receive(interval)

	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve: %v", err)
		return
	}

}

func receive(interval int) {

	fmt.Printf("commp self %s, anch %s \n", Self_add, Anchor_add)

	if Self_add == Anchor_add {
		return
	}

	conn, err := grpc.Dial(Anchor_add+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	grpc := pb.NewUpdateClient(conn)

	for {
		if (IsOut) {
			fmt.Println("exit listen")
			break
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		data := pb.UpateData{Self_add, "ping", -1, struct{}{}, nil, 0}
		back_data, err := grpc.DoUpdate(ctx, &data)
		if err != nil {
			fmt.Printf("could not comm: %v \n", err)
		} else {
			fmt.Printf("back_data --> from %s value %s\n", back_data.Key, back_data.Value)
		}

		time.Sleep(time.Duration(interval) * time.Second)
	}
}
