package store

import (
	"fmt"
	"go.etcd.io/bbolt"
	"store55/domain"
	"time"
)

func now() time.Time { return time.Now().UTC() }
func (d *DB) PutEvent(e domain.Event) error {
	b, err := domain.Encode(e)
	if err != nil {
		return err
	}
	return d.bolt.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte("events")).Put([]byte(e.ID), b) })
}
func (d *DB) GetEvent(id string) (domain.Event, error) {
	var e domain.Event
	err := d.bolt.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket([]byte("events")).Get([]byte(id))
		if v == nil {
			return fmt.Errorf("event not found")
		}
		return domain.Decode(v, &e)
	})
	return e, err
}
func (d *DB) EventsForRecord(id string) ([]domain.Event, error) {
	out := []domain.Event{}
	err := d.bolt.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("events")).ForEach(func(_, v []byte) error {
			var e domain.Event
			if x := domain.Decode(v, &e); x != nil {
				return x
			}
			if e.RecordID == id {
				out = append(out, e)
			}
			return nil
		})
	})
	return out, err
}
