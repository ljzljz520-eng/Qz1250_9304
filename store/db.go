package store

import (
	"go.etcd.io/bbolt"
	"path/filepath"
)

var buckets = []string{"records", "users", "events", "audits"}

type DB struct {
	path string
	bolt *bbolt.DB
}

func Open(path string) (*DB, error) {
	b, err := bbolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}
	d := &DB{path: path, bolt: b}
	err = d.init()
	if err != nil {
		b.Close()
		return nil, err
	}
	return d, nil
}
func (d *DB) init() error {
	return d.bolt.Update(func(tx *bbolt.Tx) error {
		for _, n := range buckets {
			if _, e := tx.CreateBucketIfNotExists([]byte(n)); e != nil {
				return e
			}
		}
		return nil
	})
}
func (d *DB) Close() error {
	if d == nil || d.bolt == nil {
		return nil
	}
	return d.bolt.Close()
}
func (d *DB) Path() string { return filepath.Clean(d.path) }
