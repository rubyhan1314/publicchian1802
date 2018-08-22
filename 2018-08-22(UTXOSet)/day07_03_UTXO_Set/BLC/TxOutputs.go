package BLC

import (
	"bytes"
	"encoding/gob"
	"log"
)

type TxOutputs struct {
	UTXOs []*UTXO
}


//序列化
func (outs *TxOutputs) Serialize()[]byte{
	var buff bytes.Buffer

	encoder:=gob.NewEncoder(&buff)

	err :=encoder.Encode(outs)
	if err != nil{
		log.Panic(err)
	}
	return buff.Bytes()
}