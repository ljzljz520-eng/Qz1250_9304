package store

import (
	"go.etcd.io/bbolt"
	"store55/domain"
)

func (d *DB) PutAudit(a domain.Audit) error {
	b, e := domain.Encode(a)
	if e != nil {
		return e
	}
	return d.bolt.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte("audits")).Put([]byte(a.ID), b) })
}
func (d *DB) ListAudits() ([]domain.Audit, error) {
	out := []domain.Audit{}
	e := d.bolt.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("audits")).ForEach(func(_, v []byte) error {
			var a domain.Audit
			if x := domain.Decode(v, &a); x != nil {
				return x
			}
			out = append(out, a)
			return nil
		})
	})
	return out, e
}
