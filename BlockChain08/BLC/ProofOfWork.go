package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

const targetBit = 16

//hash：32bytes,256Bit
//0000 0010 1000 0000 1000 1001 1001 1001
//0000 0011 1111 1111 1111 1111 1111 1111
//6 * 8 = 48 难度值
type ProofOfWork struct {
	Block  *Block   //当前要验证的区块
	target *big.Int //大数存储(防止溢出)
}

func (pow *ProofOfWork) prepareData(nonce int64) []byte {
	block := pow.Block

	data := bytes.Join(
		[][]byte{
			block.PrevBlockHash,
			block.Data,
			IntToHex(block.Timestamp),
			IntToHex(targetBit),
			IntToHex(nonce),
			IntToHex(block.Height),
		},
		[]byte{},
	)

	return data
}

func (proofOfWork *ProofOfWork) Run() ([]byte, int64) {

	//1.将block属性拼接成字节数组
	var nonce int64 = 0
	var hashInt big.Int
	var hash [32]byte
	for {
		//准备hash
		dataBytes := proofOfWork.prepareData(nonce)
		hash = sha256.Sum256(dataBytes)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if proofOfWork.target.Cmp(&hashInt) == 1 {
			break
		}

		nonce = nonce + 1

	}
	//2.计算hash值并验证

	//3.判断hash有效性，不满足则改变nonce值
	fmt.Println()
	fmt.Println("Nonce :", nonce)
	return hash[:], nonce
}

func NewProofOfWork(block *Block) *ProofOfWork {
	//1.big.Int对象 1 << 256 - target
	//只要小于这个值
	//1.创建一个初始值为1的target
	//2.左
	target := big.NewInt(1)

	target = target.Lsh(target, 256-targetBit)

	return &ProofOfWork{block, target}
}

func (pow *ProofOfWork) IsValid() bool {
	var hashInt big.Int
	hashInt.SetBytes(pow.Block.Hash)
	if pow.target.Cmp(&hashInt) == 1 {
		return true
	} else {
		return false
	}
}
