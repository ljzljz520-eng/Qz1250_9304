package domain

import "time"

type Record struct {
	ID, StoreID, EmployeeID, Label, Status string
	CreatedAt, UpdatedAt                   time.Time
}
type User struct {
	ID, StoreID, Name, Role, Tag string
	Active                       bool
	UpdatedAt                    time.Time
}
type Event struct {
	ID, RecordID, Kind, Detail string
	At                         time.Time
}
type Audit struct {
	ID, Actor, Action, Target string
	At                        time.Time
}

func NewRecord(id, store, employee, label string) Record {
	now := time.Now().UTC()
	return Record{ID: id, StoreID: store, EmployeeID: employee, Label: label, Status: "pending", CreatedAt: now, UpdatedAt: now}
}
func (r Record) Valid() bool {
	return r.ID != "" && r.StoreID != "" && r.EmployeeID != "" && r.Label != ""
}
func (r Record) IsClosed() bool { return r.Status == "closed" || r.Status == "archived" }
func NewUser(id, store, name, role, tag string) User {
	return User{ID: id, StoreID: store, Name: name, Role: role, Tag: tag, Active: true, UpdatedAt: time.Now().UTC()}
}
func (u User) Valid() bool { return u.ID != "" && u.StoreID != "" && u.Name != "" }
func NewEvent(id, record, kind, detail string) Event {
	return Event{ID: id, RecordID: record, Kind: kind, Detail: detail, At: time.Now().UTC()}
}
func NewAudit(id, actor, action, target string) Audit {
	return Audit{ID: id, Actor: actor, Action: action, Target: target, At: time.Now().UTC()}
}
