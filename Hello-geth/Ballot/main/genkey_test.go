package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"os"
	"testing"
)

func TestGenKey(test *testing.T) {
	//	创建公私钥
	ks := keystore.NewKeyStore("./keys", keystore.StandardScryptN, keystore.StandardScryptP)
	account, _ := ks.NewAccount("123456")
	keyJson, err := ks.Export(account, "123456", "123456")
	if err != nil {
		panic(err)
	}
	key := string(keyJson)
	fmt.Println("get key: ", key)
	addr := account.Address.Hex()
	fmt.Println("get pub address: ", addr)
}

func TestImportKey(test *testing.T) {
	keyJson, err := os.ReadFile("./keys/UTC--2022-10-20T08-41-30.747047739Z--7359eeb9da4a47c9f671a74c4905753c69209627")
	if err != nil {
		panic(err)
	}
	key, _ := keystore.DecryptKey(keyJson, "")
	fmt.Println("get pub address: ", key.Address.Hex())
}
