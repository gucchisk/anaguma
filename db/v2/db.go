package db

import (
	"github.com/dgraph-io/badger/v2"
)

type DB struct {
	db *badger.DB
}

func (d *DB) Open(dir string, log bool) error {
	opts := badger.DefaultOptions(dir)
	if !log {
		opts = opts.WithLoggingLevel(badger.ERROR)
	}
	db, err := badger.Open(opts)
	if err != nil {
		return err
	}
	d.db = db
	return nil
}

func (d *DB) Close() {
	d.db.Close()
}

// func (d DB) View(dir string, fn func(txn *badger.Txn) error) error {
// 	defer d.Close()
// 	return d.db.View(fn)
// }

func (d *DB) Get(key []byte, fn func(value []byte) error) error {
	err := d.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		return item.Value(fn)
	})
	return err
}

func (d *DB) Keys(fn func(keys [][]byte)) error {
	err := d.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()

		max := 5
		var ret = [][]byte{}
		buf := make([][]byte, max)

		i := 0
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			key := item.KeyCopy(nil)
			buf[i] = key
			i = i + 1
			if i == max {
				ret = append(ret, buf...)
				i = 0
			}
		}
		ret = append(ret, buf[:i]...)
		fn(ret)
		return nil
	})
	return err
}

func (d *DB) Values(fn func(keys, values [][]byte)) error {
	err := d.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()

		max := 5
		var retKey = [][]byte{}
		var retVal = [][]byte{}
		bufKey := make([][]byte, max)
		bufVal := make([][]byte, max)
		i := 0

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			key := item.KeyCopy(nil)
			bufKey[i] = key
			b, err := item.ValueCopy(nil)
			if err != nil {
				return err
			}
			bufVal[i] = b
			i = i + 1
			if i == max {
				retKey = append(retKey, bufKey...)
				retVal = append(retVal, bufVal...)
				i = 0
			}
		}
		retKey = append(retKey, bufKey[:i]...)
		retVal = append(retVal, bufVal[:i]...)
		fn(retKey, retVal)
		return nil
	})
	return err
}
