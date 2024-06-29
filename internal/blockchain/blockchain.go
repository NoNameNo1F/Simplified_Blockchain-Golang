package blockchain

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"lab1-blockchain-v1.1.0/internal/utils"
)

type Blockchain struct {
	Blocks []*Block `json:"blocks"`
}

// NewBlockchain creates a new Blockchain instance with an initial block.
//
// It returns a pointer to the newly created Blockchain.
func NewBlockchain() *Blockchain {
	return &Blockchain{Blocks: []*Block{InitBlock()}}
}

// LoadBlockchain loads the blockchain data from a file.
//
// It returns a pointer to the loaded Blockchain.
func LoadBlockchain() *Blockchain {
	//1. get file path
	folderPath := utils.GetPath("data")
	fileName := utils.GetFileByExtens(folderPath, "json")
	//2. get content of file
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("error reading file:", err)
	}
	if len(content) == 0 {
		// fmt.Printf("Blockchain is empty => Created new Blockchain\n")
		return NewBlockchain()
	}

	//3. read content
	var blockchain *Blockchain
	err = json.Unmarshal(content, &blockchain)

	return blockchain
}

func (blockchain *Blockchain) SaveBlockchainToFile() {
	folderPath := utils.GetPath("data")
	fileName := utils.GetFileByExtens(folderPath, "json")

	// bytes, _ := json.Marshal(blockchain)
	bytes, _ := json.MarshalIndent(blockchain, "", "\t")
	// fmt.Println(string(bytes))
	err := os.WriteFile(fileName, bytes, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func InitBlock() *Block {
	return NewBlock(
		[]*Transaction{{Data: []byte("Initialize Block")}},
		[]byte{},
	)
}

func (blockchain *Blockchain) AddBlock(transactions []*Transaction) {
	// 1. Create new block
	// get the previous hash block
	prevBlockHash := blockchain.GetLatestBlock().Hash
	newBlock := NewBlock(transactions, prevBlockHash)
	// 2. Add new block into blockchain
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
	// Logging New Block added
}

// ViewBlockchain prints the blockchain by iterating over each block and printing its headers and transactions.
//
// No parameters.
// No return value.
func (blockchain *Blockchain) ViewBlockchain() {
	for index, block := range blockchain.Blocks {
		fmt.Printf("|--------------------------------------- Block %d ----------------------------------------|\n", index+1)

		fmt.Println("|-------------------- Headers -----------------------------------------------------------|")
		fmt.Printf("|Timestamp: %d                                                                   |\n", block.Timestamp)
		fmt.Printf("|Previous Block Hash: %x   |\n", block.PrevBlockHash)
		fmt.Printf("|Hash: %x                  |\n", block.Hash)
		fmt.Printf("|MerkleRoot: %x            |\n", block.MerkleRoot)

		fmt.Println("|--------------------- Body -------------------------------------------------------------|")
		ViewTransactions(block.Transactions)
		fmt.Println("\n")
	}
}

// GetLatestBlock returns the latest block in the blockchain.
// No parameters.
// Returns a pointer to the latest block.
func (blockchain *Blockchain) GetLatestBlock() *Block {
	return blockchain.Blocks[len(blockchain.Blocks)-1]
}

// VerifyBlock searches for a block in the blockchain by its MerkleRoot hash value and verifies its transactions.
//
// hashValue: a string representing the hash value to search for.
// No return value.
func (blockchain *Blockchain) VerifyBlock(hashValue string) {
	isFound := false
	isVerified := false
	data, err := hex.DecodeString(hashValue)
	if err != nil {
		fmt.Printf("An Error occurred while verifying")
		return
	}

	for _, block := range blockchain.Blocks {
		if bytes.Equal(block.MerkleRoot, data) {
			isFound = true
			isVerified = block.Verify()
		}
	}

	if isFound {
		if !isVerified {
			fmt.Printf("Transactions of Block: %x is not valid\n", data)
			return
		}
		fmt.Printf("Transactions of Block: %x is valid\n", data)
		return
	}

	fmt.Printf("Hash of MerkleRoot is not found\n")
}

// ViewMerkleTree searches for a specific block in the blockchain by its MerkleRoot hash value.
//
// hashValue: a string representing the hash value to search for.
// No return value.
func (blockchain *Blockchain) ViewMerkleTree(hashValue string) {
	isFound := false
	var target *Block

	data, err := hex.DecodeString(hashValue)
	if err != nil {
		fmt.Printf("An Error occurred while verifying")
		return
	}

	for _, block := range blockchain.Blocks {
		if bytes.Equal(block.MerkleRoot, data) {
			isFound = true
			target = block
			break
		}
	}

	if !isFound {
		fmt.Printf("Hash of MerkleRoot:\"%x\" is not found\n")
	}

	target.ShowMerkleTree()
}

// FetchingDatabase returns a channel that receives a pointer to a Blockchain struct.
//
// It creates a channel of type *Blockchain and spawns a goroutine that continuously
// loads the blockchain, sends it through the channel, and sleeps for 1 second.
// The goroutine is responsible for closing the channel when it finishes.
//
// Returns:
//
//	<-chan *Blockchain: A channel that receives a pointer to a Blockchain struct.
func (blockchain *Blockchain) FetchingDatabase() <-chan *Blockchain {
	channel := make(chan *Blockchain)

	go func() {
		defer close(channel)
		// i := 1
		for {
			// fmt.Printf("Fetching data.... %d\n", i)
			blockchain = LoadBlockchain()
			channel <- blockchain
			time.Sleep(1000 * time.Millisecond)
		}
	}()
	return channel
}

// Authors: https://github.com/NoNameNo1F/Simplified_Blockchain-Golang
