package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	merkle "lab1-blockchain-v1.1.0/internal/security"
	"lab1-blockchain-v1.1.0/internal/utils"
)

type Block struct {
	Timestamp     int64          `json:"timestamp"`
	PrevBlockHash []byte         `json:"prevBlockHash"`
	Hash          []byte         `json:"hash"`
	MerkleRoot    []byte         `json:"merkleRoot"`
	Transactions  []*Transaction `json:"transactions"`
}

// Create a new block with a transactions and prevBlockHash
func NewBlock(
	transactions []*Transaction,
	prevBlockHash []byte,
) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Transactions:  transactions,
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		MerkleRoot:    []byte{},
	}

	block.Hash = block.CalculateHash()

	tree := block.CalculateMerkleRoot()
	block.MerkleRoot = tree.RootNode.Data

	return block
}

// CalculateHash calculates the hash of the block.
//
// It returns the hash of the block as a byte slice.
func (block *Block) CalculateHash() []byte {
	headers := [][]byte{
		utils.IntToHex(block.Timestamp),
		block.PrevBlockHash,
	}
	headers = append(headers, block.HashTransactions())
	hash := sha256.Sum256(bytes.Join(headers, []byte{}))

	return hash[:]
}

// GetHashTransactions calculates the hash of all the transactions in the block.
//
// It returns the hashes of the transactions as a slice of byte slices.
func (block *Block) GetHashTransactions() [][]byte {
	var txHashes [][]byte

	// append the hashes into list transactions
	for _, tx := range block.Transactions {
		txHashes = append(txHashes, tx.Data)
	}
	return txHashes[:]
}

// HashTransactions calculates the hash of all the transactions in the block.
//
// It returns the hash of the transactions as a byte slice.
func (block *Block) HashTransactions() []byte {
	txHashes := block.GetHashTransactions()
	txHash := sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}

// CalculateMerkleRoot calculates the merkle root of the transactions in the block.
//
// It returns a MerkleTree object representing the merkle root.
func (block *Block) CalculateMerkleRoot() *merkle.MerkleTree {
	txHashes := block.GetHashTransactions()
	tree := merkle.CreateMerkleTree(txHashes)

	return tree
}

// Verify checks if the merkle root of the block matches the calculated merkle root.
//
// It returns true if the merkle roots match, false otherwise.
func (block *Block) Verify() bool {

	merkleRootHash := block.CalculateMerkleRoot().RootNode.Data

	if bytes.Equal(merkleRootHash, block.MerkleRoot) {
		return true
	}

	return false
}

// ShowMerkleTree displays the merkle tree of the block.
func (block *Block) ShowMerkleTree() {
	tree := block.CalculateMerkleRoot()
	displayMerkleNode(tree.RootNode, "", true)
}

func displayMerkleNode(node *merkle.MerkleNode, prefix string, isTail bool) {
	if node == nil {
		return
	}

	// Determine the appropriate prefix for the current node
	indent := prefix + "├── "
	if isTail {
		indent = prefix + "└── "
	}

	fmt.Printf("%s%s\n", indent, hex.EncodeToString(node.Data))
	// Determine the new prefix for the child nodes
	childPrefix := prefix + "    "
	if !isTail {
		childPrefix = prefix + "│   "
	}
	if node.Left != nil || node.Right != nil {
		if node.Left != nil {
			displayMerkleNode(node.Left, childPrefix, node.Right == nil)
		}
		if node.Right != nil {
			displayMerkleNode(node.Right, childPrefix, true)
		}
	}
}

// Authors: https://github.com/NoNameNo1F/Simplified_Blockchain-Golang
