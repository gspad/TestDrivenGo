package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a", 7)
	want := "aaaaaaa"

	if got != want {
		t.Errorf("expected %q but got %q", want, got)
	}
}
