package BLC

import (
	"bytes"
	"encoding/gob"
	"log"
	"crypto/sha256"
	"encoding/hex"
	"crypto/ecdsa"
	"crypto/rand"
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

func NewCoinBaseTransaction(address string) *Transaction {
	txInput := &TxInput{[]byte{}, -1, nil, nil}
	//txOutput := &TxOutput{10, address}
	txOutput := NewTxOutput(10, address)
	txCoinBaseTransaction := &Transaction{[]byte{}, []*TxInput{txInput}, []*TxOutput{txOutput}}
	//设置交易ID
	txCoinBaseTransaction.SetID()
	return txCoinBaseTransaction
}

//交易ID--->根据tx，生成一个hash
func (tx *Transaction) SetID() {
	//1.tx--->[]byte
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	//2.[]byte-->hash
	hash := sha256.Sum256(buf.Bytes())
	//3.为tx设置ID
	tx.TxID = hash[:]
}

//根据转账的信息，创建一个普通的交易
func NewSimpleTransaction(from, to string, amount int64, bc *BlockChain, txs []*Transaction) *Transaction {
	//1.定义Input和Output的数组
	var txInputs []*TxInput
	var txOuputs [] *TxOutput

	//2.创建Input
	/*
	创世区块中交易ID：c16d3ad93450cd532dcd7ef53d8f396e46b2e59aa853ad44c284314c7b9db1b4
	 */

	//获取本次转账要使用output
	total, spentableUTXO := bc.FindSpentableUTXOs(from, amount, txs) //map[txID]-->[]int{index}

	//获取钱包的集合：
	wallets := NewWallets()
	wallet := wallets.WalletMap[from]

	for txID, indexArray := range spentableUTXO {
		txIDBytes, _ := hex.DecodeString(txID)
		for _, index := range indexArray {
			txInput := &TxInput{txIDBytes, index, nil, wallet.PublickKey}
			txInputs = append(txInputs, txInput)
		}
	}

	//idBytes, _ := hex.DecodeString("c16d3ad93450cd532dcd7ef53d8f396e46b2e59aa853ad44c284314c7b9db1b4")
	//idBytes, _ := hex.DecodeString("143d7db0d5cce24645edb2ba0b503fe15969ade0c721edfd3578cd731c563a16")
	//txInput := &TxInput{idBytes, 1, from}
	//txInputs = append(txInputs, txInput)

	//3.创建Output

	//转账
	//txOutput := &TxOutput{amount, to}
	txOutput := NewTxOutput(amount, to)
	txOuputs = append(txOuputs, txOutput)

	//找零
	//txOutput2 := &TxOutput{total - amount, from}
	txOutput2 := NewTxOutput(total-amount, from)
	txOuputs = append(txOuputs, txOutput2)

	//创建交易
	tx := &Transaction{[]byte{}, txInputs, txOuputs}

	//设置交易的ID
	tx.SetID()


	//设置签名
	bc.SignTrasanction(tx,wallet.PrivateKey)


	return tx

}

//判断tx是否时CoinBase交易
func (tx *Transaction) IsCoinBaseTransaction() bool {

	return len(tx.Vins[0].TxID) == 0 && tx.Vins[0].Vout == -1
}

//签名
/*
签名：为了对一笔交易进行签名
	私钥：
	要获取交易的Input，引用的output，所在的之前的交易：
 */
func (tx *Transaction) Sign(privateKey ecdsa.PrivateKey, prevTxsmap map[string]*Transaction) {
	//1.判断当前tx是否时coinbase交易
	if tx.IsCoinBaseTransaction() {
		return
	}

	//2.获取input对应的output所在的tx，如果不存在，无法进行签名
	for _, input := range tx.Vins {
		if prevTxsmap[hex.EncodeToString(input.TxID)] == nil {
			log.Panic("当前的Input，没有找到对应的output所在的Transaction，无法签名。。")
		}
	}

	//即将进行签名:私钥，要签名的数据
	txCopy := tx.TrimmedCopy()

	for index, input := range txCopy.Vins {
		// input--->5566

		prevTx := prevTxsmap[hex.EncodeToString(input.TxID)]

		//txCopy.Vins[index].Signature = nil                                 //仅仅是一个双重保险，保证签名一定为空
		input.Signature = nil

		//txCopy.Vins[index].PublicKey = prevTx.Vouts[input.Vout].PubKeyHash //设置input中的publickey为对应的output的公钥哈希
		input.PublicKey = prevTx.Vouts[input.Vout].PubKeyHash

		txCopy.TxID = txCopy.NewTxID() //设置要签名的数据：设置到id

		//为了方便下一个input，将数据再置为空
		txCopy.Vins[index].PublicKey = nil


		//获取要交易的数据

		/*
		第一个参数
		第二个参数：私钥
		第三个参数：要签名的数据


		func Sign(rand io.Reader, priv *PrivateKey, hash []byte) (r, s *big.Int, err error)
		r + s--->sign
		input.Signatrue = sign
	 */
		//r,s,err:=ecdsa.Sign(rand.Reader, &privateKey, txCopy.TxID )
		r,s,err:=ecdsa.Sign(rand.Reader, &privateKey, txCopy.NewTxID() )
		if err != nil{
			log.Panic(err)
		}

		sign:=append(r.Bytes(),s.Bytes()...)
		tx.Vins[index].Signature = sign
	}

}

//获取要签名tx的副本
/*
要签名tx中，并不是所有的数据都要作为签名数据，生成签名
txCopy = tx{签名所需要的部分数据}
TxID

Inputs
	txid,vout,sign,publickey

Outputs
	value,pubkeyhash


交易的副本中包含的数据
	包含了原来tx中的输入和输出。
		输入中：sign，publickey
 */

func (tx *Transaction) TrimmedCopy() *Transaction {
	var inputs [] *TxInput
	var outputs [] *TxOutput
	for _, in := range tx.Vins {
		inputs = append(inputs, &TxInput{in.TxID, in.Vout, nil, nil})
	}

	for _, out := range tx.Vouts {
		outputs = append(outputs, &TxOutput{out.Value, out.PubKeyHash})
	}

	txCopy := &Transaction{tx.TxID, inputs, outputs}
	return txCopy

}

func (tx *Transaction) Serialize() [] byte {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	return buf.Bytes()
}

func (tx *Transaction) NewTxID() []byte {
	txCopy := tx
	txCopy.TxID = []byte{}
	hash := sha256.Sum256(txCopy.Serialize())
	return hash[:]
}
