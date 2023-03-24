package Testing

import "testing"

func TestAddNumber(t *testing.T) {
	got := AddNumber(3, 6)
	want := 7

	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}
