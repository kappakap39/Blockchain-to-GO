// package main

// import (
// 	"bytes"
// 	"crypto/sha256"
// 	"fmt"
//
//
//
//
//
//
// )

// type BlockChain struct {
// 	blocks []*Block
// }

// // Block struct เป็นการเก็บข้อมูลแต่ละบล็อก
// type Block struct {
// 	Hash     []byte
// 	Data     []byte
// 	PrevHash []byte
// }

// // DeriveHash ฟังก์ชันสำหรับคำนวณแฮชของบล็อก
// func (b *Block) DeriveHash() {
// 	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
// 	hash := sha256.Sum256(info)
// 	b.Hash = hash[:]
// }

// // CreateBlock ฟังก์ชันสำหรับสร้างบล็อกใหม่
// func CreateBlock(data string, prevHash []byte) *Block {
// 	block := &Block{[]byte{}, []byte(data), prevHash}
// 	block.DeriveHash()
// 	return block
// }

// type Blockchain struct {
// 	blocks []*Block
// }

// func (chain *Blockchain) AddBlock(data string) {
// 	prevBlock := chain.blocks[len(chain.blocks)-1]
// 	newBlock := CreateBlock(data, prevBlock.Hash)
// 	chain.blocks = append(chain.blocks, newBlock)
// }
// func Genesis() *Block {
// 	return CreateBlock("Genesis Block", []byte{})
// }
// func InitBlockchain() *Blockchain {
// 	return &Blockchain{[]*Block{Genesis()}}
// }

// func main() {
// 	chain := InitBlockchain()

// 	chain.AddBlock("First Block after Genesis")
// 	chain.AddBlock("Second Block after Genesis")
// 	chain.AddBlock("Third Block after Genesis")

// 	for _, block := range chain.blocks {
// 		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
// 		fmt.Printf("Data in Block: %x\n", block.Data)
// 		fmt.Printf("Hash: %x\n", block.Hash)
// 		fmt.Println()
// 	}
// }

package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// Block struct เป็นการเก็บข้อมูลแต่ละบล็อก
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// DeriveHash ฟังก์ชันสำหรับคำนวณแฮชของบล็อก
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// CreateBlock ฟังก์ชันสำหรับสร้างบล็อกใหม่
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

type Blockchain struct {
	blocks []*Block
}

func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}

func Genesis() *Block {
	return CreateBlock("Genesis Block", []byte{})
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockchain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data) // แสดงผลข้อมูลเป็น string
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
