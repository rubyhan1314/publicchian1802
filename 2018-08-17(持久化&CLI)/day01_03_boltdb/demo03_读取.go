package main

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
)

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	//读取数据：
	err = db.View(func(tx *bolt.Tx) error {
		//打开桶
		b := tx.Bucket([]byte("mybucket"))

		if b != nil {
			//读取数据
			data := b.Get([]byte("k")) //[]byte
			fmt.Println(data)
			fmt.Println(string(data))

			data2 := b.Get([]byte("kk"))
			fmt.Println(data2)
		}
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

}
