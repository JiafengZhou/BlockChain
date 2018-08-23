package BLC

import (
	"time"
)

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
	Hash []byte

	//6.Nonce值
	Nonce int64
}

//func (block *Block) SetHash() {
//	//1.将Height转化为字节数组
//	heghtBytes := IntToHex(block.Height)
//	fmt.Println("heghtBytes:", heghtBytes)
//	//2.将时间戳转[]byte
//	//base:2代表二进制
//	timeString := strconv.FormatInt(block.Timestamp, 2)
//	fmt.Println("timeString:", timeString)
//
//	timeBytes := []byte(timeString)
//	fmt.Println("timeBytes:", timeBytes)
//	//3.拼接所有属性
//	blockBytes := bytes.Join([][]byte{heghtBytes, block.PrevBlockHash, block.Data, timeBytes, block.Hash}, []byte{})
//	//4.计算Hash值
//	hash := sha256.Sum256(blockBytes)
//
//	block.Hash = hash[:]
//
//}

//1.创建新的区块
func NewBlock(data string, height int64, preBlockHash []byte) *Block {
	block := &Block{height,
		preBlockHash,
		[]byte(data),
		time.Now().Unix(),
		nil,
		0,
	}
	//block.SetHash()
	//工作量证明，返回有效的Hash和Nonce值
	//000000 > hash
	pow := NewProofOfWork(block)
	hash, nonce := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func CreateGenesisBlock(data string) *Block {
	return NewBlock(data, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
