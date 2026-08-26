package config

import (
	"os"
	"strconv"
)

type Config struct {
	DataPath string
	Port     int
	StoreID  string
}

func Load() Config {
	p := os.Getenv("STORE55_DATA")
	if p == "" {
		p = "store55.db"
	}
	port := 8080
	if x, e := strconv.Atoi(os.Getenv("STORE55_PORT")); e == nil && x > 0 {
		port = x
	}
	sid := os.Getenv("STORE55_ID")
	if sid == "" {
		sid = "55"
	}
	return Config{p, port, sid}
}
func (c Config) Addr() string { return ":" + strconv.Itoa(c.Port) }
