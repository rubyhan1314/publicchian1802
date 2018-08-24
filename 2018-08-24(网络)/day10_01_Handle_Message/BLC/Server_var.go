package BLC

/*
全节点：地址硬编码
	localhost：3000

钱包节点，矿工节点

 */

var knowNodes = []string{"localhost:3000"}

var nodeAddress string //当前节点自己的地址


//记录因该同步，但尚未同步的区块的hash
var blocksArray [][]byte