package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	/*
	Base64编码，go语言提供了base64的编码和解码实现
		编码---->编码后的数据
		解码---->原始数据

	 */

	//一：普通的数据
	//1.原始数据：
	msg := "Man2"
	/*
	Base64编码原理：
		1.字符，按照ASCII码转为数值：10进制
		2.10进制转为2进制-->8bit
		3.按照2进制，6位一组
		4.高位补0,转为10进制



010011	010110	000101	101110
19		22		5		46
T		W		F		u
	 */
	 fmt.Println([]byte(msg)) //[77 97 110]
	 for _,c:=range msg{
	 	fmt.Printf("%b\n",c)
	 }



















	//编码
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println("Base64编码后：", encoded)

	//解码
	decoded,err:=base64.StdEncoding.DecodeString(encoded)
	if err != nil{
		log.Panic(err)
	}
	fmt.Println("Base64解码后：",decoded,string(decoded))

	//二：网址数据
	// 1.原始数据：
	msg2 := "http://www.baidu.com" //网址
	encoded2:=base64.URLEncoding.EncodeToString([]byte(msg2)) //string
	fmt.Println("Base64编码url：",encoded2) //aHR0cDovL3d3dy5iYWlkdS5jb20=

	decoded2,err:=base64.URLEncoding.DecodeString(encoded2)
	if err != nil{
		log.Panic(err)
	}
	fmt.Println("Base64解码url：",decoded2,string(decoded2))
}
