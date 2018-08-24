package BLC

import (
	"fmt"
	"os"
)

func (cli *CLI) Send(from, to, amount []string,nodeID string) {
	bc := GetBlockChainObject(nodeID)
	if bc == nil {
		fmt.Println("没有BlockChain，无法转账。。")
		os.Exit(1)
	}
	defer bc.DB.Close()

	bc.MineNewBlock(from, to, amount,nodeID)
	//添加更新
	utsoSet :=&UTXOSet{bc}
	utsoSet.Update()
}
