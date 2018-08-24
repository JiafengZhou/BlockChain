package BLC

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"time"
)

type BlockchainIterator struct {
	CurrentHash []byte   //当前正在遍历的区块hash
	DB          *bolt.DB //数据库
}

//迭代器
func (blockchain *Blockchain) Iterator() *BlockchainIterator {

	return &BlockchainIterator{blockchain.Tip, blockchain.DB}
}

func (blockchainIterator *BlockchainIterator) Next() *Block {
	//第一次调用，先将最新区块返回，并且更新迭代器数据结构
	var block *Block
	err := blockchainIterator.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))

		if b != nil {
			block = DeserializeBlock(b.Get(blockchainIterator.CurrentHash))
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	blockchainIterator.CurrentHash = block.PrevBlockHash
	return block
}

func (blc *Blockchain) PrintchainWithIterator() {

	iter := blc.Iterator()
	for {

		block := iter.Next()

		fmt.Println("Height:", block.Height)
		fmt.Printf("PrevBlockHash:%x\n", block.PrevBlockHash)
		fmt.Printf("Data:%s\n", block.Data)
		fmt.Println("Timestamp:", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("Hash:%x\n", block.Hash)
		fmt.Println("Nonce:", block.Nonce)
		fmt.Println("---------------------------------------------")

		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)

		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}
	}

}
