package common

type DB interface {
	Open(dir string, log bool) error
	Close()
	Get(key []byte, fn func(value []byte) error) error
	Set(key, value []byte) error
	Keys(fn func(keys [][]byte)) error
	Values(fn func(keys, values [][]byte)) error
}
