package BLC

import (
	"fmt"
	"os"
)

func (cli *CLI) StartNode(nodeID string, mineAddress string) {
	//启动服务器
	if mineAddress == "" || IsValidAddress([]byte(mineAddress)) {
		//启动服务器
		startServer(nodeID, mineAddress)
	} else {
		fmt.Println("地址无效。。")
		os.Exit(1)
	}
}
