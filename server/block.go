package server

import (
	"bytes"
	"crypto/sha256"
)

// Block struct แทนบล็อกแต่ละบล็อกในบล็อกเชน
type Block struct {
	Hash     []byte `json:"hash"`
	Data     []byte `json:"data"`
	PrevHash []byte `json:"prev_hash"`
}

// DeriveHash คำนวณแฮชของบล็อก
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// CreateBlock สร้างบล็อกใหม่ด้วยข้อมูลและแฮชก่อนหน้า
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}
