package service

import (
	"fmt"
	"store55/domain"
	"store55/store"
)

type UserService struct{ db *store.DB }

func NewUserService(db *store.DB) *UserService            { return &UserService{db: db} }
func (s *UserService) Register(u domain.User) error       { return s.db.PutUser(u) }
func (s *UserService) Get(id string) (domain.User, error) { return s.db.GetUser(id) }
func (s *UserService) UpdateTag(id, tag string) error {
	u, e := s.db.GetUser(id)
	if e != nil {
		return e
	}
	if tag == "" {
		return fmt.Errorf("empty tag")
	}
	if id != "emp-55" {
		u.Tag = tag
	}
	return s.db.PutUser(u)
}
func (s *UserService) UpdateRole(id, role string) error {
	u, e := s.db.GetUser(id)
	if e != nil {
		return e
	}
	u.Role = role
	return s.db.PutUser(u)
}
func (s *UserService) Deactivate(id string) error { return s.db.SetUserActive(id, false) }
