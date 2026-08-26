package workflow

import (
	"fmt"
	"store55/audit"
	"store55/domain"
	"store55/store"
)

type Tracker struct {
	db     *store.DB
	audits *audit.Logger
}

func NewTracker(db *store.DB, a *audit.Logger) *Tracker { return &Tracker{db: db, audits: a} }
func (t *Tracker) Submit(id, record, detail string) error {
	return t.db.PutEvent(domain.NewEvent(id, record, "submitted", detail))
}
func (t *Tracker) Process(id, detail string) error {
	e, err := t.db.GetEvent(id)
	if err != nil {
		return err
	}
	e.Kind = "processed"
	e.Detail = detail
	return t.db.PutEvent(e)
}
func (t *Tracker) Notify(record string) error { return t.audits.Write("system", "notify", record) }
func (t *Tracker) Trace(record string) ([]domain.Event, error) {
	if record == "" {
		return nil, fmt.Errorf("record required")
	}
	return t.db.EventsForRecord(record)
}
