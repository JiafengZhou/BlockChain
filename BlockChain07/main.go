package main

import (
	"./BLC"
	"fmt"
	"time"
)

func main() {
	blockChain := BLC.CreateBlockchainWithGenesisBlock()

	blockChain.AddBlockToBlockchain("Send 100RMB To zhangqiang", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	blockChain.AddBlockToBlockchain("Send 200RMB To changjingkong", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	blockChain.AddBlockToBlockchain("Send 300RMB To zhangying", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	blockChain.AddBlockToBlockchain("Send 400RMB To zhangying", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	fmt.Println("----------------------------------------------------")

	for id, block := range blockChain.Blocks {
		proofOfWork := BLC.NewProofOfWork(blockChain.Blocks[id])
		fmt.Println(proofOfWork.IsValid())
		fmt.Println(id, ":", block.Nonce)
	}
	time.Sleep(time.Second)
	fmt.Println(blockChain)

}
