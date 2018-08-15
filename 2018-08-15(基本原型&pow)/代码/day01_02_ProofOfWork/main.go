package main

import (
	"1802pubicchain/day01_02_ProofOfWork/BLC"
	"fmt"
)

func main() {

	//bytes:=make([]byte,3,3)
	//fmt.Println(bytes)
	//1.测试区块
	//block:=BLC.NewBlock("I am  a block",make([]byte,32,32),0)
	//fmt.Println(block)

	//2.测试创世区块
	//genesisBlock:=BLC.CreateGenesisBlock("Genesis Block")
	//fmt.Println(genesisBlock)

	//3.创建一个区块链
	blockChain:=BLC.CreateBlockChainWithGenesisBlock("Genesis Block")
	//fmt.Println(blockChain)
	//fmt.Println(blockChain.Blocks)
	//fmt.Println(blockChain.Blocks[0])

	//4.测试添加区块
	//blockChain:=BLC.CreateBlockChainWithGenesisBlock("Genesis Block")
	//blockChain.AddBlockToBlockChain("Send 100RMB to wangergou",blockChain.Blocks[len(blockChain.Blocks)-1].Hash,blockChain.Blocks[len(blockChain.Blocks)-1].Height+1)
	//blockChain.AddBlockToBlockChain("Send 100RMB to rose",blockChain.Blocks[len(blockChain.Blocks)-1].Hash,blockChain.Blocks[len(blockChain.Blocks)-1].Height+1)
	//blockChain.AddBlockToBlockChain("Send 1000RMB to lixiaohua",blockChain.Blocks[len(blockChain.Blocks)-1].Hash,blockChain.Blocks[len(blockChain.Blocks)-1].Height+1)
	//fmt.Println(blockChain)


	pow:=BLC.NewProofOfWork(blockChain.Blocks[0])
	fmt.Println(pow.IsValid())

	//s1:="hello world" // string
	//sha256的hash

	//写法一：
	//hash1:=sha256.Sum256([]byte(s1))
	//fmt.Printf("%T\n",hash1) //[32]uint8
	//fmt.Printf("%x\n",hash1)

	//写法二：
	//hasher:=sha256.New()
	//hasher.Write([]byte(s1))
	//hash2:= hasher.Sum(nil)
	//fmt.Printf("%x\n",hash2)
	//fmt.Printf("%T\n",hash2) //[]unit8



}