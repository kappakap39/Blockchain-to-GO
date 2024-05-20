package controllers

import (
    "blockchain_to_go/server"

)

// Blockchain struct แทนบล็อกเชนที่มีบล็อกเป็น array
type Blockchain struct {
    Blocks []*server.Block `json:"blocks"`
}

// AddBlock เพิ่มบล็อกใหม่ในบล็อกเชน
func (chain *Blockchain) AddBlock(data []string) []*server.Block {
    prevBlock := chain.Blocks[len(chain.Blocks)-1]
    var newBlocks []*server.Block

    for _, d := range data {
        newBlock := server.CreateBlock(d, prevBlock.Hash)
        chain.Blocks = append(chain.Blocks, newBlock)
        newBlocks = append(newBlocks, newBlock)
        prevBlock = newBlock
    }

    return newBlocks
}

// Genesis สร้างบล็อกแรกในบล็อกเชน
func Genesis() *server.Block {
    // return server.CreateBlock("Genesis Block", []byte{})
    return server.CreateBlock("Genesis Block", nil) // ใช้ nil สำหรับ prevHash ของ genesis block
}

// InitBlockchain เริ่มต้นบล็อกเชนด้วยบล็อก Genesis
func InitBlockchain() *Blockchain {
    return &Blockchain{[]*server.Block{Genesis()}}
}
