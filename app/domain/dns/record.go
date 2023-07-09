package dns

import "fmt"

type Record struct {
	Name   string
	Type   string
	TTL    int64
	Values []string
}

func (r Record) String() string {
	return fmt.Sprintf("Record{Name: %q, Type: %s, TTL: %d, Values: %v}", r.Name, r.Type, r.TTL, r.Values)
}
