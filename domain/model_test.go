package domain

import "testing"

func TestRecordValidation(t *testing.T) {
	if NewRecord("r", "55", "e", "tag").Valid() == false {
		t.Fatal()
	}
	if (Record{}).Valid() {
		t.Fatal()
	}
}
