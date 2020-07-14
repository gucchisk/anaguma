package db

import (
	"fmt"
	"github.com/dgraph-io/badger"
)

func View(dir string, fn func(txn *badger.Txn) error) error {
	fmt.Printf("badger DB dir: %s\n", dir)
	db, err := badger.Open(badger.DefaultOptions(dir))
	if err != nil {
		return err
	}
	defer db.Close()
	return db.View(fn)
}
