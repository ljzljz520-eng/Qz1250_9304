package domain

import "encoding/json"

func Encode(v any) ([]byte, error)    { return json.Marshal(v) }
func Decode(data []byte, v any) error { return json.Unmarshal(data, v) }
func CloneRecord(r Record) Record     { b, _ := Encode(r); var out Record; _ = Decode(b, &out); return out }
func CloneUser(u User) User           { b, _ := Encode(u); var out User; _ = Decode(b, &out); return out }
func NormalizeStatus(s string) string {
	switch s {
	case "pending", "review", "approved", "closed", "archived":
		return s
	default:
		return "pending"
	}
}
