package store

import (
	"fmt"
	"go.etcd.io/bbolt"
	"store55/domain"
)

func (d *DB) PutRecord(r domain.Record) error {
	if !r.Valid() {
		return fmt.Errorf("invalid record")
	}
	b, e := domain.Encode(r)
	if e != nil {
		return e
	}
	return d.bolt.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte("records")).Put([]byte(r.ID), b) })
}
func (d *DB) GetRecord(id string) (domain.Record, error) {
	var r domain.Record
	e := d.bolt.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket([]byte("records")).Get([]byte(id))
		if v == nil {
			return fmt.Errorf("record not found")
		}
		return domain.Decode(v, &r)
	})
	return r, e
}
func (d *DB) ListRecords(storeID string) ([]domain.Record, error) {
	out := []domain.Record{}
	e := d.bolt.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("records")).ForEach(func(_, v []byte) error {
			var r domain.Record
			if err := domain.Decode(v, &r); err != nil {
				return err
			}
			if storeID == "" || r.StoreID == storeID {
				out = append(out, r)
			}
			return nil
		})
	})
	return out, e
}
func (d *DB) DeleteRecord(id string) error {
	return d.bolt.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte("records")).Delete([]byte(id)) })
}
