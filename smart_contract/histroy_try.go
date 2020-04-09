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


/* if fabric 2.0------------------------------------------------------------------------------------------------------------------------------
#### 修改fabric-samples/first-network/docker-compose-cli.yaml文件，添加mod的映射，避免重复下载依赖
    volumes:
        - /mnt/hgfs/go-env-1/go-path/pkg/mod/:/opt/gopath/pkg/mod/
$ docker exec -it cli /bin/bash
# cd /opt/gopath/src/github.com/hyperledger/fabric-samples/chaincode
#### 将go.mod文件里的module改为github.com/hyperledger/fabric-samples/chaincode/histroy_try/go
#### module github.com/hyperledger/fabric-samples/chaincode/histroy_try/go
# export GO111MODULE=on
# peer lifecycle chaincode package histroy_try_v1.tar.gz --path histroy_try/go/ --lang golang --label histroy_try_v1
# ls ./histroy_try_v1.tar.gz
# peer lifecycle chaincode install histroy_try_v1.tar.gz
# peer lifecycle chaincode queryinstalled
#### 注意，下面的包id每次打包都会不一样
# peer lifecycle chaincode approveformyorg --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --channelID mychannel --name histroy_try_v1 --version 1 --init-required --package-id histroy_try_v1:9e11241ded7f81b8c2fbd835be8caa7990fd07184c0f4117c2a23cb85efdd5c6 --sequence 1 --waitForEvent
# peer lifecycle chaincode checkcommitreadiness --channelID mychannel --name histroy_try_v1 --version 1 --sequence 1 --output json --init-required

#### 转到peer0.org2，安装 & 批准链码
# export CORE_PEER_ADDRESS=peer0.org2.example.com:9051
# export CORE_PEER_LOCALMSPID=Org2MSP
# export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
# export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin\@org2.example.com/msp/
# cd /opt/gopath/src/github.com/hyperledger/fabric-samples/chaincode
# peer lifecycle chaincode install histroy_try_v1.tar.gz
# peer lifecycle chaincode queryinstalled
# peer lifecycle chaincode checkcommitreadiness --channelID mychannel --name histroy_try_v1 --version 1 --sequence 1 --output json --init-required
# peer lifecycle chaincode approveformyorg --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --channelID mychannel --name histroy_try_v1 --version 1 --init-required --package-id histroy_try_v1:9e11241ded7f81b8c2fbd835be8caa7990fd07184c0f4117c2a23cb85efdd5c6 --sequence 1 --waitForEvent
# peer lifecycle chaincode checkcommitreadiness --channelID mychannel --name histroy_try_v1 --version 1 --sequence 1 --output json --init-required

#### 转到peer0.org1，部署链码
# peer lifecycle chaincode commit -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --channelID mychannel --name histroy_try_v1 --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt --version 1 --sequence 1 --init-required

# peer lifecycle chaincode querycommitted --channelID mychannel --name histroy_try_v1
# peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n histroy_try_v1 --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt --isInit -c '{"Args":[]}'
# peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n histroy_try_v1 --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt  -c '{"Args":["put_kv","test_key","test-001"]}'
# peer chaincode query -C mychannel -n histroy_try_v1 -c '{"Args":["get_value","test_key"]}'
# peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n histroy_try_v1 --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt  -c '{"Args":["put_kv","test_key","test-002"]}'
# peer chaincode query -C mychannel -n histroy_try_v1 -c '{"Args":["get_value_history","test_key"]}'
*/
