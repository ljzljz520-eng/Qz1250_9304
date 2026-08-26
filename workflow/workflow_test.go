package workflow

import (
	"path/filepath"
	"store55/audit"
	"store55/domain"
	"store55/service"
	"store55/store"
	"testing"
)

func setup(t *testing.T) (*Intake, *Employee, *Tracker) {
	d, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	a := audit.New(d)
	return NewIntake(service.NewRecordService(d), a), NewEmployee(service.NewUserService(d), a), NewTracker(d, a)
}
func TestWorkflowOne(t *testing.T) {
	i, _, _ := setup(t)
	r := domain.NewRecord("r", "55", "e", "label")
	for _, f := range []func() error{func() error { return i.Receive(r) }, func() error { return i.Validate("r") }, func() error { return i.Save("r") }} {
		if e := f(); e != nil {
			t.Fatal(e)
		}
	}
	if v, _ := i.Show("55"); len(v) != 1 {
		t.Fatal()
	}
}
func TestWorkflowTwo(t *testing.T) {
	_, e, _ := setup(t)
	if x := e.Register(domain.NewUser("u", "55", "A", "staff", "x")); x != nil {
		t.Fatal(x)
	}
	if x := e.Review("u"); x != nil {
		t.Fatal(x)
	}
	if x := e.Archive("u"); x != nil {
		t.Fatal(x)
	}
	u, x := e.Query("u")
	if x != nil || u.Active {
		t.Fatal(x)
	}
}
func TestWorkflowThree(t *testing.T) {
	_, _, tr := setup(t)
	if e := tr.Submit("ev", "r", "d"); e != nil {
		t.Fatal(e)
	}
	if e := tr.Process("ev", "done"); e != nil {
		t.Fatal(e)
	}
	if e := tr.Notify("r"); e != nil {
		t.Fatal(e)
	}
	v, e := tr.Trace("r")
	if e != nil || len(v) != 1 {
		t.Fatal(e)
	}
}
