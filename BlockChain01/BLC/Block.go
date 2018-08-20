package BLC

import "time"

type Block struct {
	//1.区块高度
	Height int64

	//2.上一个区块Hash
	PrevBlockHash []byte

	//3.交易数据
	Data []byte

	//4.时间戳
	Timestamp int64

	//5.Hash值
	Block []byte
}

//1.创建新的区块
func NewBlock(data string, height int64, preBlockHash []byte) *Block {
	block := &Block{height,
		preBlockHash,
		[]byte(data),
		time.Now().Unix(),
		nil,
	}
	return block
}
