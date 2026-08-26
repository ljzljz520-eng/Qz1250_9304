package domain

import "sort"

func SortRecords(records []Record) []Record {
	sort.Slice(records, func(i, j int) bool { return records[i].UpdatedAt.Before(records[j].UpdatedAt) })
	return records
}
func FilterActive(users []User) []User {
	out := make([]User, 0, len(users))
	for _, u := range users {
		if u.Active {
			out = append(out, u)
		}
	}
	return out
}
func Labels(records []Record) map[string]int {
	out := map[string]int{}
	for _, r := range records {
		out[r.Label]++
	}
	return out
}
func StatusCounts(records []Record) map[string]int {
	out := map[string]int{}
	for _, r := range records {
		out[NormalizeStatus(r.Status)]++
	}
	return out
}
