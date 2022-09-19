package main

import (
	"fmt"
	"github.com/alwaysfocus/focus-blockchain/blockchain"
	"strconv"
)

func main() {
	// Create the initial Blockchain
	chain := blockchain.InitBlockChain()

	// Add a few blocks
	chain.AddBlock("First Block after Genesis Block")
	chain.AddBlock("Second Block after Genesis Block")
	chain.AddBlock("Third Block after Genesis Block")

	// Let's test out our blockchain :)
	for _, block := range chain.Blocks {
		fmt.Printf("Prev Hash: %x\n", block.PrevHash)
		fmt.Printf("Block Data: %s\n", block.Data)
		fmt.Printf("Block Hash: %x\n", block.Hash)

		// Run our proof of work
		pow := blockchain.NewProof(block)

		// Is the block valid?
		fmt.Printf("Proof of Work: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

	}
}
