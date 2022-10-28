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
	fmt.Println("start test...")
	//	abigen -abi YongCoin.abi -type YongCoin -pkg main -out YongCoin.go
	//	跟本地geth节点建立rpc通信------------------------------------------------------------------
	gethAddress := "http://127.0.0.1:8545"
	cl, err := ethclient.Dial(gethAddress)
	defer cl.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("success conn to geth node at:	" + gethAddress)
	//	根据交易哈希，查询合约部署交易----------------------------------------------------------------
	contractSetUpTranx, isPending, err := cl.TransactionByHash(context.Background(), common.HexToHash("0xf61f4b997bb62137e243621ede6008f6f36fc3770e0b7d6ed71166b1e64cd8dd"))
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
		./geth --dev --dev.period 0 --http --http.api eth,web3,personal,net --http.corsdomain="package://6fd22d6fe5549ad4c4d8fd3ca0b7816b.mod" --datadir ./tmpdata
		eth.sendTransaction({from:"0x44FF39a22D1c54960ae9A1e16e88EA8df4D656De",to:"25a00d211b6e0459d3b593b59285d16e7c06111b",value: web3.toWei(50,"ether")})
		eth.getBalance("25a00d211b6e0459d3b593b59285d16e7c06111b")
	*/
	//amount := big.NewInt(10)
	//var gasLimit uint64 = 300000
	//var gasPrice *big.Int = big.NewInt(200)
	//nonce, err := cl.PendingNonceAt(context.Background(), deployerKey.Address)
	//auth, err := bind.NewTransactorWithChainID(bytes.NewReader(deployerKeyJson), "", chainId)
	//auth.Nonce = big.NewInt(int64(nonce))
	//auth.Value = amount
	//auth.GasLimit = gasLimit
	//auth.GasPrice = gasPrice
	//auth.From = deployerKey.Address
	//tx := types.NewTransaction(nonce, userAccount.Address, amount, gasLimit, gasPrice, []byte{})
	//signedTx, err := auth.Signer(auth.From, tx)
	//sendEthTranx := cl.SendTransaction(context.Background(), signedTx)
	//if sendEthTranx != nil {
	//	panic(err)
	//}
	//bind.WaitMined(context.Background(), cl, signedTx)
	//fmt.Println("send eth hash: ", signedTx.Hash())
	//	call mint----------------------------------------------------------------------------------
	yongcoin, err := NewYongCoin(common.HexToAddress("0xa0E8a077b907b34E6c874DaBEa169009de010B58"), cl)
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
