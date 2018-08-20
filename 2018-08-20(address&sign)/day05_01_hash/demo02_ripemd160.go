package main

import (
	"golang.org/x/crypto/ripemd160"
	"fmt"
)

func main() {
	/*
	先需要安装：go get "golang.org/x/crypto/ripemd160"
	 */

	bytes := []byte{12, 48, 79, 32, 90, 46}
	hasher := ripemd160.New()
	hasher.Write(bytes)
	hash3 := hasher.Sum(nil)
	fmt.Printf("%T\n", hash3) //[]uint8
	fmt.Printf("%x\n", hash3)

}
