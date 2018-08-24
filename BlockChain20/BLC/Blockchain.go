package BLC

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"time"
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

//遍历输出所有区块的信息
func (blc *Blockchain) Printchain() {

	var block *Block
	b := 1
	fmt.Println(b)
	var currentHash []byte = blc.Tip

	for {
		err := blc.DB.Update(func(tx *bolt.Tx) error {
			//获取表
			b := tx.Bucket([]byte(blockTableName))
			if b != nil {
				blockBytes := b.Get(currentHash)
				block = DeserializeBlock(blockBytes)
				////1. 区块高度
				//Height int64
				////2. 上一个区块HASH
				//PrevBlockHash []byte
				////3. 交易数据
				//Data []byte
				////4. 时间戳
				//Timestamp int64
				////5. Hash
				//Hash []byte
				//// 6. Nonce
				//Nonce int64
				fmt.Println("Height:", block.Height)
				fmt.Println("PrevBlockHash:", block.PrevBlockHash)
				fmt.Println("Data:", block.Data)
				fmt.Println("Timestamp:", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
				fmt.Println("Hash:", block.Hash)
				fmt.Println("Nonce:", block.Nonce)
				fmt.Println("---------------------------------------------")
			}
			//
			return nil
		})

		if err != nil {
			log.Panic(err)
		}

		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)

		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}
		currentHash = block.PrevBlockHash
	}

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
