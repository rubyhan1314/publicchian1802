package BLC

import (
	"github.com/boltdb/bolt"
	"log"
	"encoding/hex"
	"fmt"
)

/*
持久化：
	数据库：blockchain.db
		数据表(bucket) blocks
			存储所有的block

		数据表(bucket) utxoset
			存储所有的未花费utxo


查询余额，转账
 */
type UTXOSet struct {
	BlockChian *BlockChain
}

const utxosettable = "utxoset"

//提供一个重置的功能：获取blockchain中所有的未花费utxo

/*
查询block块中所有的未花费utxo：执行FindUnspentUTXOMap--->map

 */
func (utxoset *UTXOSet) ResetUTXOSet() {
	err := utxoset.BlockChian.DB.Update(func(tx *bolt.Tx) error {
		//1.utxoset表存在，删除
		b := tx.Bucket([]byte(utxosettable))
		if b != nil {
			err := tx.DeleteBucket([]byte(utxosettable))
			if err != nil {
				log.Panic("重置时，删除表失败。。")
			}
		}

		//2.创建utxoset
		b, err := tx.CreateBucket([]byte(utxosettable))
		if err != nil {
			log.Panic("重置时，创建表失败。。")
		}
		if b != nil {
			//3.将map数据--->表
			unUTXOMap := utxoset.BlockChian.FindUnspentUTXOMap()
			/*
			map:
				key:[string]-->[]byte
				value:*Txoutputs{[]UTXO}



			 */
			for txIDStr, outs := range unUTXOMap {
				txID, _ := hex.DecodeString(txIDStr) //[]byte
				b.Put(txID, outs.Serialize())
			}
			fmt.Println("啦啦啦啦。。。。。")
		}

		return nil

	})

	if err != nil {
		log.Panic(err)
	}

}

//查询余额
func (utxoSet *UTXOSet) GetBalance(address string) int64 {
	utxos := utxoSet.FindUnspentUTXOsByAddress(address)

	var total int64

	for _, utxo := range utxos {
		total += utxo.Output.Value
	}
	return total
}

//根据指定的地址，找出所有的utxo
func (utxoSet *UTXOSet) FindUnspentUTXOsByAddress(address string) []*UTXO {
	var utxos []*UTXO
	err := utxoSet.BlockChian.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(utxosettable))
		if b != nil {
			/*
			获取表中的所有的数据
			key,value
			key:TxID
			value：TxOuputs
			 */
			c := b.Cursor()
			for k, v := c.First(); k != nil; k, v = c.Next() {
				txOutputs := DeserializeTxOutputs(v)
				for _, utxo := range txOutputs.UTXOs { //txid, index,output
					if utxo.Output.UnlockWithAddress(address) {
						utxos = append(utxos, utxo)
					}
				}
			}
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	return utxos
}

//增加功能：
/*
添加一个方法，用于查询要转账的utxo
二狗，转账小花：5
	二狗：
		utxo
 */

func (utxoSet *UTXOSet) FindSpentableUTXOs(from string, amount int64, txs []*Transaction) (int64, map[string][]int) {
	var total int64
	//用于存储转账所使用utxo
	spentableUTXOMap := make(map[string][]int)
	//1.查询未打包utxo：txs
	unPackageSpentableUTXOs := utxoSet.FindUnpackeSpentableUTXO(from, txs)

	for _,utxo:=range unPackageSpentableUTXOs{
		total += utxo.Output.Value
		txIDStr:=hex.EncodeToString(utxo.TxID)
		spentableUTXOMap[txIDStr] =append(spentableUTXOMap[txIDStr],utxo.Index)
		if total >= amount{
			return total,spentableUTXOMap
		}
	}

	//钱不够，


	//2.查询utxotable，查询utxo
	//已经存储的都是未花费
	err :=utxoSet.BlockChian.DB.View(func(tx *bolt.Tx) error {
		//查询utxotable中，未花费的utxo
		b :=tx.Bucket([]byte(utxosettable))
		if b != nil{
			//查询
			c := b.Cursor()
			dbLoop:
			for k,v:=c.First();k!=nil;k,v= c.Next(){
				txOutputs:=DeserializeTxOutputs(v)
				for _,utxo:=range txOutputs.UTXOs{
					if utxo.Output.UnlockWithAddress(from){
						total += utxo.Output.Value
						txIDStr:=hex.EncodeToString(utxo.TxID)
						spentableUTXOMap[txIDStr] = append(spentableUTXOMap[txIDStr],utxo.Index)
						if total >= amount{
							break dbLoop
							//return nil
						}
					}
				}



			}


		}

		return nil

	})
	if err != nil{
		log.Panic(err)
	}

	return total,spentableUTXOMap
}

//查询未打包的tx中，可以使用的utxo
func (utxoSet *UTXOSet) FindUnpackeSpentableUTXO(from string, txs []*Transaction) []*UTXO {
	//存储可以使用的未花费utxo
 	var unUTXOs []*UTXO

 	//存储已经花费的input
 	spentedMap:=make(map[string][]int)

 	for i:=len(txs) -1;i>=0;i--{
 		//func caculate(tx *Transaction, address string, spentTxOutputMap map[string][]int, unSpentUTXOs []*UTXO) []*UTXO {
		unUTXOs = caculate(txs[i],from,spentedMap,unUTXOs)
	}


 	return unUTXOs
}
