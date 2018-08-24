package main

import (
	"fmt"
	"net"
	"log"
	"io/ioutil"
)

func main() {
	/*
	服务端：
	 */
	fmt.Println("服务端的程序。。。")
	listen, err := net.Listen("tcp", ":9528")
	if err != nil {
		log.Panic(err)
	}

	defer listen.Close()

	//等待客户端的链接申请。。
	for{
		conn,err:=listen.Accept()//阻塞式
		if err != nil{
			log.Panic(err)
		}

		fmt.Println("已有客户端链接。。",conn.RemoteAddr())
		//读取对方传来的数据
		request,err:=ioutil.ReadAll(conn)
		if err != nil{
			log.Panic(err)
		}

		fmt.Printf("接收到的数据是：%s\n",request)


	}

}
