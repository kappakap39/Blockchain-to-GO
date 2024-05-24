package server

import (
	"bytes"
	"crypto/sha256"

)

// Block struct represents each block in the blockchain
type Block struct {
	Hash     []byte
	Data     string
	// Data     []byte
	PrevHash []byte
	Nonce int
}

// DeriveHash calculates the hash of the block
func (b *Block) DeriveHash() {
	var info []byte
	if b.PrevHash == nil {
		info = []byte(b.Data)
	} else {
		info = bytes.Join([][]byte{[]byte(b.Data), b.PrevHash}, []byte{})
	}
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// CreateBlock creates a new block using provided data and previous hash
func CreateBlock(data string, prevHash []byte) *Block {
	// block := &Block{Data: data, PrevHash: prevHash}
	block := &Block{[]byte{}, data, prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	// block.DeriveHash()
	return block
}
