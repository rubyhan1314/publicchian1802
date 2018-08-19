package BLC

import (
	"bytes"
	"encoding/gob"
	"log"
	"crypto/sha256"
)

//定义交易的数据
type Transaction struct {
	//1.交易ID-->就是交易的Hash
	TxID []byte
	//2.输入
	Vins []*TxInput
	//3.输出
	Vouts []*TxOutput
}

/*
交易：
1.CoinBase交易：创世区块中
2.转账产生的普通交易：
 */

 func NewCoinBaseTransaction(address string) *Transaction{
 	txInput:=&TxInput{[]byte{},-1,"Genesis Data"}
 	txOutput:=&TxOutput{10,address}
 	txCoinBaseTransaction:=&Transaction{[]byte{},[]*TxInput{txInput},[]*TxOutput{txOutput}}
 	//设置交易ID
 	txCoinBaseTransaction.SetID()
 	return txCoinBaseTransaction
 }

 //交易ID--->根据tx，生成一个hash
 func (tx *Transaction) SetID(){
 	//1.tx--->[]byte
 	var buf bytes.Buffer
 	encoder:=gob.NewEncoder(&buf)
 	err:=encoder.Encode(tx)
 	if err != nil{
 		log.Panic(err)
	}
 	//2.[]byte-->hash
 	hash:=sha256.Sum256(buf.Bytes())
 	//3.为tx设置ID
 	tx.TxID = hash[:]
 }
