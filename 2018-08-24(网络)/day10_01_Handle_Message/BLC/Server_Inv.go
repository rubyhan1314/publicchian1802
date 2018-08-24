package BLC

//传递给对方的数据信息：block，tx
type Inv struct {
	AddrFrom string   //当前节点自己的地址
	Type     string   //要发送数据的种类
	Items    [][]byte //要传递的数据的hash
}
