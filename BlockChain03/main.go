package main

import (
	"./BLC"
	"fmt"
)

func main() {
	genesisBlock := BLC.CreateGenesisBlock("Genesis Block...")
	//block := BLC.NewBlock("Genenis",1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
	fmt.Println(genesisBlock)
	fmt.Printf("%v", genesisBlock.Hash)
}
