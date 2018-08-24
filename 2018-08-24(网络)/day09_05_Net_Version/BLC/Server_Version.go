package BLC

type Version struct {
	Version    int64  //版本
	BestHeight int64  //当前节点的区块链中的最后一个区块的高度
	AddrFrom   string //当前节点自己的地址
}
