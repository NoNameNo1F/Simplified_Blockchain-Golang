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

func (block *Block) CalculateHash() []byte {
	headers := [][]byte{
		utils.IntToHex(block.Timestamp),
		block.PrevBlockHash,
	}
	headers = append(headers, block.HashTransactions())
	hash := sha256.Sum256(bytes.Join(headers, []byte{}))

	return hash[:]
}

/*
Summary: Function calculates Hash of all transactions are stored in a block into byte[] for representing all transactions

- Input: None
- Output: byte[]

Usage: block.HashTransactions()
*/
func (block *Block) GetHashTransactions() [][]byte {
	var txHashes [][]byte

	// append the hashes into list transactions
	for _, tx := range block.Transactions {
		txHashes = append(txHashes, tx.Data)
	}
	return txHashes[:]
}

func (block *Block) HashTransactions() []byte {
	txHashes := block.GetHashTransactions()
	txHash := sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}

func (block *Block) CalculateMerkleRoot() *merkle.MerkleTree {
	txHashes := block.GetHashTransactions()
	tree := merkle.CreateMerkleTree(txHashes)

	return tree
}

func (block *Block) Verify() bool {

	merkleRootHash := block.CalculateMerkleRoot().RootNode.Data

	if bytes.Equal(merkleRootHash, block.MerkleRoot) {
		return true
	}

	return false
}

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
