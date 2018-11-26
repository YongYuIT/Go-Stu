package test

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"testing"
)

func Test0003(test *testing.T) {

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

}
