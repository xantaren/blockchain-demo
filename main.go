package main

import (
	"blockchaindemo/blockchain"
	"os"
)

func main() {
	chain := blockchain.InitBlockChain()
	defer os.Exit(0)
	defer chain.Database.Close()

	cli := blockchain.CommandLine{
		Blockchain: chain,
	}
	cli.Run()
}
