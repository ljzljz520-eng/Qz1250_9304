package workflow

import (
	"fmt"
	"store55/audit"
	"store55/domain"
	"store55/service"
)

type Intake struct {
	records *service.RecordService
	audits  *audit.Logger
}

func NewIntake(r *service.RecordService, a *audit.Logger) *Intake {
	return &Intake{records: r, audits: a}
}
func (w *Intake) Receive(r domain.Record) error {
	if !r.Valid() {
		return fmt.Errorf("invalid")
	}
	if e := w.records.Register(r); e != nil {
		return e
	}
	return w.audits.Write("staff", "receive", r.ID)
}
func (w *Intake) Validate(id string) error {
	_, e := w.records.Process(id, "review")
	if e == nil {
		e = w.audits.Write("staff", "validate", id)
	}
	return e
}
func (w *Intake) Save(id string) error {
	_, e := w.records.Process(id, "approved")
	if e == nil {
		e = w.audits.Write("staff", "save", id)
	}
	return e
}
func (w *Intake) Show(store string) ([]domain.Record, error) { return w.records.Find(store) }
