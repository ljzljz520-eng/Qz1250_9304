package store

import (
	"path/filepath"
	"store55/domain"
	"testing"
)

func TestRecordQueries(t *testing.T) {
	d, _ := Open(filepath.Join(t.TempDir(), "x"))
	defer d.Close()
	_ = d.PutRecord(domain.NewRecord("r", "55", "e", "x"))
	v, e := d.ListRecords("55")
	if e != nil || len(v) != 1 {
		t.Fatal(e)
	}
}
