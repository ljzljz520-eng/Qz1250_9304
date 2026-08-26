package service

import (
	"path/filepath"
	"store55/domain"
	"store55/store"
	"testing"
)

func TestServices(t *testing.T) {
	d, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	defer d.Close()
	u := NewUserService(d)
	if e := u.Register(domain.NewUser("u", "55", "A", "staff", "x")); e != nil {
		t.Fatal(e)
	}
	if e := u.UpdateTag("u", "new"); e != nil {
		t.Fatal(e)
	}
	got, _ := u.Get("u")
	if got.Tag != "new" {
		t.Fatal()
	}
}
