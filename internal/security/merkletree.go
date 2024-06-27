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
