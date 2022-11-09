package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	/**
	./geth --dev --dev.period 0 --http --http.api eth,web3,personal,net --http.corsdomain="package://6fd22d6fe5549ad4c4d8fd3ca0b7816b.mod" --datadir ./tmpdata --ws --ws.port 3334
	*/
	fmt.Println("start test...")
	//	abigen -abi YongCoin.abi -type YongCoin -pkg main -out YongCoin.go
	//	跟本地geth节点建立rpc通信------------------------------------------------------------------
	//gethAddress := "http://127.0.0.1:8545" //http not support event
	gethAddress := "ws://127.0.0.1:3334" //websocket support event, 注意：如果需要getn节点支持ws协议，启动时需要加上 --ws --ws.port 3334
	cl, err := ethclient.Dial(gethAddress)
	defer cl.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("success conn to geth node at:	" + gethAddress)
	//	根据交易哈希，查询合约部署交易----------------------------------------------------------------
	contractSetUpTranx, isPending, err := cl.TransactionByHash(context.Background(), common.HexToHash("0xd5c1852f4da3b46bf110a0bf746054f47622891c922c04f91c066913c9f35166"))
	if err == nil && !isPending {
		fmt.Println("get contractSetUpTranx use gas: " + strconv.FormatUint(contractSetUpTranx.Gas(), 10))
	}
	//	创建公私钥---------------------------------------------------------------------------------
	userKeyStore := keystore.NewKeyStore("./keys", keystore.StandardScryptN, keystore.StandardScryptP)
	userAccount, _ := userKeyStore.NewAccount("123456")
	userKeyJson, err := userKeyStore.Export(userAccount, "123456", "123456")
	if err != nil {
		panic(err)
	}
	userKeyStr := string(userKeyJson)
	fmt.Println("get userKeyStr: ", userKeyStr)
	addr := userAccount.Address.Hex()
	fmt.Println("get pub address: ", addr)
	//	get deployer pri-pub userKeyStr-----------------------------------------------------------
	deployerKeyJson, err := os.ReadFile("./keys/UTC--2022-11-09T03-23-27.218433604Z--29768bce1cdb6cd22b4a47209d375d3951e3cf8d")
	if err != nil {
		panic(err)
	}
	deployerKey, _ := keystore.DecryptKey(deployerKeyJson, "")
	fmt.Println("get deployer address: ", deployerKey.Address.Hex())
	chainId, err := cl.ChainID(context.Background())
	if err != nil {
		panic(err)
	}
	deployerOpts, err := bind.NewKeyedTransactorWithChainID(deployerKey.PrivateKey, chainId)
	if err != nil {
		panic(err)
	}
	//	deployer give some money to user-------------------------------------------------------------
	/*
		eth.sendTransaction({from:"0x29768bCe1CDB6cD22B4a47209D375d3951e3cF8D",to:"0x9bbe4c52EeDe518775F0f776b6944904E6037c38",value: web3.toWei(50,"ether")})
		eth.getBalance("0x9bbe4c52EeDe518775F0f776b6944904E6037c38")
	*/
	//	call mint----------------------------------------------------------------------------------
	yongcoin, err := NewYongCoin(common.HexToAddress("0x863459c4c6975d0d289Df40f9eBf14e193da0B9C"), cl)
	if err != nil {
		panic(err)
	}
	mintTranx, err := yongcoin.Mint(deployerOpts, userAccount.Address, new(big.Int).SetUint64(uint64(1000)))
	if err != nil {
		panic(err)
	}
	fmt.Println("mint tranx hash: ", mintTranx.Hash().Hex())
	//	tranf 500 to deployer----------------------------------------------------------------------------------
	userDecKey, _ := keystore.DecryptKey(userKeyJson, "123456")
	fmt.Println("get user address: ", userDecKey.Address.Hex())
	userOpts, err := bind.NewKeyedTransactorWithChainID(userDecKey.PrivateKey, chainId)
	if err != nil {
		panic(err)
	}
	//	WatchSent------------
	chanSent := make(chan *YongCoinSent)
	watchOpt := new(bind.WatchOpts)
	watchOpt.Context = context.Background()
	yongcoin.WatchSent(watchOpt, chanSent)
	go listenSentEvent(chanSent)
	//	WatchSent------------
	sendTranx, err := yongcoin.Send(userOpts, deployerKey.Address, new(big.Int).SetUint64(uint64(500)))
	if err != nil {
		panic(err)
	}
	fmt.Println("send tranx hash: ", sendTranx.Hash().Hex())
	//	query deployer----------------------------------------------------------------------------------
	deployerBalance, err := yongcoin.Balance(nil, deployerKey.Address)
	userBalance, err := yongcoin.Balance(nil, userDecKey.Address)
	fmt.Println("deployerBalance:", deployerBalance)
	fmt.Println("userBalance:", userBalance)
	//	解析send交易回执
	readReceipt(cl, sendTranx.Hash())
}

func readReceipt(client *ethclient.Client, hash common.Hash) {

	abi, err := abi.JSON(strings.NewReader(YongCoinMetaData.ABI))
	if err != nil {
		panic(err)
	}
	receipt, err := client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		panic(err)
	}
	fmt.Println("get receipt: ", receipt.TxHash.Hex())
	for i, log := range receipt.Logs {
		fmt.Println("get log start --> ", i)
		for s, event := range abi.Events {
			fmt.Println("event str-->", s, "; event name-->", event.Name)
			logValue, _ := event.Inputs.Unpack(log.Data)
			fmt.Println("get log value:", logValue)
		}
		for i2, topic := range log.Topics {
			fmt.Println("get log topic --> ", i2, " topic-->", topic.Hex())
		}
		fmt.Println("get log end --> ", i)
	}
}

func listenSentEvent(sent <-chan *YongCoinSent) {
	fmt.Println(" #### ready to listen event")
	sentEvent := <-sent
	fmt.Println(" #### got sent: ", sentEvent)
	fmt.Println(" #### got sent: from:  ", sentEvent.From.Hex())
	fmt.Println(" #### got sent: to:  ", sentEvent.To.Hex())
	fmt.Println(" #### got sent: amount:  ", sentEvent.Amount)
}
