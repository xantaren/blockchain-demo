package blockchain

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
)

type CommandLine struct {
	Blockchain *BlockChain
}

func (cli *CommandLine) printUsage() {
	fmt.Println("Usage:")
}

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()

		// Exits by shutting down the goroutine unlike os.exit() which exits mid-process
		// This is to promote proper GC and prevent corrupting the db
		runtime.Goexit()
	}
}

func (cli *CommandLine) addBlock(data string) {
	cli.Blockchain.AddBlock(data)
	fmt.Println("Block Added")
}

func (cli *CommandLine) printBlockChainInfo() {
	iter := cli.Blockchain.Iterator()

	for {
		block := iter.Next()

		fmt.Println("------")
		fmt.Println("Data in Block: ", string(block.Data))
		fmt.Println("Hash: ", hex.EncodeToString(block.Hash))
		fmt.Println("Previous Hash: ", hex.EncodeToString(block.PrevHash))

		pow := NewProof(block)
		fmt.Println("Proof of work: ", pow.Validate())

		// The genesis block does not contain a previous block, thus no previous hash
		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func (cli *CommandLine) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)
	addBlockData := addBlockCmd.String("block", "", "Block data")

	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panicln(err)
		}

	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panicln(err)
		}

	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printBlockChainInfo()
	}
}
