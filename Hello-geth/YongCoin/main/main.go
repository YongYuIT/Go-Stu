package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os"
	"strconv"
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
	contractSetUpTranx, isPending, err := cl.TransactionByHash(context.Background(), common.HexToHash("0x7ec39001b485d996bee6bb67f58d95f004fe9b0bd124cf938c156ac9b5f95543"))
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
	deployerKeyJson, err := os.ReadFile("./keys/UTC--2022-10-28T07-05-43.432642863Z--44ff39a22d1c54960ae9a1e16e88ea8df4d656de")
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
		eth.sendTransaction({from:"0x44FF39a22D1c54960ae9A1e16e88EA8df4D656De",to:"0x86aE6D6d8C3E84A1ABCC9B6e146d7Ff980a62d01",value: web3.toWei(50,"ether")})
		eth.getBalance("0x86aE6D6d8C3E84A1ABCC9B6e146d7Ff980a62d01")
	*/
	//	call mint----------------------------------------------------------------------------------
	yongcoin, err := NewYongCoin(common.HexToAddress("0x542e1998371C9939F08a8aC2992EA3674E5293Af"), cl)
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
}

func listenSentEvent(sent <-chan *YongCoinSent) {
	fmt.Println(" #### ready to listen event")
	sentEvent := <-sent
	fmt.Println(" #### got sent: ", sentEvent)
	fmt.Println(" #### got sent: from:  ", sentEvent.From.Hex())
	fmt.Println(" #### got sent: to:  ", sentEvent.To.Hex())
	fmt.Println(" #### got sent: amount:  ", sentEvent.Amount)
}
