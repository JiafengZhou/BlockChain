package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	//创建或者打开数据库
	//
	db, err := bolt.Open("my.db", 0600, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	//创建表

	err = db.Update(func(tx *bolt.Tx) error {
		//创建BlockBucket表
		b, err := tx.CreateBucket([]byte("BlockBucket"))

		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		if b != nil {
			err := b.Put([]byte("l"), []byte("Send 100 BTC To 强哥......"))

			if err != nil {
				log.Panic("数据存储失败......")
			}
		}

		return nil
	})
}
