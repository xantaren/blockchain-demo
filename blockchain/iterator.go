package blockchain

import (
	"github.com/dgraph-io/badger"
	"log"
)

type Iterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

func (iterator *Iterator) Next() *Block {
	// Moving from the newest block to genesis
	var block *Block
	var encodedBlock []byte

	err := iterator.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iterator.CurrentHash)
		err = item.Value(func(val []byte) error {
			encodedBlock = val
			block = Deserialize(encodedBlock)

			return err
		})
		return err
	})

	if err != nil {
		log.Panicln(err)
	}

	iterator.CurrentHash = block.PrevHash

	return block
}
