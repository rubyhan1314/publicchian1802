package main

import (
	"fmt"
	"net"
	"log"
	"io"
	"bytes"
)

func main() {
	fmt.Println("客户端程序申请链接服务端。。。。。")
	conn, err := net.Dial("tcp", "127.0.0.1:9528")
	if err != nil {
		log.Panic(err)
	}

	defer conn.Close()

	//给服务端发送数据。。
	_, err = io.Copy(conn, bytes.NewReader([]byte("helloworld....")))
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("客户端程序结束。。。")
}
