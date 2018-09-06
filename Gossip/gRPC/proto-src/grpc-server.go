package proto_src

import (
	pb "./update-msg"
	"fmt"
	"strconv"
	"golang.org/x/net/context"
)

type Server struct {
	safe_add string
}

func NewServer(_self_add string) *Server {
	result := new(Server)
	result.safe_add = _self_add
	return result
}

func (*Server) DoUpdateAll(ser pb.Update_DoUpdateAllServer) error {
	inDatas, err := ser.Recv()
	if err != nil {
		return fmt.Errorf("error when recv: %s", err.Error())
	}
	fmt.Println(strconv.Itoa(len(inDatas.Datas)))

	return nil
}

func (this *Server) DoUpdate(ctx context.Context, in_data *pb.UpateData) (*pb.UpateData, error) {
	fmt.Printf("rec data --> from %s value %s\n", in_data.Key, in_data.Value)
	data := pb.UpateData{this.safe_add, "ping", -1, nil, nil, nil}
	return &data, nil
}
