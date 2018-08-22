package main

import (
	"./BLC"
	"fmt"
)

func main() {
	genesisBlockChain := BLC.CreateBlockchainWithGenesisBlock()
	fmt.Println(genesisBlockChain)
	fmt.Println(genesisBlockChain.Blocks)
	fmt.Println(genesisBlockChain.Blocks[0])
}
