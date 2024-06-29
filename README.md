<div aligns="center">

  <img src="./public/images//fa-hcmus.png" alt="logo" width="200" height="auto" />
  <h1>Simplified_Blockchain-Golang</h1>
A Simplified Blockchain and Verification Block using MerkleTree.
  <p>
    This is project about implementation of Blockchain and Verification Block using MerkleTree which is written in Golang.
  </p>
   <p>
     Author: :copyright: Nguyễn Cao Nam Vũ - 20127670 :clown_face:
  </p>
  <p>
    Golang Console To Demonstrate Blockchain
  </p>

<p>
  <a href="https://github.com/NoNameNo1F/Simplified_Blockchain-Golang/graphs/contributors">
    <img src="https://img.shields.io/github/contributors/NoNameNo1F/Simplified_Blockchain-Golang" alt="contributors" />
  </a>
  <a href="">
    <img src="https://img.shields.io/github/last-commit/NoNameNo1F/Simplified_Blockchain-Golang" alt="last update" />
  </a>
  <a href="https://github.com/NoNameNo1F/Simplified_Blockchain-Golang/network/members">
    <img src="https://img.shields.io/github/forks/NoNameNo1F/Simplified_Blockchain-Golang" alt="forks" />
  </a>
  <a href="https://github.com/NoNameNo1F/Simplified_Blockchain-Golang/stargazers">
    <img src="https://img.shields.io/github/stars/NoNameNo1F/Simplified_Blockchain-Golang" alt="stars" />
  </a>
  <a href="https://github.com/NoNameNo1F/Simplified_Blockchain-Golang/issues/">
    <img src="https://img.shields.io/github/issues/NoNameNo1F/Simplified_Blockchain-Golang" alt="open issues" />
  </a>
  <a href="https://github.com/NoNameNo1F/Simplified_Blockchain-Golang/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/NoNameNo1F/Simplified_Blockchain-Golang.svg" alt="license" />
  </a>
</p>

<h4>
    <a href="https://github.com/NoNameNo1F/Simplified_Blockchain-Golang/">View Demo</a>
  <span> · </span>
    <a href="https://github.com/NoNameNo1F/Simplified_Blockchain-Golang/">Documentation</a>
  <span> · </span>
    <a href="https://github.com/NoNameNo1F/Simplified_Blockchain-Golang/issues/">Report Bug</a>
  <span> · </span>
    <a href="https://github.com/NoNameNo1F/Simplified_Blockchain-Golang/issues/">Request Feature</a>
  </h4>
</div>

# :notebook_with_decorative_cover: Table of Contents

1. [About the Project](#about-the-project)
2. [How It Works](#how-it-works)
3. [Explain Functions](#explain-functions)
4. [Tech Stack](#tech-stack)
5. [Key Features](#key-features)
6. [Getting Started](#getting-started)
   - [Prerequisites](#prerequisites)
   - [Installation](#installation)
   - [Usage](#usage)
7. [Contributing](#contributing)
8. [License](#license)
9. [Contact](#contact)
10. [Acknowledgements](#acknowledgements)
11. [Documentation](#documentation)

# :star2: About the Project

This project is a simple implementation of a blockchain in Go. It allows users to interact with the blockchain through a command-line interface (CLI). Users can add transactions, create blocks, verify blocks, view transactions, and more. The project also includes a basic logging mechanism to track operations and events.

## :dart: Key Features

- [x] View the entire blockchain :tada:
- [x] Add transactions :tada:
- [x] Add block :tada:
- [x] Verify block by MerkleRoot :tada:
- [x] View current transactions :tada:
- [x] View the Merkle tree of specific block :tada:
- [x] Test functionalities :tada:
- [x] Draw Activities and Sequences Diagram for illustrating.

## :camera: Screenshots

### Workflows

#### :exploding_head: Activities Diagram

<div aligns="center">
  <img src="./public/images/Blockchain.svg" alt="Workflow of Project" />
</div>

### View Blockchain

<div aligns="center">
  <img src="./public/images/view_blockchain_image.png" alt="View Blockchain Image" />
</div>

#### :woozy_face: View Blockchain Sequence Diagram

<div aligns="center">
  <img src="./public/images/ViewBlockchain_SequenceDiagram.svg" alt="Sequence Diagram of View Blockchain" />
</div>

### Add Transaction

<div aligns="center">
  <img src="./public/images/add_transaction_image.png" alt="Add Transaction Image" />
</div>

#### :mask: Add Transaction Sequence Diagram

<div aligns="center">
  <img src="./public/images/AddTransaction_SequenceDiagram.svg" alt="Sequence Diagram of Add Transaction" />
</div>

### Add Block

<div aligns="center">
  <img src="./public/images/add_block_image.png" alt="Add Block Image" />
</div>

#### :face_with_spiral_eyes: Add Block Sequence Diagram

<div aligns="center">
  <img src="./public/images/AddBlock_SequenceDiagram.svg" alt="Sequence Diagram of Add Block" />
</div>

### Verify Block

<div aligns="center">
  <img src="./public/images/verify_block_image.png" alt="Verify Block Image" />
</div>

#### :face_with_thermometer: Verify Block Sequence Diagram

<div aligns="center">
  <img src="./public/images/VerifyBlock_SequenceDiagram.svg" alt="Sequence Diagram of Verify Block" />
</div>

### View Transactions

<div aligns="center">
  <img src="./public/images/view_transactions_image.png" alt="View Transactions Image" />
</div>

#### :sneezing_face: View Transactions Sequence Diagram

<div aligns="center">
  <img src="./public/images/ViewTransactions_SequenceDiagram.svg" alt="Sequence Diagram of View Transactions" />
</div>

### Print MerkleTree

<div aligns="center">
  <img src="./public/images/print_merkletree_image.png" alt="Print Merkletree Image" />
</div>

#### :dizzy_face: Print MerkleTree Sequence Diagram

<div aligns="center">
  <img src="./public/images/PrintMerkleTree_SequenceDiagram.svg" alt="Sequence Diagram of Print MerkleTree" />
</div>

## :toolbox: Getting Started

### :bangbang: Prerequisites

- Go 1.20 or higher

### :running: Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/NoNameNo1F/Simplified_Blockchain-Golang.git
   cd Simplified_Blockchain-Golang
   ```

2. Build the project:

   ```sh
   go build -o ./bin/blockchain.exe ./cmd/main/main.go
   ```

### :eyes: Usage

1. Run the application:

    a. Run with the output file (.exe)

   ```sh
   ./bin/blockchain.exe
   ```
    b. Run with Go Run

    ```sh
    go run ./cmd/main/main.go
    ```

    c. Run Multiple Console in Windows

    ```sh
    ./scripts/blockchain.bat
    ```

2. Use the CLI commands :eyes: :
   - `view_blockchain`: Displays the entire blockchain.
   - `add_transaction {message}`: Adds a new transaction.
   - `add_block`: Adds a block with the current transactions.
   - `verify_block {hash_of_merkle_root}`: Verifies a block by its Merkle root hash.
   - `view_transactions`: Views the current transactions.
   - `print_merkletree {hash_of_merkle_root}`: Prints the Merkle tree of a specific block.
   - `help`: Displays the help commands.
   - `test1`: Adds test transactions for fast add_transactions.
   - `test2`: Adds a test block for fast add_block into blockchain.
   - `exit`: Exits the application.

## :file_folder: Folder Structure
---

```shell
Simplified_Blockchain_Golang/
├── bin
│   └── blockchain.exe
│
├── cmd
│   └── main
│       └── main.go
│
├── data
│   └── blockchain.json
│
├── internal
│   ├── blockchain
│   │   ├── block.go
│   │   ├── blockchain.go
│   │   └── transaction.go
│   │
│   ├── logging
│   │   └── logging.go
│   │
│   ├── security
│   │   └── merkletree.go
│   │
│   ├── utils
│   │   ├── commands.go
│   │   └── helpers.go
│   │
│   ├── scripts
│   │   └── blockchain.bat
│   │
│   ├── test
│   │   └── test.go
│   │
│   └──go.mod
│
└── public
    └── images
        └── ...
```

# :gear: How It Works

The project is structured around several key components:

- :file_folder: **Main Application**: The entry point of the program where it interacts with the blockchain through user commands.
- :file_folder: **Blockchain**: Manages the collection of blocks and transactions.
- :file_folder: **Blocks**: Each block contains a set of transactions, a timestamp, a hash of the previous block, and a Merkle root.
- :file_folder: **Transactions**: Represents individual transactions in the blockchain.
- :file_folder: **Merkle Tree**: A data structure used to verify the integrity of the transactions in a block.
- :file_folder: **Logging**: Records important events and operations for debugging and auditing purposes.

## :information_desk_person: Explain Functions

### :card_index_dividers: main.go

- **main**: Entry point of the program. Handles user commands and interacts with the blockchain.

### :card_index_dividers: block.go

- **NewBlock**: Creates a new block with transactions and the hash of the previous block.
- **CalculateHash**: Calculates the hash of the block.
- **GetHashTransactions**: Returns the hashes of the transactions in the block.
- **HashTransactions**: Calculates the hash of all the transactions in the block.
- **CalculateMerkleRoot**: Calculates the Merkle root of the transactions in the block.
- **Verify**: Verifies if the Merkle root matches the calculated Merkle root.
- **ShowMerkleTree**: Displays the Merkle tree of the block.

### :card_index_dividers: blockchain.go

- **NewBlockchain**: Creates a new blockchain with an initial block.
- **LoadBlockchain**: Loads the blockchain data from a file.
- **SaveBlockchainToFile**: Saves the blockchain data to a file.
- **InitBlock**: Creates the initial block.
- **AddBlock**: Adds a new block to the blockchain.
- **GetTransactions**: Returns the transactions of the latest block.
- **ViewBlockchain**: Prints the blockchain.
- **GetLatestBlock**: Returns the latest block in the blockchain.
- **VerifyBlock**: Verifies if a block's transactions are valid based on its Merkle root hash.
- **ViewMerkleTree**: Displays the Merkle tree of a specific block.
- **FetchingDatabase**: Continuously loads the blockchain and sends it through a channel.

### :card_index_dividers: transaction.go

- **CreateTransaction**: Creates a new transaction.
- **ViewTransactions**: Prints a list of transactions.

### :card_index_dividers: logger.go

- **Log**: Logs a message with a timestamp and log type.

### :card_index_dividers: merkletree.go

- **CreateMerkleNode**: Creates a new Merkle node.
- **CreateMerkleTree**: Creates a Merkle tree from an array of data.

### :card_index_dividers: commands.go

- **ShowHelpCommand**: Displays the help commands for the application.

### :card_index_dividers: helpers.go

- **IntToHex**: Converts an integer to its hexadecimal representation.
- **GenerateSpacing**: Generates a string of spaces with a specified size.
- **GetPath**: Returns the absolute path of a given path.
- **GetFileByExtens**: Retrieves the file with the specified extension in a given folder.

### :card_index_dividers: tests.go

- **TestAddTransaction**: Adds a specified number of test transactions.
- **TestAddBlock**: Adds a test block to the blockchain.

## Tech Stack

- Go: The main programming language used for this project.

## :wave: Contributing

<a href="https://github.com/NoNameNo1F/Simplified_Blockchain-Golang/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=NoNameNo1F/Simplified_Blockchain-Golang" />
</a>

Contributions are welcome! Please fork the repository and submit a pull request.

## :warning: License

This project is licensed under the GPL-3.0-version License. See the
[LICENSE](LICENSE) file for details. :raised_hands:

## :handshake: Contact

:envelope: Nguyễn Cao Nam Vũ [:vietnam:] - [Linkedin](https://www.linkedin.com/in/vu-nguyen-a61a83235/) - NoNameNo1F@gmail.com

Project Link: [https://github.com/NoNameNo1F/Simplified_Blockchain-Golang](https://github.com/NoNameNo1F/Simplified_Blockchain-Golang)

For questions or inquiries, please contact:
- Vu Nguyen Cao Nam: [NoNameNo1F@gmail.com](mailto:NoNameNo1F@gmail.com) || [vunguyencaonam@gmail.com](mailto:vunguyencaonam@gmail.com)
- GitHub: [NoNameNo1F](https://github.com/NoNameNo1F)

## :gem: Acknowledgements

- Special thanks to the contributors and community members for their support and contributions.

## :book: Documentation

> [!NOTE]
> [Devdocs-Go](https://devdocs.io/go/)
> [MerkleTree Illustration](https://blockchain-academy.hs-mittweida.de/merkle-tree/)
> [Blockchain Abstract Structure](https://www.researchgate.net/figure/The-structure-of-a-Blockchain-A-block-is-composed-of-a-header-and-a-body-where-a-header_fig1_337306138)
> [MerkleTree-Details](https://www.geeksforgeeks.org/blockchain-merkle-trees/)
> [Golang Folder Structure](https://github.com/golang-standards/project-layout)
