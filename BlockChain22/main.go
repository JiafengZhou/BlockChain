package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args

	fmt.Printf("%v\n", args)

}

//bc ./bc addBlock -data "zhoujiafeng"

//bc ./bc printchain 即将输出所有block

//不带-：
