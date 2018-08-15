package main

import (
	"1802pubicchain/day01_01_Base_Prototype/BLC"
	"fmt"
)

func main() {

	//bytes:=make([]byte,3,3)
	//fmt.Println(bytes)
	//1.测试区块
	//block:=BLC.NewBlock("I am  a block",make([]byte,32,32),0)
	//fmt.Println(block)
	//fmt.Println(block.TimeStamp)
	//fmt.Println(block.PrevBlockHash)
	//fmt.Println(block.Hash)

	//2.测试创世区块
	//genesisBlock:=BLC.CreateGenesisBlock("Genesis Block")
	//fmt.Println(genesisBlock)

	//3.创建一个区块链
	blockChain := BLC.CreateBlockChainWithGenesisBlock("Genesis Block")
	//fmt.Println(blockChain)
	//fmt.Println(blockChain.Blocks) // blockchain 的字段数组
	//fmt.Println(blockChain.Blocks[0]) //创世区块

	//4.测试添加区块
	fmt.Println(len(blockChain.Blocks)) //区块链中，有一个区块：创世区块
	blockChain.AddBlockToBlockChain("send 100RMB to wangergou", blockChain.Blocks[len(blockChain.Blocks)-1].Hash, blockChain.Blocks[len(blockChain.Blocks)-1].Height+1)
	blockChain.AddBlockToBlockChain("send 50RMB to lixiaohua", blockChain.Blocks[len(blockChain.Blocks)-1].Hash, blockChain.Blocks[len(blockChain.Blocks)-1].Height+1)
	blockChain.AddBlockToBlockChain("send 3RMB to rose", blockChain.Blocks[len(blockChain.Blocks)-1].Hash, blockChain.Blocks[len(blockChain.Blocks)-1].Height+1)
	blockChain.AddBlockToBlockChain("send 1000RMB to jack", blockChain.Blocks[len(blockChain.Blocks)-1].Hash, blockChain.Blocks[len(blockChain.Blocks)-1].Height+1)
	fmt.Println(blockChain)
	fmt.Println(blockChain.Blocks[1].Height)
	fmt.Println(blockChain.Blocks[1].PrevBlockHash)
	fmt.Println(blockChain.Blocks[0].Hash)

	fmt.Println(len(blockChain.Blocks))
}
