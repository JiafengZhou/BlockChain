package BLC

type BlockChain struct {
	Blocks []*Block //存储有序的区块
}

//1.创建带有创世区块的区块链

func CreateBlockchainWithGenesisBlock() *BlockChain {
	//创建创世区块
	genesisBlock := CreateGenesisBlock("Genesis Data.....")
	//返回区块数组
	return &BlockChain{[]*Block{genesisBlock}}

}

//增加区块到区块里面
func (blc *BlockChain) AddBlockToBlockchain(data string, height int64, preHash []byte) {
	//创建新区块
	newBlock := NewBlock(data, height, preHash)
	//添加区块到区块链
	blc.Blocks = append(blc.Blocks, newBlock)

}
