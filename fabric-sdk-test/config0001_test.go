package test

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"testing"
)

func Test0001(test *testing.T) {

	configProvider := config.FromFile("./config0002.yaml")
	sdk, err := fabsdk.New(configProvider)
	if err != nil {
		fmt.Printf("create sdk failed --> %s\n", err.Error())
		return
	}
	fmt.Println("create sdk success")
	defer sdk.Close()

}
