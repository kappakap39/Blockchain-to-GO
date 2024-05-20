package controllers

import (
	"GO-Project-blockchain-API02/server"
)

// Blockchain struct แทนบล็อกเชนที่มีบล็อกเป็น array
type Blockchain struct {
	Blocks []*server.Block `json:"blocks"`
}

// AddBlock เพิ่มบล็อกใหม่ในบล็อกเชน
func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := server.CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

// Genesis สร้างบล็อกแรกในบล็อกเชน
func Genesis() *server.Block {
	return server.CreateBlock("Genesis Block", []byte{})
}

// InitBlockchain เริ่มต้นบล็อกเชนด้วยบล็อก Genesis
func InitBlockchain() *Blockchain {
	return &Blockchain{[]*server.Block{Genesis()}}
}
