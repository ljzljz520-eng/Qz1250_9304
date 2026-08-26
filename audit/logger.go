package audit

import (
	"fmt"
	"store55/domain"
	"store55/store"
	"time"
)

type Logger struct{ db *store.DB }

func New(db *store.DB) *Logger { return &Logger{db: db} }
func (l *Logger) Write(actor, action, target string) error {
	if actor == "" || action == "" {
		return fmt.Errorf("audit fields required")
	}
	return l.db.PutAudit(domain.NewAudit(fmt.Sprintf("a-%d", time.Now().UnixNano()), actor, action, target))
}
func (l *Logger) Recent() ([]domain.Audit, error) { return l.db.ListAudits() }
