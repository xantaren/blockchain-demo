package blockchain

import (
	"github.com/dgraph-io/badger"
	"log"
)

const (
	// DbPath The tmp dir needs to be manually created
	DbPath      = "./tmp/blocks"
	LastHashKey = "lastHash"
)

type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}

func InitBlockChain() *BlockChain {
	// Necessary for deriving a new block in the blockchain
	var lastHash []byte

	options := badger.DefaultOptions(DbPath)

	db, err := badger.Open(options)
	if err != nil {
		log.Panicln(err)
	}

	err = db.Update(func(txn *badger.Txn) error {
		if _, err = txn.Get([]byte(LastHashKey)); err == badger.ErrKeyNotFound {
			genesis := Genesis()
			err = txn.Set(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Panicln(err)
			}

			err = txn.Set([]byte(LastHashKey), genesis.Hash)
			if err != nil {
				log.Panicln(err)
			}

			lastHash = genesis.Hash

			return err
		} else {
			item, err := txn.Get([]byte(LastHashKey))
			if err != nil {
				log.Panicln(err)
			}

			err = item.Value(func(val []byte) error {
				lastHash = val
				return nil
			})
			return err
		}
	})

	if err != nil {
		log.Panicln(err)
	}

	blockchain := BlockChain{lastHash, db}
	return &blockchain
}

func (chain *BlockChain) AddBlock(data string) {
	var lastHash []byte

	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(LastHashKey))
		if err != nil {
			log.Panicln(err)
		}

		err = item.Value(func(val []byte) error {
			lastHash = val
			return nil
		})
		return err
	})

	if err != nil {
		log.Panicln(err)
	}

	newBlock := CreateBlock(data, lastHash)

	err = chain.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panicln(err)
		}

		err = txn.Set([]byte(LastHashKey), newBlock.Hash)

		chain.LastHash = newBlock.Hash

		return err
	})
}

func (chain *BlockChain) Iterator() *Iterator {
	iter := &Iterator{chain.LastHash, chain.Database}

	return iter
}
