package BLC

import (
	"time"
	"strconv"
	"bytes"
	"crypto/sha256"
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
	TimeStamp int64  //2018-08-15 ---->"15909394005"-->[]byte
	//5.自己的hash
	Hash []byte

}

//step2：提供一个函数用于创建一个区块
func NewBlock(data string, prevBlockHash [] byte, height int64) *Block {
	//创建区块
	block := &Block{height, prevBlockHash, []byte(data), time.Now().Unix(), nil}
	//设置hash
	block.SetHash()
	return block
}

//step3：为区块设置hash值
/*
就是为指定的区块计算hash。
 */
func (block *Block) SetHash() {
	//可以通过当前的block的属性值来生成hash--->[]byte
	//1.Height-->[]byte
	heightBytes := IntToHex(block.Height)
	//2.TimeStamp-->[]byte
	//timeBytes:=IntToHex(block.TimeStamp) //"2018-10-10" ,"158030384844"
	//转为二进制的字符串
	timeString := strconv.FormatInt(block.TimeStamp, 2)
	timeBytes := []byte(timeString)

	//3.拼接所有的byte
	blockBytes := bytes.Join([][]byte{
		heightBytes,
		block.PrevBlockHash,
		block.Data,
		timeBytes}, []byte{})
	//4.设置到block上
	hash := sha256.Sum256(blockBytes)
	block.Hash = hash[:]

}

//step4:创建一个创世区块
func CreateGenesisBlock(data string) *Block{

	return NewBlock(data,make([]byte,32,32),0)
}
