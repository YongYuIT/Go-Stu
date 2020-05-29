package service

import (
	"fmt"
	pb "hello_grpc_stream/proto"
	"hello_grpc_stream/tools"
	"math/rand"
	"strings"
	"time"
)

type SchoolStuInfoServiceServerImp struct {
}

func (this *SchoolStuInfoServiceServerImp) GetStusByClassInfo(classinfo *pb.ClassInfo, stream pb.SchoolStuInfoService_GetStusByClassInfoServer) error {
	fmt.Println("get class info-->", classinfo.GradeName, classinfo.ClassName)
	if strings.EqualFold(classinfo.GradeName, "test_remote_err") {
		return fmt.Errorf("this is an funk remote err")
	}
	for i := 0; i < 20; i++ {
		info := pb.StuInfo{}
		if rand.Intn(1) == 0 {
			info.StuGender = "M"
		} else {
			info.StuGender = "F"
		}
		info.StuHomeAdd = tools.GetRandStr(10 + rand.Intn(10))
		info.StuName = tools.GetName()
		stream.Send(&info)
		time.Sleep(3 * time.Second)
	}
	return nil
}
