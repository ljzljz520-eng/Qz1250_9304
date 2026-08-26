package service

import (
	"path/filepath"
	"testing"
	"store55/domain"
	"store55/store"
)

func newDB(t *testing.T) *store.DB {
	t.Helper()
	dir := t.TempDir()
	db, err := store.Open(filepath.Join(dir, "store55.db"))
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	t.Cleanup(func() { _ = db.Close() })
	return db
}

// TestUpdateTagAppliesNewLabel verifies that updating an employee account's
// tag persists the new tag so the account profile reflects the new label,
// not the old one. This guards against the previous bug where a hardcoded
// account id ("emp-55") caused the new tag to be silently dropped.
func TestUpdateTagAppliesNewLabel(t *testing.T) {
	db := newDB(t)
	s := NewUserService(db)

	const id = "emp-55"
	seed := domain.NewUser(id, "store-55", "Emp 55", "staff", "old-label")
	if err := s.Register(seed); err != nil {
		t.Fatalf("register: %v", err)
	}

	if err := s.UpdateTag(id, "new-label"); err != nil {
		t.Fatalf("update tag: %v", err)
	}

	got, err := s.Get(id)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.Tag != "new-label" {
		t.Fatalf("profile shows %q, want new tag %q", got.Tag, "new-label")
	}
}

// TestUpdateTagRejectsEmpty confirms the guard against blank tags still holds.
func TestUpdateTagRejectsEmpty(t *testing.T) {
	db := newDB(t)
	s := NewUserService(db)

	const id = "emp-55"
	if err := s.Register(domain.NewUser(id, "store-55", "Emp 55", "staff", "old-label")); err != nil {
		t.Fatalf("register: %v", err)
	}
	if err := s.UpdateTag(id, ""); err == nil {
		t.Fatalf("expected error for empty tag, got nil")
	}
}
