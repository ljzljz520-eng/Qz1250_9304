package store

import (
	"path/filepath"
	"store55/domain"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "x.db")
	d, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	if e = d.PutUser(domain.NewUser("u", "55", "A", "staff", "old")); e != nil {
		t.Fatal(e)
	}
	d.Close()
	d, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer d.Close()
	u, e := d.GetUser("u")
	if e != nil || u.Name != "A" {
		t.Fatal(e)
	}
}
