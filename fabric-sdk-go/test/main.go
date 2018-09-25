package main

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"log"
)

func main() {

	//读取配置文件，创建SDK
	configProvider := config.FromFile("./config.yaml")
	sdk, err := fabsdk.New(configProvider)
	if err != nil {
		log.Fatalf("create sdk fail: %s\n", err.Error())
	}

	//读取配置文件(config.yaml)中的组织(yong.thinking.com)的用户(Admin)
	mspClient, err := mspclient.New(sdk.Context(),
		mspclient.WithOrg("yong.thinking.com"))
	if err != nil {
		log.Fatalf("create msp client fail: %s\n", err.Error())
	}

	adminIdentity, err := mspClient.GetSigningIdentity("Admin")
	if err != nil {
		log.Fatalf("get admin identify fail: %s\n", err.Error())
	} else {
		fmt.Println("AdminIdentify is found:")
		fmt.Println(adminIdentity)
	}

	//读取配置文件(config.yaml)中的通道(channels)的当前通道(mythnkingchannel)，这一步需要首先通过cli节点创建通道并加入，参考《区块链二十五 Hello-word链码》
	channelProvider := sdk.ChannelContext("mythnkingchannel",
		fabsdk.WithUser("Admin"),
		fabsdk.WithOrg("yong.thinking.com"))
	channelClient, err := channel.New(channelProvider)
	if err != nil {
		log.Fatalf("create channel client fail: %s\n", err.Error())
	}

	//调用链码，这一步需要在cli节点上手动安装链码，参考《区块链二十七 Hello-word链码 分析》
	var args [][]byte
	args = append(args, []byte("{\"Args\":[\"init\",\"a\",\"100\",\"b\",\"200\"]}"))

	request := channel.Request{
		ChaincodeID: "iptest20180816001",
		Fcn:         "invoke",
		Args:        args,
	}
	response, err := channelClient.Query(request)
	if err != nil {
		log.Fatal("query fail: ", err.Error())
	} else {
		fmt.Printf("response is %s\n", response.Payload)
	}
}
