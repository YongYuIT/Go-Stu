package main

import (
	"encoding/json"
	"fmt"
	//"github.com/hyperledger/fabric/core/chaincode/shim" //1.4
	"github.com/hyperledger/fabric-chaincode-go/shim" //2.0 change
	"strings"
)

//import pb "github.com/hyperledger/fabric/protos/peer" //1.4
import pb "github.com/hyperledger/fabric-protos-go/peer" //2.0 change

type HistoryContract struct {
}

func (this *HistoryContract) Init(stub shim.ChaincodeStubInterface) pb.Response {

	var response pb.Response
	response.Status = shim.OK
	response.Payload = []byte("init is called")
	return response

}

func (this *HistoryContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	var response pb.Response
	response.Status = shim.ERROR
	response.Payload = []byte("unknown error")

	func_name, params := stub.GetFunctionAndParameters()
	if (strings.Compare("put_kv", func_name) == 0) {
		isSuccess, err_msg := put_kv(stub, params[0], params[1])
		if (isSuccess) {
			response.Status = shim.OK
			response.Payload = []byte("put success")
		} else {
			response.Payload = []byte("put falied-->" + err_msg)
		}
	}
	if (strings.Compare("get_value", func_name) == 0) {

		isSuccess, msg := get_value(stub, params[0])
		if (isSuccess) {
			response.Status = shim.OK
			response.Payload = []byte("get success-->" + msg)
		} else {
			response.Payload = []byte("get falied-->" + msg)
		}
	}

	if (strings.Compare("get_value_history", func_name) == 0) {

		isSuccess, msg := get_value_history(stub, params[0])
		if (isSuccess) {
			response.Status = shim.OK
			response.Payload = []byte("get history success-->" + msg)
		} else {
			response.Payload = []byte("get history falied-->" + msg)
		}
	}

	return response
}

func put_kv(stub shim.ChaincodeStubInterface, key string, value string) (isSuccess bool, err_msg string) {

	err := stub.PutState(key, []byte(value))
	if (err != nil) {
		return false, err.Error()
	}
	return true, ""
}

func get_value(stub shim.ChaincodeStubInterface, key string) (isSuccess bool, value_err string) {

	value, err := stub.GetState(key)
	if (err != nil) {
		return false, err.Error()
	}
	return true, string(value)

}

func get_value_history(stub shim.ChaincodeStubInterface, key string) (isSuccess bool, value_err string) {
	history, err := stub.GetHistoryForKey(key)
	if (err != nil) {
		return false, err.Error()
	}
	var result string = ""
	for history.HasNext() {
		response, err := history.Next()
		if (err != nil) {
			return false, err.Error()
		}
		history_item, err := json.Marshal(response)
		if (err != nil) {
			return false, err.Error()
		}
		result += "item-->" + string(history_item) + "\n"
	}
	return true, string(result)
}

func main() {
	err := shim.Start(new(HistoryContract))
	if err != nil {
		fmt.Println("fuck error -> " + err.Error())
	}

}

/*
$ docker exec -it cli /bin/bash
# peer chaincode install -n histroy_try -v v0 -p github.com/chaincode/histroy_try/go/
# peer chaincode instantiate -o orderer.example.com:7050 -C mychannel -n histroy_try -v v0 -c '{"Args":[]}' -P "AND('Org1MSP.member')"
# peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n histroy_try --peerAddresses peer0.org1.example.com:7051 -c '{"Args":["put_kv","test_key","test-001"]}'
# peer chaincode query -C mychannel -n histroy_try -c '{"Args":["get_value","test_key"]}'
# peer chaincode query -C mychannel -n histroy_try -c '{"Args":["get_value_history","test_key"]}'
*/
