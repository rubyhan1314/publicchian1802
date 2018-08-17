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
	err = db.View(func(tx *bolt.Tx) error {
		/*
		游标：Cursor
		 */
		b := tx.Bucket([]byte("mybucket"))
		if b != nil {
			//获取该表中的所有的key-value对
			c := b.Cursor()
			for k, v := c.First(); k != nil; k, v = c.Next() {
				fmt.Printf("key:%s,value:%s\n", k, v)
			}
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

}
