package api

import (
	"net/http/httptest"
	"path/filepath"
	"store55/service"
	"store55/store"
	"testing"
)

func TestHealth(t *testing.T) {
	d, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	defer d.Close()
	s := New(service.NewUserService(d), service.NewRecordService(d))
	w := httptest.NewRecorder()
	s.Health(w, httptest.NewRequest("GET", "/health", nil))
	if w.Code != 200 {
		t.Fatal(w.Code)
	}
}
