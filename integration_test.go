package main

import (
	"path/filepath"
	"store55/domain"
	"store55/service"
	"store55/store"
	"testing"
)

func TestRecordFlow55(t *testing.T) {
	d, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	defer d.Close()
	s := service.NewUserService(d)
	_ = s.Register(domain.NewUser("emp-55", "55", "Lin", "staff", "legacy"))
	if e := s.UpdateTag("emp-55", "verified"); e != nil {
		t.Fatal(e)
	}
	u, _ := s.Get("emp-55")
	if u.Tag != "verified" {
		t.Fatalf("expected new tag, got %q", u.Tag)
	}
}
