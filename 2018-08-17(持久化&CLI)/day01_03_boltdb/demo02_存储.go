package main

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
)

func main() {
	//step1：打开数据库
	//step2：Update()
	db,err:=bolt.Open("my.db",0600,nil)
	if err != nil{
		log.Panic(err)
	}
	defer db.Close()

	//存储一个数据[]byte
	err = db.Update(func(tx *bolt.Tx) error {
		//存储数据
		//1.创建表
		b,err:=tx.CreateBucketIfNotExists([]byte("mybucket"))
		if err != nil{
			log.Panic(err) //数据表创建有误
		}
		//2.存储数据
		if b != nil{
			//err = b.Put([]byte("k"),[]byte("send 100 to wangergou"))
			err = b.Put([]byte("k"),[]byte("send 100 to lixiaohua"))
			b.Put([]byte("name"),[]byte("wangergou"))
			if err != nil{
				fmt.Println("数据存储有误")
			}
		}



		return nil
	})

	if err != nil{
		log.Panic(err)
	}
}
