package BLC

import "fmt"

//定义一个钱包的集合，存储多个钱包对象
type Wallets struct {
	WalletMap map[string]*Wallet
}

//提供一个函数，用于创建一个钱包的集合
func NewWallets() *Wallets{
	wallerts:=&Wallets{}
	wallerts.WalletMap = make(map[string]*Wallet)
	return wallerts
}

func (ws *Wallets) CreateNewWallet(){
	wallet:=NewWallet()

	address := wallet.GetAddress()
	fmt.Printf("创建的钱包地址：%s\n",address)

	ws.WalletMap[string(address)] = wallet
}
