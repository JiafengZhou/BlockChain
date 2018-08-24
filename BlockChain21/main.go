package main

import (
	"flag"
	"fmt"
)

func main() {

	flagPrintChainCmd := flag.String("printchain", "即将打印区块......", "输出所有block.")

	flagInt := flag.Int("number", 6, "输出一个整数")

	flag.Parse()

	fmt.Printf("%s\n", *flagPrintChainCmd)

	fmt.Printf("%s\n", *flagInt)

}

//bc ./bc addBlock -data "zhoujiafeng"

//bc ./bc printchain 即将输出所有block

//不带-：
