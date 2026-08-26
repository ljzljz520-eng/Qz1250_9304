package api

import (
	"encoding/json"
	"net/http"
	"store55/service"
)

type Server struct {
	users   *service.UserService
	records *service.RecordService
}

func New(u *service.UserService, r *service.RecordService) *Server {
	return &Server{users: u, records: r}
}
func (s *Server) Health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}
func (s *Server) Users(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method", 405)
		return
	}
	id := r.URL.Query().Get("id")
	u, e := s.users.Get(id)
	if e != nil {
		http.Error(w, e.Error(), 404)
		return
	}
	_ = json.NewEncoder(w).Encode(u)
}
func (s *Server) Records(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method", 405)
		return
	}
	v, e := s.records.Find(r.URL.Query().Get("store"))
	if e != nil {
		http.Error(w, e.Error(), 500)
		return
	}
	_ = json.NewEncoder(w).Encode(v)
}
