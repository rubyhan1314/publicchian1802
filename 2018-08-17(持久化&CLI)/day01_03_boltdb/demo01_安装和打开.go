package main

import (
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	/*
	step1：安装
		go get "github.com/boltdb/bolt"

	step2：打开一个数据库

	go run xxx.go
		--->直接执行，相对于当前的工程下

	go build -o 别名 xxx.go

	 */
	 db,err:=bolt.Open("my.db",0600,nil)
	 if err != nil{
	 	log.Panic(err)
	 }
	 defer db.Close()

	 /*
	 db.Update(),读和写
	 db.View()，只读


	  */

}
