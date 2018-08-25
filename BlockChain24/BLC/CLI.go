package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\taddblock -data DATA - 交易数据")
	fmt.Println("\tprintchain -- 输出区块链信息")

}

func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

type CLI struct {
	Blockchain *Blockchain
}

func (cli *CLI) addBlock(data string) {
	cli.Blockchain.AddBlockToBlockchain(data)
}

func (cli *CLI) printChain() {
	cli.Blockchain.PrintchainWithIterator()
}

func (cli *CLI) Run() {
	isValidArgs()

	args := os.Args

	fmt.Printf("%v\n", args)

	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	flagAddBlockData := addBlockCmd.String("data", "https://github.com/JiafengZhou", "交易数据")
	//flagprintchain := addBlockCmd.String("data", "https://github.com/JiafengZhou", "交易数据")

	switch os.Args[1] {
	case "addBlock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		fmt.Println(*flagAddBlockData)
		if *flagAddBlockData == "" {
			printUsage()
			os.Exit(1)
		}

		//fmt.Println(*flagAddBlockData)
		cli.addBlock(*flagAddBlockData)
	}

	if printChainCmd.Parsed() {
		fmt.Println("输出所有区块的数据")
		//fmt.Println(*printChainCmd)
		cli.printChain()
	}

}
