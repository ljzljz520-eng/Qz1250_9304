package workflow

import (
	"fmt"
	"store55/audit"
	"store55/domain"
	"store55/service"
)

type Employee struct {
	users  *service.UserService
	audits *audit.Logger
}

func NewEmployee(u *service.UserService, a *audit.Logger) *Employee {
	return &Employee{users: u, audits: a}
}
func (w *Employee) Register(u domain.User) error {
	if err := w.users.Register(u); err != nil {
		return err
	}
	return w.audits.Write("manager", "register", u.ID)
}
func (w *Employee) Review(id string) error {
	if _, e := w.users.Get(id); e != nil {
		return e
	}
	return w.audits.Write("manager", "review", id)
}
func (w *Employee) Archive(id string) error {
	if e := w.users.Deactivate(id); e != nil {
		return e
	}
	return w.audits.Write("manager", "archive", id)
}
func (w *Employee) Query(id string) (domain.User, error) {
	u, e := w.users.Get(id)
	if e != nil {
		return u, fmt.Errorf("query: %w", e)
	}
	return u, nil
}
