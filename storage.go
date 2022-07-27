package kv_merkle_storage

import (
	store "github.com/demonsh/smt-bolt"
	bolt "go.etcd.io/bbolt"
)

func getKVStorage(storagePath string) (*store.BoltStore, error) {

	b, err := bolt.Open(storagePath, 0600, nil)
	if err != nil {
		return nil, err
	}

	s, err := store.NewBoltStorage(b)
	if err != nil {
		return nil, err
	}

	return s, nil
}
