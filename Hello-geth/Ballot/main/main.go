package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
	"strconv"
)

func main() {
	//	geth attach http://127.0.0.1:8545
	//	跟本地geth节点建立rpc通信------------------------------------------------------------------
	gethAddress := "http://127.0.0.1:8545"
	cl, err := ethclient.Dial(gethAddress)
	defer cl.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("success conn to geth node at:	" + gethAddress)
	//	根据交易哈希，查询合约部署交易----------------------------------------------------------------
	contractSetUpTranx, isPending, err := cl.TransactionByHash(context.Background(), common.HexToHash("0xf16be0d428a89b756a4fd01d783b6b622c17177e9df4781f9c7a64d582dc3c64"))
	if err == nil && !isPending {
		fmt.Println("get contractSetUpTranx use gas: " + strconv.FormatUint(contractSetUpTranx.Gas(), 10))
	}
	//	调用合约----------------------------------------------------------------------------------
	//	abigen -abi Ballot.abi -type Ballot -pkg main -out Ballot.go
	ballot, err := NewBallot(common.HexToAddress("0x00DdEC4d9Cc8d27cBC66174bDEFb98197B5D4899"), cl)
	if err != nil {
		panic(err)
	}
	//	查询
	chairAddr, err := ballot.Chairperson(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("get chairAddr-->", chairAddr.Hex())
	//	发交易
	keyJson, err := os.ReadFile("./keys/UTC--2022-10-20T08-41-30.747047739Z--7359eeb9da4a47c9f671a74c4905753c69209627")
	if err != nil {
		panic(err)
	}
	key, _ := keystore.DecryptKey(keyJson, "")
	fmt.Println("get pub address: ", key.Address.Hex())
	chainId, err := cl.ChainID(context.Background())
	if err != nil {
		panic(err)
	}
	opts, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainId)
	if err != nil {
		panic(err)
	}
	tranx, err := ballot.GiveRightToVote(opts, common.HexToAddress("75cf56b6a261fe3635bbff45ae750dd6a74f0986"))
	fmt.Println("send tranx:", tranx.Hash())
	voter, err := ballot.Voters(nil, common.HexToAddress("75cf56b6a261fe3635bbff45ae750dd6a74f0986"))
	fmt.Println("get voter:", voter)
}
