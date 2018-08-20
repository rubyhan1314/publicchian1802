package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	//原始数据
	bytes:=[]byte{12,48,79,32,90,46}

	//方法1
	hash1:=sha256.Sum256(bytes)
	fmt.Printf("%T\n",hash1) //[32]uint8
	fmt.Printf("%x\n",hash1) //c7b06162ac3d0e23046ee87f50ff3a5eacd330afad0fcc551e7cb891f4632add

	//方法2
	hasher:=sha256.New()
	hasher.Write(bytes)
	hash2:=hasher.Sum(nil)
	fmt.Printf("%T\n",hash2) //[]uint8
	fmt.Printf("%x\n",hash2) //c7b06162ac3d0e23046ee87f50ff3a5eacd330afad0fcc551e7cb891f4632add
}
