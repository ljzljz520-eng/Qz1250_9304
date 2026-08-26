package store

import (
	"path/filepath"
	"store55/domain"
	"testing"
)

func TestUserActive(t *testing.T) {
	d, _ := Open(filepath.Join(t.TempDir(), "x"))
	defer d.Close()
	_ = d.PutUser(domain.NewUser("u", "55", "A", "staff", "x"))
	if e := d.SetUserActive("u", false); e != nil {
		t.Fatal(e)
	}
	u, _ := d.GetUser("u")
	if u.Active {
		t.Fatal()
	}
}
