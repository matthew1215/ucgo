package watchtransferscript

import (
	"ucgo/app/model/active/blockchain"
)

func Do() {
	watchTransferInModel := blockchain.WatchTransferInModel{}
	watchTransferInModel.WatchTransfer()
}
