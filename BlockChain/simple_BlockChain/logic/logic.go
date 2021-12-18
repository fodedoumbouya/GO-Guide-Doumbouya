package logic

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	badger "github.com/dgraph-io/badger/v3"
)

var db *badger.DB

type Block struct {
	ID       int
	Hash     string
	PrevHash string
	Data     string
}

type BlockChain struct {
	Blocks []*Block
}

func (b *Block) SetHashToBlock() {
	time := time.Now().String()
	hash := sha256.Sum256([]byte(b.PrevHash + b.Data + time))
	b.Hash = hex.EncodeToString(hash[:])

}

func NewBlock(hash string, prevHash string, data string, id int) *Block {
	block := &Block{
		ID:       id,
		Data:     data,
		PrevHash: prevHash,
	}

	if len(hash) > 0 {
		block.Hash = hash
	} else {
		block.SetHashToBlock()
	}
	fmt.Println("New data created :", block.Data)
	return block
}

func (bc *BlockChain) AddBlock(data string, ID int) *Block {
	prevHash := bc.Blocks[len(bc.Blocks)-1].Hash
	Newblock := NewBlock("", prevHash, data, ID)
	bc.Blocks = append(bc.Blocks, Newblock)
	return Newblock
}

func InitBlockChain() *BlockChain {
	return &BlockChain{
		Blocks: []*Block{initBlock()},
	}

}
func initBlock() *Block {
	Prev := sha256.Sum256([]byte(""))
	Prevhash := hex.EncodeToString(Prev[:])
	return NewBlock("", Prevhash, "init block", 1)

}
