package BLC


/*
用于获取某个区块，或者某个交易的数据
 */

 type GetData struct {
 	AddrFrom string//当前节点自己的地址
 	Type string //要获取的数据类型：block，tx
 	Hash []byte//块的hash，tx的hash
 }
