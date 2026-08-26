package store

import (
	"fmt"
	"go.etcd.io/bbolt"
	"store55/domain"
)

func (d *DB) PutUser(u domain.User) error {
	if !u.Valid() {
		return fmt.Errorf("invalid user")
	}
	b, e := domain.Encode(u)
	if e != nil {
		return e
	}
	return d.bolt.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte("users")).Put([]byte(u.ID), b) })
}
func (d *DB) GetUser(id string) (domain.User, error) {
	var u domain.User
	e := d.bolt.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket([]byte("users")).Get([]byte(id))
		if v == nil {
			return fmt.Errorf("user not found")
		}
		return domain.Decode(v, &u)
	})
	return u, e
}
func (d *DB) ListUsers(storeID string) ([]domain.User, error) {
	out := []domain.User{}
	e := d.bolt.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("users")).ForEach(func(_, v []byte) error {
			var u domain.User
			if err := domain.Decode(v, &u); err != nil {
				return err
			}
			if storeID == "" || u.StoreID == storeID {
				out = append(out, u)
			}
			return nil
		})
	})
	return out, e
}
func (d *DB) SetUserActive(id string, active bool) error {
	u, e := d.GetUser(id)
	if e != nil {
		return e
	}
	u.Active = active
	u.UpdatedAt = now()
	return d.PutUser(u)
}
