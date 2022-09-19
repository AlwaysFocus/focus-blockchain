package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// BlockChain - Contains the array of pointers to all blocks. Later on we will implement more complicated structures
type BlockChain struct {
	Blocks []*Block
}

// Block - This is our base block for our blockchain
type Block struct {
	// Represents the hash of a given block
	// A Hash is derived using the Data and the Hash of the previous block
	Hash []byte
	// Represents the data of a given block
	Data []byte
	// Represents the previous hash of a given block
	// This lets us link blocks together
	PrevHash []byte
}

// DeriveHash - This function creates the hash for a given block using the Data and the previous Hash
func (b *Block) DeriveHash() {
	// Combine the Data and the PrevHash into a new byte slice
	joinedData := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	// Create a new hash using the joined data. We are going to use SHA256 for now but will implement other algorithms later
	hash := sha256.Sum256(joinedData)
	// Set the hash of the block to the new hash
	b.Hash = hash[:]
}

// CreateBlock - This function creates a new block using the data and previous hash
func CreateBlock(data string, prevHash []byte) *Block {
	// Create the new Block
	block := &Block{[]byte{}, []byte(data), prevHash}
	// Derive the hash for the new block
	block.DeriveHash()
	// return the new block
	return block
}

// AddBlock - This function is used to add new blocks to the blockchain
func (chain *BlockChain) AddBlock(data string) {
	// Get the previous block
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	// Create the current block
	newBlock := CreateBlock(data, prevBlock.Hash)
	// Add the new block to the blockchain
	chain.Blocks = append(chain.Blocks, newBlock)
}

// Genesis - The Genesis function is responsible for creating the first block in a new blockchain
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain - Initialize the blockchain with a Genesis block
func InitBlockChain() *BlockChain {
	// Initialize the blockchain
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	// Create the initial Blockchain
	chain := InitBlockChain()

	// Add a few blocks
	chain.AddBlock("First Block after Genesis Block")
	chain.AddBlock("Second Block after Genesis Block")
	chain.AddBlock("Third Block after Genesis Block")

	// Let's look at the data in the blockchain
	for _, block := range chain.Blocks {
		fmt.Printf("Prev Hash: %x\n", block.PrevHash)
		fmt.Printf("Block Data: %s\n", block.Data)
		fmt.Printf("Block Hash: %x\n", block.Hash)
	}
}
