package BLC

import (
	"time"

)

//step1:定一个一Block
type Block struct {
	//字段属性
	//1.高度：区块在区块链中的编号，第一个区块页叫创世区块，为0
	Height int64
	//2.上一个区块的Hash值
	PrevBlockHash []byte
	//3.数据：data，交易数据
	Data []byte
	//4.时间戳
	TimeStamp int64
	//5.自己的hash
	Hash []byte
	//6.Nonce
	Nonce int64
}

//step2：提供一个函数用于创建一个区块
func NewBlock(data string, prevBlockHash [] byte, height int64) *Block {
	//创建区块
	block := &Block{height, prevBlockHash, []byte(data), time.Now().Unix(), nil,0}
	//设置hash
	//block.SetHash()
	pow:=NewProofOfWork(block)
	hash,nonce:=pow.Run()
	block.Hash = hash
	block.Nonce = nonce

	return block
}



//step4:创建一个创世区块
func CreateGenesisBlock(data string) *Block{

	return NewBlock(data,make([]byte,32,32),0)
}
