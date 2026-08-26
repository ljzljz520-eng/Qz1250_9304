package domain

import "testing"

func TestCodecRoundTrip(t *testing.T) {
	r := NewRecord("r", "55", "e", "tag")
	b, e := Encode(r)
	if e != nil {
		t.Fatal(e)
	}
	var x Record
	if e = Decode(b, &x); e != nil || x.ID != r.ID {
		t.Fatal(e)
	}
}
