package main

import (
	"log"
	"net/http"
	"store55/api"
	"store55/audit"
	"store55/config"
	"store55/service"
	"store55/store"
)

func main() {
	c := config.Load()
	db, e := store.Open(c.DataPath)
	if e != nil {
		log.Fatal(e)
	}
	defer db.Close()
	a := audit.New(db)
	_ = a
	u := service.NewUserService(db)
	r := service.NewRecordService(db)
	s := api.New(u, r)
	mux := http.NewServeMux()
	mux.HandleFunc("/health", s.Health)
	mux.HandleFunc("/users", s.Users)
	mux.HandleFunc("/records", s.Records)
	log.Printf("store55 listening on %s", c.Addr())
	log.Fatal(http.ListenAndServe(c.Addr(), mux))
}
