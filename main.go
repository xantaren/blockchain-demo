package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

type BlockChain struct {
	blocks []*Block
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func printBlockChainInfo(chain *BlockChain) {
	for _, block := range chain.blocks {
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
