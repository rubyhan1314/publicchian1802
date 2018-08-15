package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"fmt"
)

func main() {
	var num int64 = 256
	fmt.Printf("十进制：%d\n",num)
	fmt.Printf("二进制：%b\n",num)
	fmt.Printf("十六进制：%x\n",num)

	bytes:=IntToHex(num)
	fmt.Println(bytes)

}


/*
将一个int64的整数：转为二进制后，每8bit一个byte。转为[]byte
 */
func IntToHex(num int64)[]byte{ //uint8-->0-255
	buff :=new(bytes.Buffer)
	//将二进制数据写入w
	//func Write(w io.Writer, order ByteOrder, data interface{}) error
	err:=binary.Write(buff,binary.BigEndian,num)
	if err!=nil{
		log.Panic(err)
	}
	//转为[]byte并返回
	return buff.Bytes()
}