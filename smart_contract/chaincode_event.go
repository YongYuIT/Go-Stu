package main

import (
	"fmt"
	//"github.com/hyperledger/fabric/core/chaincode/shim" //1.4
	"github.com/hyperledger/fabric-chaincode-go/shim" //2.0 change
	"strings"
)
//import pb "github.com/hyperledger/fabric/protos/peer" //1.4
import pb "github.com/hyperledger/fabric-protos-go/peer" //2.0 change

type ChainCodeContract struct {
}

func (this *ChainCodeContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	var response pb.Response
	response.Status = shim.OK
	response.Payload = []byte("init is called")
	return response
}

func (this *ChainCodeContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	var response pb.Response
	response.Status = shim.ERROR
	response.Payload = []byte("unknown error")

	func_name, params := stub.GetFunctionAndParameters()
	if strings.Compare("test_env", func_name) == 0 {
		err := stub.SetEvent("test_event_name", []byte("this is test event message-->"+params[0]))
		if err == nil {
			response.Status = shim.OK
			response.Payload = []byte("test event success")
		}
		return response
	}
	return response
}

func main() {
	err := shim.Start(new(ChainCodeContract))
	if err != nil {
		fmt.Println("chaincode start error --> " + err.Error())
	}

}
