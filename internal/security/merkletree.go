package security

import (
	"crypto/sha256"
)

type MerkleTree struct {
	RootNode *MerkleNode `json:"rootNode"`
}

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

// CreateMerkleNode creates a new MerkleNode by hashing the data and setting the left and right child nodes.
//
// Parameters:
// - left: the left child node of the new MerkleNode.
// - right: the right child node of the new MerkleNode.
// - data: the data to be hashed and used as the data of the new MerkleNode.
//
// Returns:
// - The newly created MerkleNode.
func CreateMerkleNode(
	left, right *MerkleNode,
	data []byte,
) *MerkleNode {
	node := &MerkleNode{}
	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		node.Data = hash[:]
	} else {
		prevHashes := append(left.Data, right.Data...)
		hash := sha256.Sum256(prevHashes)
		node.Data = hash[:]
	}

	node.Left = left
	node.Right = right

	return node
}

// CreateMerkleTree creates a Merkle tree from the given array of data.
//
// Parameter:
// datas: a 2D byte array containing the data for each node.
// Return:
// A pointer to the root of the Merkle tree.
func CreateMerkleTree(datas [][]byte) *MerkleTree {
	var nodes []MerkleNode
	// creating node with a specific data of transaction
	for _, data := range datas {
		node := CreateMerkleNode(nil, nil, data)
		nodes = append(nodes, *node)
	}

	// Adding 1 node if node standalone
	if len(nodes)%2 != 0 {
		nodes = append(nodes, nodes[len(nodes)-1])
	}
	// for i := 0; i < len(nodes)/2; i++
	for len(nodes) > 1 {
		var newLevel []MerkleNode

		for j := 0; j < len(nodes); j += 2 {
			left := &nodes[j]
			right := &nodes[j+1]
			node := CreateMerkleNode(left, right, nil)
			newLevel = append(newLevel, *node)
		}

		if len(newLevel)%2 != 0 && len(newLevel) != 1 {
			newLevel = append(newLevel, newLevel[len(newLevel)-1])
		}
		nodes = newLevel
	}

	tree := MerkleTree{RootNode: &nodes[0]}
	return &tree
}

// Authors: https://github.com/NoNameNo1F/Simplified_Blockchain-Golang
