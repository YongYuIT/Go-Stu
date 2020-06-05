package service

import (
	"context"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"hello_hdf/tools"
	"os"
)

func RegisterNode(ctx context.Context, port int) error {
	cid := uuid.NewV4().String()
	err := createCidFile(cid)
	if err != nil {
		fmt.Println("get uuid err-->", err)
		return err
	}
	return tools.CreateNodeInfo(cid, ctx, port)
}
func createCidFile(cid string) error {
	cid_file_info, err := os.Stat(cid)
	if os.IsExist(err) && !cid_file_info.IsDir() {
		return err
	}
	cid_file, err := os.Create(cid)
	if err != nil {
		fmt.Println("create cid file err-->", err)
		return err
	}
	defer cid_file.Close()
	cid_file.WriteString(cid)
	return nil
}
