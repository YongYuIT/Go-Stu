package test

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"testing"
)

func Test0004(test *testing.T) {

	configProvider := config.FromFile("./config0003.yaml")
	sdk, err := fabsdk.New(configProvider)
	if err != nil {
		fmt.Printf("create sdk failed --> %s\n", err.Error())
		return
	}
	fmt.Println("create sdk success")
	defer sdk.Close()

	channelProvider := sdk.ChannelContext("mychannel", fabsdk.WithUser("Admin"), fabsdk.WithOrg("org1.example.com"))
	channelClient, err := channel.New(channelProvider)
	if err != nil {
		fmt.Printf("create channel client failed --> %s\n", err.Error())
		return
	}
	if channelClient != nil {
		fmt.Println("create channel client success")
	}

	var args [][]byte
	args = append(args, []byte("{\"Args\":[]}"))

	//https://blog.csdn.net/YongYu_IT/article/details/84066898
	request := channel.Request{
		ChaincodeID: "test_my_hello",
		Fcn:         "query",
		Args:        args,
	}
	response, err := channelClient.Query(request)
	if err != nil {
		fmt.Printf("query fail --> %s\n", err.Error())
		return
	} else {
		fmt.Printf("response is --> %s\n", response.Payload)
	}

}
