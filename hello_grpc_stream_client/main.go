package main

import (
	"context"
	"fmt"
	pb "github.com/YongYuIT/Go-Stu/tree/master/hello_grpc_stream/proto"
	"google.golang.org/grpc"
	"io"
)

func main() {

	//不使用认证建立连接
	conn, err := grpc.Dial("0.0.0.0:8989", grpc.WithInsecure())
	if err != nil {
		fmt.Println("conn to server err-->", err)
		return
	}
	defer conn.Close()

	grpcClient := pb.NewSchoolStuInfoServiceClient(conn)
	req_test_err := pb.ClassInfo{}
	req_test_err.GradeName = "test_remote_err"
	req_test_err.ClassName = "test_remote_err"
	stream, err := grpcClient.GetStusByClassInfo(context.Background(), &req_test_err)
	if err != nil {
		fmt.Println("get err-->", err)
	} else {
		getMessageFromStream(stream)
	}

	req := pb.ClassInfo{}
	req.ClassName = "VIP C1"
	req.GradeName = "VVIP G1"
	stream, err = grpcClient.GetStusByClassInfo(context.Background(), &req)
	if err != nil {
		fmt.Println("get err-->", err)
	} else {
		getMessageFromStream(stream)
	}
}

func getMessageFromStream(stream pb.SchoolStuInfoService_GetStusByClassInfoClient) {
	for {
		resp, err := stream.Recv()
		//服务端数据发送完毕
		if err == io.EOF {
			fmt.Println("recv msg finish")
			break
		}
		if err != nil {
			fmt.Println("recv msr err-->", err)
			break
		}
		fmt.Printf("get stu %s, %s, at %s\n", resp.StuName, resp.GetStuGender(), resp.StuHomeAdd)
	}
}
