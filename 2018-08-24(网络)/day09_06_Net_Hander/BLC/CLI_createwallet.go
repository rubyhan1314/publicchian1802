package BLC

import "fmt"

func (cli *CLI) CreateWallet(nodeID string){
	wallets:=NewWallets(nodeID) //获取钱包集合
	wallets.CreateNewWallet(nodeID)//创建钱包对象
	fmt.Println("钱包：",wallets.WalletMap)
}
