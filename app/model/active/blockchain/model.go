package blockchain

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
	"time"
	"ucgo/app/model/active/blockchain/myadvancedtoken"
	"ucgo/library/ires"
)

const key = `
{
	"address": "2f5707b0eec7577ad7bdb9259e293580ed33279c",
	"crypto": {
		"cipher": "aes-128-ctr",
		"ciphertext": "912e93d8dcf05f6a74c408cb171fded43faf5984c218f94f26450f005286fe6a",
		"cipherparams": {
			"iv": "a2d2e6fc9f129283aa932c3fee888bdd"
		},
		"kdf": "scrypt",
		"kdfparams": {
			"dklen": 32,
			"n": 262144,
			"p": 1,
			"r": 8,
			"salt": "dfb6ea6f608c38d6bc9e4e1a29549df874ac6e0aba49c23a8b1e282f440a48fe"
		},
		"mac": "0a1106a2f37163fe48ac14fbf635ec76af66084dd02a4c3bff2e27b52a5ad839"
	},
	"id": "1ec39291-fea1-47c9-92aa-2d6fade3670a",
	"version": 3
}
`
const passphrase = "123"
const contract_address = "0x89c7237c9848ede386d4ee0d4ae69bc374df131b"
const ipc_path = "/Users/tong/ethereum/data/geth.ipc"

type AddTransferInModel struct {
	Num int64
}
type AddTransferOutModel struct {
	TokenName    string
	TransferHash string
}

// 添加交易
func (this AddTransferInModel) AddTransfer() ires.ResponseStruct {
	var res AddTransferOutModel
	conn, err := ethclient.Dial(ipc_path)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	token, err := myadvancedtoken.NewMyadvancedtoken(common.HexToAddress(contract_address), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	name, err := token.Name(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve token name: %v", err)
	}
	res.TokenName = name

	auth, err := bind.NewTransactor(strings.NewReader(key), passphrase)
	toAddress := "0x970b5064504d914c58ea0a8e510f259b087a9e4949e11f6d7f878fdf7e6349de"
	tx, err := token.Transfer(auth, common.HexToAddress(toAddress), big.NewInt(this.Num))
	res.TransferHash = tx.Hash().String()

	return ires.Ok(res)
}

type WatchTransferInModel struct{}

// 观察交易
func (this WatchTransferInModel) WatchTransfer() {
	conn, err := ethclient.Dial(ipc_path)
	if err != nil {
		log.Fatal(err)
	}
	token, err := myadvancedtoken.NewMyadvancedtoken(common.HexToAddress(contract_address), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	logs := make(chan *myadvancedtoken.MyadvancedtokenTransfer)
	ctx := context.Background()
	opts := &bind.WatchOpts{nil, ctx}
	sub, err := token.WatchTransfer(opts, logs, nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog.Value.String())
		}
		// 协程跑脚本时（处理快只处理数据库的脚本)必须休眠1毫秒 减少数据库压力
		time.Sleep(time.Duration(1) * time.Millisecond)
	}
}
