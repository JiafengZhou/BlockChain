## 工作量证明

原理：寻找合适的Nonce值使得Hash值满足一定要求


```go
func (proofOfWork *ProofOfWork) Run() ([]byte, int64) {

	//1.将block属性拼接成字节数组
	var nonce int64 = 0
	var hashInt big.Int
	var hash [32]byte
	for {
		//准备hash
		dataBytes := proofOfWork.prepareData(nonce)
		hash = sha256.Sum256(dataBytes)
		fmt.Printf("\r%x",hash)
		hashInt.SetBytes(hash[:])

		if proofOfWork.target.Cmp(&hashInt) == 1 {
			break
		}

			nonce = nonce + 1

	}
	//2.计算hash值并验证


	//3.判断hash有效性，不满足则改变nonce值

	fmt.Println("Nonce :",nonce)
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
```

























