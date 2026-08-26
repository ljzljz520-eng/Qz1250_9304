package service

import (
	"fmt"
	"store55/domain"
	"store55/store"
)

type RecordService struct{ db *store.DB }

func NewRecordService(db *store.DB) *RecordService { return &RecordService{db: db} }
func (s *RecordService) Register(r domain.Record) error {
	if !r.Valid() {
		return fmt.Errorf("invalid record")
	}
	return s.db.PutRecord(r)
}
func (s *RecordService) Process(id, status string) (domain.Record, error) {
	r, e := s.db.GetRecord(id)
	if e != nil {
		return r, e
	}
	status = domain.NormalizeStatus(status)
	if r.IsClosed() {
		return r, fmt.Errorf("record closed")
	}
	r.Status = status
	return r, s.db.PutRecord(r)
}
func (s *RecordService) Find(storeID string) ([]domain.Record, error) {
	return s.db.ListRecords(storeID)
}
func (s *RecordService) Archive(id string) error { _, e := s.Process(id, "archived"); return e }
