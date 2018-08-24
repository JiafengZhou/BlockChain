package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

//数据库名字
const dbName = "blockchain.db"

//表名
const blockTableName = "blocks"

type Blockchain struct {
	//Blocks []*Block // 存储有序的区块
	Tip []byte //最新的区块的Hash值
	DB  *bolt.DB
}

// 增加区块到区块链里面
func (blc *Blockchain) AddBlockToBlockchain(data string) {

	// 往链里面添加区块
	//blc.Blocks = append(blc.Blocks, newBlock)
	err := blc.DB.Update(func(tx *bolt.Tx) error {
		//获取表
		b := tx.Bucket([]byte(blockTableName))

		if b != nil {
			// 创建新区块
			var height int64 = DeserializeBlock(b.Get(blc.Tip)).Height + 1
			var preHash []byte = DeserializeBlock(b.Get(blc.Tip)).Hash
			newBlock := NewBlock(data, height, preHash)
			//表里添加区块
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {

				log.Panic(err)
			}

			// 序列化更新最新区块
			err = b.Put([]byte("l"), newBlock.Hash)
			if err != nil {
				log.Panic(err)
			}

			//更新Tip
			blc.Tip = newBlock.Hash
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

}

//1. 创建带有创世区块的区块链
func CreateBlockchainWithGenesisBlock() *Blockchain {

	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	var blockHash []byte

	err = db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(blockTableName))

		if b == nil {

			b, err = tx.CreateBucket([]byte(blockTableName))
		}

		if err != nil {
			log.Panic(err)
		}

		if b != nil {
			// 创建创世区块
			genesisBlock := CreateGenesisBlock("Genesis Data.......")

			//将创世区块存储到表中
			err := b.Put(genesisBlock.Hash, genesisBlock.Serialize())

			if err != nil {
				log.Panic(err)
			}
			//更新最新的区块链hash值
			err = b.Put([]byte("l"), genesisBlock.Hash)

			if err != nil {
				log.Panic(err)
			}

			blockHash = genesisBlock.Hash
		}
		return nil
	})

	// 返回区块链对象
	return &Blockchain{blockHash, db}
}
