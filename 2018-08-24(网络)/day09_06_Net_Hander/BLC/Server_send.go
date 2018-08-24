package BLC

import (
	"fmt"
	"net"
	"log"
	"io"
	"bytes"
)
/*
处理节点之间的发送的数据
 */


//发送version
func sendVersion(toAddr string, bc *BlockChain) {
	//1.获取当前区块链的队在高高度
	bestHeight := bc.GetBestHeight()
	//创建version对象
	version := Version{NODE_VERSION, bestHeight, nodeAddress}
	//将version序列化
	payload := gobEncode(version)

	//拼接命令+数据
	request := append(commandToBytes(COMMAND_VERSION), payload...)

	//发送
	sendData(toAddr, request)
}

//发送消息
func sendData(to string, data []byte) {
	fmt.Println("当前节点可以发送数据。。。")
	conn, err := net.Dial("tcp", to)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	//发送数据。。
	_, err = io.Copy(conn, bytes.NewReader(data))
	if err != nil {
		log.Panic(err)
	}

}

