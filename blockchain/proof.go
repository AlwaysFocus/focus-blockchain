package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

// Steps:
// Take data from a block
// Create a nonce (counter) which starts at 0
// Create a hash of the data with the nonce
// Check to see if the hash matches our set of requirements.
// Requirements:
// First bytes must contain 0's

// Difficulty - The difficulty to calculate the hash is the number of leading 0's required. We will use a constant for this
const Difficulty = 18

// ProofOfWork - This struct will contain the block and the target
type ProofOfWork struct {
	// The block in the blockchain
	Block *Block
	// A Number that satisfies the requirements previously listed - it is derived from our difficulty
	Target *big.Int
}

// NewProof - Produce a pointer to a proof of work for a given block
func NewProof(b *Block) *ProofOfWork {
	// Our new target
	target := big.NewInt(1)
	// Subtract the difficulty from the number of bytes in a hash then shift bytes left by the result
	target.Lsh(target, uint(256-Difficulty))

	// Our proof of work
	pow := &ProofOfWork{b, target}

	return pow
}

// InitData - This will be our Hash derivation function
func (pow *ProofOfWork) InitData(nonce int) []byte {
	// Generate our data for a new block using the preHash, Data, nonce and Difficulty
	data := bytes.Join([][]byte{
		pow.Block.PrevHash,
		pow.Block.Data,
		ToHex(int64(nonce)),
		ToHex(int64(Difficulty)),
	}, []byte{})

	return data
}

// Run - Our main computation function
func (pow *ProofOfWork) Run() (int, []byte) {
	// Our integer used for hashing
	var intHash big.Int
	// Our hash
	var hash [32]byte

	// Our nonce
	nonce := 0

	// Do our proof of work
	for nonce < math.MaxInt64 {
		// Our Data
		data := pow.InitData(nonce)
		// Our Hash
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)

		// Convert our hash into a bigint
		intHash.SetBytes(hash[:])

		// See if our bigInt hash matches our pow target
		if intHash.Cmp(pow.Target) == -1 {
			// Hash is less than our target - we have signed the block
			break
		} else {
			// Keep doing work!
			nonce++
		}
	}
	fmt.Println()
	return nonce, hash[:]
}

// Validate - Used to validate that our proof of work was successful and our hash is valid
func (pow *ProofOfWork) Validate() bool {
	// BigInt to hold the int version of our hash
	var intHash big.Int

	// Initialize the data for the pow using the nonce of our block
	data := pow.InitData(pow.Block.Nonce)

	// Get our hash
	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	// Run int hash computation to see if we have a valid block
	return intHash.Cmp(pow.Target) == -1

}

// ToHex - Utility function for converting an int64 to a BigEndian []byte
func ToHex(num int64) []byte {
	// Create a new buffer to write our number to
	buff := new(bytes.Buffer)

	// Write our int64 in BigEndian to our buffer
	err := binary.Write(buff, binary.BigEndian, num)

	// Catch any errors
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
