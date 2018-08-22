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
