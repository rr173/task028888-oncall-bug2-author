package oncall

import "testing"

func TestBug2_RejectsEmptyEngineerID(t *testing.T) {
	_, err := Build(Request{
		Roster: []string{""},
		Start:  "2026-03-02",
		End:    "2026-03-02",
	})
	if err == nil {
		t.Fatal("expected empty engineer ID to be rejected")
	}
}
