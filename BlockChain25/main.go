package main

import "./BLC"

func main() {
	//blockchain := BLC.CreateBlockchainWithGenesisBlock()
	//defer blockchain.DB.Close()

	cli := &BLC.CLI{}

	cli.Run()

}

//bc ./bc addBlock -data "zhoujiafeng"

//b c ./bc printchain 即将输出所有block

//不带-：
