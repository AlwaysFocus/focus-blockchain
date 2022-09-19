package blockchain

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
	// The nonce for the block
	Nonce int
}

// CreateBlock - This function creates a new block using the data and previous hash
func CreateBlock(data string, prevHash []byte) *Block {
	// Create the new Block
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	// Derive the hash for the new block
	// Get a new proof of work for our block
	pow := NewProof(block)

	// Get our Nonce and Hash for the new block
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

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
