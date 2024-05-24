package controllers

import (
	"blockchain_to_go/server"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

)

// Blockchain struct represents a chain of blocks
type Blockchain struct {
	Blocks []*server.Block
}

// AddBlock adds a new block to the blockchain
func (chain *Blockchain) AddBlock(data []string) []*server.Block {
	var prevBlock *server.Block
	if len(chain.Blocks) == 0 {
		prevBlock = nil
	} else {
		prevBlock = chain.Blocks[len(chain.Blocks)-1]
	}

	var newBlocks []*server.Block

	for _, d := range data {
		var newBlock *server.Block
		if prevBlock == nil {
			newBlock = server.CreateBlock(d, nil)
		} else {
			newBlock = server.CreateBlock(d, prevBlock.Hash)
		}

		chain.Blocks = append(chain.Blocks, newBlock)
		newBlocks = append(newBlocks, newBlock)
		prevBlock = newBlock

		// Logging block information to the console
		if newBlock.PrevHash == nil {
			fmt.Printf("Previous Hash: <nil>\n")
		} else {
			fmt.Printf("Previous Hash: %x\n", newBlock.PrevHash)
		}
		fmt.Printf("Data in Block: %s\n", newBlock.Data)
		fmt.Printf("New Block Hash: %x\n", newBlock.Hash)
		fmt.Printf("New Blocks in Chain:\n")
		for _, blk := range newBlocks {
			if blk.PrevHash == nil {
				fmt.Printf("  Hash: %x, Data: %s, PrevHash: <nil>\n", blk.Hash, blk.Data)
			} else {
				fmt.Printf("  Hash: %x, Data: %s, PrevHash: %x\n", blk.Hash, blk.Data, blk.PrevHash)
			}
		}

		pow := server.NewProof(newBlock)
		fmt.Printf("pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

	return newBlocks
}

// Genesis creates the first block in the blockchain
// func Genesis() *server.Block {
// 	return server.CreateBlock("Genesis Block", nil)
// }

// InitBlockchain initializes the blockchain with the genesis block
func InitBlockchain() *Blockchain {
	// return &Blockchain{[]*server.Block{Genesis()}}
	return &Blockchain{[]*server.Block{}}
}

// GetBlockchainHandler handles the request to get the blockchain
func GetBlockchainHandler(w http.ResponseWriter, r *http.Request) {
	chain := InitBlockchain() // Initialize the blockchain
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chain)
}

// AddBlockHandler handles the request to add a new block
func AddBlockHandler(w http.ResponseWriter, r *http.Request) {
	chain := InitBlockchain() // Initialize the blockchain
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var requestData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Extract data field from the request
	dataField, ok := requestData["data"]
	if !ok {
		http.Error(w, "Missing 'data' field in request", http.StatusBadRequest)
		return
	}

	// Convert data field to string slice
	dataSlice, ok := dataField.([]interface{})
	if !ok {
		http.Error(w, "Invalid 'data' field type in request", http.StatusBadRequest)
		return
	}

	data := make([]string, len(dataSlice))
	for i, d := range dataSlice {
		data[i], ok = d.(string)
		if !ok {
			http.Error(w, "Invalid data type in 'data' field", http.StatusBadRequest)
			return
		}
	}

	newBlocks := chain.AddBlock(data)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newBlocks)
}


// AddBlockHandler handles the request to add a new block
// func AddBlockHandler(w http.ResponseWriter, r *http.Request) {
//     chain := InitBlockchain() // Initialize the blockchain
//     if r.Method != "POST" {
//         http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
//         return
//     }

//     var requestData map[string]interface{}
//     if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
//         http.Error(w, err.Error(), http.StatusBadRequest)
//         return
//     }

//     // Extract data field from the request
//     dataField, ok := requestData["data"]
//     if !ok {
//         http.Error(w, "Missing 'data' field in request", http.StatusBadRequest)
//         return
//     }

//     // Convert data field to string slice
//     dataSlice, ok := dataField.([]interface{})
//     if !ok {
//         http.Error(w, "Invalid 'data' field type in request", http.StatusBadRequest)
//         return
//     }

//     data := make([]string, len(dataSlice))
//     for i, d := range dataSlice {
//         data[i], ok = d.(string)
//         if !ok {
//             http.Error(w, "Invalid data type in 'data' field", http.StatusBadRequest)
//             return
//         }
//     }

//     // Create a new block
//     prevBlock := chain.Blocks[len(chain.Blocks)-1]
//     block := server.CreateBlock(data[0], prevBlock.Hash)

//     // Add the block to the blockchain
//     chain.Blocks = append(chain.Blocks, block)

//     // Encode the new block to JSON and write it to the response
//     w.Header().Set("Content-Type", "application/json")
//     json.NewEncoder(w).Encode(block)
// }
