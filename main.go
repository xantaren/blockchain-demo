package main

import (
	"blockchaindemo/blockchain"
	"encoding/hex"
	"fmt"
)

func printBlockChainInfo(chain *blockchain.BlockChain) {
	for _, block := range chain.Blocks {
		fmt.Println("------")
		fmt.Println("Data in Block: ", string(block.Data))
		fmt.Println("Hash: ", hex.EncodeToString(block.Hash))
		fmt.Println("Previous Hash: ", hex.EncodeToString(block.PrevHash))

		pow := blockchain.NewProof(block)
		fmt.Println("Proof of work: ", pow.Validate())
	}
}

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First block after genesis")
	chain.AddBlock("Another block")
	chain.AddBlock("And another")

	printBlockChainInfo(chain)
}
