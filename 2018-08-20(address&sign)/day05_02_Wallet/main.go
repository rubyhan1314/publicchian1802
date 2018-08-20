package main

import (
	"./BLC"
	"fmt"
)

func main() {
	//1.测试钱包对象
	wallet:=BLC.NewWallet()
	fmt.Println(wallet)
	fmt.Println("私钥：",wallet.PrivateKey)
	fmt.Println("公钥：",wallet.PublickKey)
	fmt.Println(len(wallet.PublickKey))

	address:=wallet.GetAddress()
	fmt.Println(address)
	fmt.Println(string(address))

	fmt.Println("地址是否有效：",BLC.IsValidAddress(address))
	fmt.Println("地址是否有效：",BLC.IsValidAddress([]byte("wangergou")))
}
