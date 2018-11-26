package test

import (
	"fmt"
	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"testing"
)

func Test0002(test *testing.T) {

	configProvider := config.FromFile("./config0002.yaml")
	sdk, err := fabsdk.New(configProvider)
	if err != nil {
		fmt.Printf("create sdk failed --> %s\n", err.Error())
		return
	}
	fmt.Println("create sdk success")
	defer sdk.Close()

	mspClient, err := mspclient.New(sdk.Context(), mspclient.WithOrg("org1.example.com"))
	if err != nil {
		fmt.Printf("create msp client failed --> %s\n", err.Error())
		return
	}
	fmt.Println("create msp client success")

	adminIdentity, err := mspClient.GetSigningIdentity("Admin")
	if err != nil {
		fmt.Printf("get admin failed --> %s\n", err.Error())
	} else {
		fmt.Printf("AdminIdentify is found --> %s\n", adminIdentity)
	}

}
