package blockchaindemo

import (
	"encoding/hex"
	"fmt"
)

func printBlockChainInfo(chain *BlockChain) {
	for _, block := range chain.Blocks {
		fmt.Println("Data in Block: ", string(block.Data))
		fmt.Println("Hash: ", hex.EncodeToString(block.Hash))
		fmt.Println("Previous Hash: ", hex.EncodeToString(block.PrevHash))
		fmt.Println("------")
	}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First block after genesis")
	chain.AddBlock("Another block")
	chain.AddBlock("And another")

	printBlockChainInfo(chain)
}
