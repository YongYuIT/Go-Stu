package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strings"
)

import pb "github.com/hyperledger/fabric/protos/peer"

type RichQueryContract struct {
}

func (this *RichQueryContract) Init(stub shim.ChaincodeStubInterface) pb.Response {

	var response pb.Response
	response.Status = shim.OK
	response.Payload = []byte("init is called")
	return response

}

func (this *RichQueryContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

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

	if (strings.Compare("rich_query", func_name) == 0) {

		isSuccess, msg := rich_query(stub, params[0])
		if (isSuccess) {
			response.Status = shim.OK
			response.Payload = []byte("get rich query success-->" + msg)
		} else {
			response.Payload = []byte("get rich query falied-->" + msg)
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

func rich_query(stub shim.ChaincodeStubInterface, query_str string) (isSuccess bool, value_err string) {
	resultIter, err := stub.GetQueryResult(query_str)
	if (err != nil) {
		return false, err.Error()
	}
	var result string = ""
	for resultIter.HasNext() {
		response, err := resultIter.Next()
		if (err != nil) {
			return false, err.Error()
		}
		item, err := json.Marshal(response)
		if (err != nil) {
			return false, err.Error()
		}
		result += "item-->" + string(item) + "\n"
	}
	return true, result
}

func main() {
	err := shim.Start(new(RichQueryContract))
	if err != nil {
		fmt.Println("fuck error -> " + err.Error())
	}

}

/*
$ docker exec -it cli /bin/bash
# peer chaincode install -n rich_query -v v0 -p github.com/chaincode/rich_query/go/
# peer chaincode instantiate -o orderer.example.com:7050 -C mychannel -n rich_query -v v0 -c '{"Args":[]}' -P "AND('Org1MSP.member')"
# peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n rich_query --peerAddresses peer0.org1.example.com:7051 -c '{"Args":["put_kv","test_key","test-001"]}'
# peer chaincode query -C mychannel -n rich_query -c '{"Args":["get_value","test_key"]}'

# peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n rich_query --peerAddresses peer0.org1.example.com:7051 -c '{"Args":["put_kv","10000001","{\"address\":\"ShenZhen city,GD,China\",\"age\":20,\"country\":\"CN\",\"gender\":\"M\",\"name\":\"yong\",\"phone_num\":\"110\",\"stu_no\":\"10000001\"}"]}'
# peer chaincode query -C mychannel -n rich_query -c '{"Args":["get_value","10000001"]}'
# peer chaincode query -C mychannel -n rich_query -c '{"Args":["rich_query","{\"selector\":{\"age\":{\"$lt\":19}}}"]}'
# peer chaincode query -C mychannel -n rich_query -c '{"Args":["rich_query","{\"selector\":{\"age\":{\"$gt\":19}}}"]}'
*/
