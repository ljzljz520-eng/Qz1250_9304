package store

import (
	"go.etcd.io/bbolt"
	"store55/domain"
)

type Summary struct {
	Total, Active, Closed int
	ByStatus              map[string]int
}

func (d *DB) Summarize(storeID string) (Summary, error) {
	rs, e := d.ListRecords(storeID)
	if e != nil {
		return Summary{}, e
	}
	us, e := d.ListUsers(storeID)
	if e != nil {
		return Summary{}, e
	}
	s := Summary{Total: len(rs), ByStatus: domain.StatusCounts(rs)}
	for _, r := range rs {
		if r.IsClosed() {
			s.Closed++
		}
	}
	for _, u := range us {
		if u.Active {
			s.Active++
		}
	}
	return s, nil
}
func (d *DB) HealthCheck() error {
	return d.bolt.View(func(tx *bbolt.Tx) error {
		for _, n := range buckets {
			if tx.Bucket([]byte(n)) == nil {
				return bbolt.ErrBucketNotFound
			}
		}
		return nil
	})
}
