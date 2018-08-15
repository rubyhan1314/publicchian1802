package main

import (
	"bytes"
	"fmt"
)

func main() {
	/*
	数组拼接
		bytes.Join(s [][] byte,sep[])
		第一个参数：钥拼接的二位数组
		第二个参数：分割
	 */
	 b1:=[]byte("hello")
	 b2:=[]byte("abcd")
	 b3:=[]byte("ABCD")

	 s:=[][]byte{
	 	b1,
	 	b2,
	 	b3}
	fmt.Println(s)
	b4:= bytes.Join(s,[]byte{})
	fmt.Println(b4)
}
