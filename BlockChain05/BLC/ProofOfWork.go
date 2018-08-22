package BLC

import "math/big"

const targetBit = 16

//hash：32bytes,256Bit
//0000 0010 1000 0000 1000 1001 1001 1001
//0000 0011 1111 1111 1111 1111   1111
//6 * 8 = 48 难度值
type ProofOfWork struct {
	Block  *Block   //当前要验证的区块
	target *big.Int //大数存储(防止溢出)
}

func (proofOfWork *ProofOfWork) Run() ([]byte, int64) {

	return nil, 0
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
