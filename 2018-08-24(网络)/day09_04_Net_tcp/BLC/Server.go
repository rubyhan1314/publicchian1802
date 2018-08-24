package BLC

import (
	"fmt"
	"net"
	"log"
	"io/ioutil"
	"io"
	"bytes"
)

/*
全节点：地址硬编码
	localhost：3000

钱包节点，矿工节点

 */

var knowNodes = []string{"localhost:3000"}

//节点的服务端启动，可以接收其他节点传来的数据。。
func startServer(nodeID string, mineAddress string) {
	//当前节点自己地址："192.168.1.100:3001"
	nodeAddress := fmt.Sprintf("localhost:%s", nodeID)

	//监听：地址
	listen, err := net.Listen("tcp", nodeAddress)
	if err != nil {
		log.Panic(err)
	}
	defer listen.Close()

	//当前的节点，如果全节点：等待
	//如果不是全节点，钱包节点，矿工节点，给全节点发送一个消息。。
	if nodeAddress != knowNodes[0] {
		//钱包节点，矿工节点。。给全节点发送一个消息。。
		sendMessage(knowNodes[0], "我时王二狗，我的地址是："+nodeAddress)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Panic(err)
		}

		fmt.Println("发送方已经链入。。", conn.RemoteAddr())

		request, err := ioutil.ReadAll(conn)
		if err != nil {
			log.Panic(err)
		}

		fmt.Printf("接收到的数据是：%s\n", request) //command + 数据
	}

}

//发送消息
func sendMessage(to, data string) {
	fmt.Println("当前节点可以发送数据。。。")
	conn, err := net.Dial("tcp", to)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	//发送数据。。
	_, err = io.Copy(conn, bytes.NewReader([]byte(data)))
	if err != nil {
		log.Panic(err)
	}

}
