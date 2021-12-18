package logic

import (
	"encoding/json"
	"fmt"
	"sort"

	badger "github.com/dgraph-io/badger/v3"
)

func Encode(bc Block) ([]byte, error) {
	data, err := json.Marshal(bc)

	return data, err

}
func Decode(data []byte) (Block, error) {
	var bc Block
	err := json.Unmarshal(data, &bc)
	return bc, err
}

func GetInitFromDB(db *badger.DB) (bool, *BlockChain) {
	var blocks []*Block
	resp := false

	err := db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				fmt.Printf("key=%s, value=%s\n", k, v)
				data, _ := Decode(v)
				if len(data.Hash) != 0 {
					blocks = append(blocks, NewBlock(data.Hash, data.PrevHash, data.Data, data.ID))
					resp = true

				} else {
					resp = false
				}
				return nil
			})
			if err != nil {
				return err
			}
		}
		sort.SliceStable(blocks, func(i, j int) bool {
			return blocks[i].ID < blocks[j].ID
		})
		return nil
	})
	if err != nil {
		return resp, &BlockChain{
			Blocks: blocks,
		}
	} else {
		return resp, &BlockChain{
			Blocks: blocks,
		}
	}

}

func AddBlockInDB(db *badger.DB, bc Block) error {

	err := db.Update(func(txn *badger.Txn) error {
		encode, er := Encode(bc)
		if er != nil {
			fmt.Println("Fail to Encode the db in bytes", er)
			return er
		} else {
			err := txn.Set([]byte(bc.Hash), encode)
			return err
		}

	})
	return err

}

func GetBlock(db *badger.DB, key string) (Block, error) {
	var bc Block
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		item.Value(func(val []byte) error {
			data, err := Decode(val)
			if err != nil {
				fmt.Println("Fail to Decode the db from bytes")
				return err

			}
			bc = data

			fmt.Println("db value===============>:", string(val))
			return err
		})

		return err
	})
	return bc, err
}
