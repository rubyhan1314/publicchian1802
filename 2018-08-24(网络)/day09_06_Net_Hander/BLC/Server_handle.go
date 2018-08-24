package BLC

import (
	"bytes"
	"encoding/gob"
	"log"
	"fmt"
)

//处理接收到的version数据
func handleVersion(request [] byte, bc *BlockChain) {
	/*
	1.从request中获取版本的数据：[]byte
	2.反序列化--->version
	3.操作bc，获取自己的最后block的height
	4.根对方的比较
	 */

	versionBytes := request[COMMAND_LENGTH:]
	//反序列化
	var version Version
	reader := bytes.NewReader(versionBytes)
	decoder := gob.NewDecoder(reader)
	err := decoder.Decode(&version)
	if err != nil {
		log.Panic(err)
	}

	//获取自己的block的height
	selfHeight := bc.GetBestHeight() // 0
	foreignerBestHeight := version.BestHeight // 1

	fmt.Printf("接收到%s,传来的版本高度:%d\n", version.AddrFrom, foreignerBestHeight)
	if selfHeight > foreignerBestHeight {
		//发送版本对象给 对方
		sendVersion(version.AddrFrom, bc)
	} else if selfHeight < foreignerBestHeight {
		fmt.Println("我的高度没有你高，给我看看你的数据。。。")
	}

}
