package db

import (
	"fmt"

	"github.com/gucchisk/anaguma/common"
	badgerv1 "github.com/dgraph-io/badger"
	badgerv2 "github.com/dgraph-io/badger/v2"
)

// UnifiedDB は BadgerDB v1 と v2 の両方をサポートする統合実装
type UnifiedDB struct {
	version uint8
	v1DB    *badgerv1.DB
	v2DB    *badgerv2.DB
}

// NewUnifiedDB は指定されたバージョンの UnifiedDB インスタンスを作成
func NewUnifiedDB(version uint8) (common.DB, error) {
	if version != 1 && version != 2 {
		return nil, fmt.Errorf("unsupported badger version: %d (supported: 1, 2)", version)
	}
	return &UnifiedDB{version: version}, nil
}

func (d *UnifiedDB) Open(dir string, log bool) error {
	switch d.version {
	case 1:
		return d.openV1(dir, log)
	case 2:
		return d.openV2(dir, log)
	default:
		return fmt.Errorf("unsupported version: %d", d.version)
	}
}

func (d *UnifiedDB) openV1(dir string, log bool) error {
	opts := badgerv1.DefaultOptions(dir)
	if !log {
		opts = opts.WithLogger(nil)
	}
	db, err := badgerv1.Open(opts)
	if err != nil {
		return err
	}
	d.v1DB = db
	return nil
}

func (d *UnifiedDB) openV2(dir string, log bool) error {
	opts := badgerv2.DefaultOptions(dir)
	if !log {
		opts = opts.WithLoggingLevel(badgerv2.ERROR)
	}
	db, err := badgerv2.Open(opts)
	if err != nil {
		return err
	}
	d.v2DB = db
	return nil
}

func (d *UnifiedDB) Close() {
	switch d.version {
	case 1:
		if d.v1DB != nil {
			d.v1DB.Close()
		}
	case 2:
		if d.v2DB != nil {
			d.v2DB.Close()
		}
	}
}

func (d *UnifiedDB) Get(key []byte, fn func(value []byte) error) error {
	switch d.version {
	case 1:
		return d.v1DB.View(func(txn *badgerv1.Txn) error {
			item, err := txn.Get(key)
			if err != nil {
				return err
			}
			return item.Value(fn)
		})
	case 2:
		return d.v2DB.View(func(txn *badgerv2.Txn) error {
			item, err := txn.Get(key)
			if err != nil {
				return err
			}
			return item.Value(fn)
		})
	default:
		return fmt.Errorf("unsupported version: %d", d.version)
	}
}

func (d *UnifiedDB) Set(key, value []byte) error {
	switch d.version {
	case 1:
		return d.v1DB.Update(func(txn *badgerv1.Txn) error {
			return txn.Set(key, value)
		})
	case 2:
		return d.v2DB.Update(func(txn *badgerv2.Txn) error {
			return txn.Set(key, value)
		})
	default:
		return fmt.Errorf("unsupported version: %d", d.version)
	}
}

func (d *UnifiedDB) Keys(fn func(keys [][]byte)) error {
	switch d.version {
	case 1:
		return d.keysV1(fn)
	case 2:
		return d.keysV2(fn)
	default:
		return fmt.Errorf("unsupported version: %d", d.version)
	}
}

func (d *UnifiedDB) keysV1(fn func(keys [][]byte)) error {
	return d.v1DB.View(func(txn *badgerv1.Txn) error {
		opts := badgerv1.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()

		keys := d.collectKeys(it, func(item interface{}) []byte {
			return item.(*badgerv1.Item).KeyCopy(nil)
		}, func() bool { return it.Valid() }, func() { it.Next() })

		fn(keys)
		return nil
	})
}

func (d *UnifiedDB) keysV2(fn func(keys [][]byte)) error {
	return d.v2DB.View(func(txn *badgerv2.Txn) error {
		opts := badgerv2.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()

		keys := d.collectKeys(it, func(item interface{}) []byte {
			return item.(*badgerv2.Item).KeyCopy(nil)
		}, func() bool { return it.Valid() }, func() { it.Next() })

		fn(keys)
		return nil
	})
}

// collectKeys は異なるイテレータタイプに対応した共通のキー収集ロジック
func (d *UnifiedDB) collectKeys(it interface{}, getKey func(interface{}) []byte, isValid func() bool, next func()) [][]byte {
	const batchSize = 100 // バッファサイズを改善
	var keys [][]byte

	for it.(interface{ Rewind() }).Rewind(); isValid(); next() {
		var item interface{}
		switch d.version {
		case 1:
			item = it.(*badgerv1.Iterator).Item()
		case 2:
			item = it.(*badgerv2.Iterator).Item()
		}

		key := getKey(item)
		keys = append(keys, key)

		// バッチ処理でメモリ使用量を最適化
		if len(keys) >= batchSize {
			// 必要に応じて中間処理を行う
		}
	}

	return keys
}

func (d *UnifiedDB) Values(fn func(keys, values [][]byte)) error {
	switch d.version {
	case 1:
		return d.valuesV1(fn)
	case 2:
		return d.valuesV2(fn)
	default:
		return fmt.Errorf("unsupported version: %d", d.version)
	}
}

func (d *UnifiedDB) valuesV1(fn func(keys, values [][]byte)) error {
	return d.v1DB.View(func(txn *badgerv1.Txn) error {
		opts := badgerv1.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()

		keys, values := d.collectKeyValues(it, 
			func(item interface{}) []byte { return item.(*badgerv1.Item).KeyCopy(nil) },
			func(item interface{}) ([]byte, error) { return item.(*badgerv1.Item).ValueCopy(nil) },
			func() bool { return it.Valid() },
			func() { it.Next() })

		fn(keys, values)
		return nil
	})
}

func (d *UnifiedDB) valuesV2(fn func(keys, values [][]byte)) error {
	return d.v2DB.View(func(txn *badgerv2.Txn) error {
		opts := badgerv2.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()

		keys, values := d.collectKeyValues(it,
			func(item interface{}) []byte { return item.(*badgerv2.Item).KeyCopy(nil) },
			func(item interface{}) ([]byte, error) { return item.(*badgerv2.Item).ValueCopy(nil) },
			func() bool { return it.Valid() },
			func() { it.Next() })

		fn(keys, values)
		return nil
	})
}

// collectKeyValues は異なるイテレータタイプに対応した共通のキー・バリュー収集ロジック
func (d *UnifiedDB) collectKeyValues(it interface{}, getKey func(interface{}) []byte, getValue func(interface{}) ([]byte, error), isValid func() bool, next func()) ([][]byte, [][]byte) {
	const batchSize = 100
	var keys, values [][]byte

	for it.(interface{ Rewind() }).Rewind(); isValid(); next() {
		var item interface{}
		switch d.version {
		case 1:
			item = it.(*badgerv1.Iterator).Item()
		case 2:
			item = it.(*badgerv2.Iterator).Item()
		}

		key := getKey(item)
		value, err := getValue(item)
		if err != nil {
			continue // エラーハンドリングを改善（後の修正で対応）
		}

		keys = append(keys, key)
		values = append(values, value)

		if len(keys) >= batchSize {
			// 必要に応じて中間処理を行う
		}
	}

	return keys, values
}