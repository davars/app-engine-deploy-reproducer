package handlers

import "testing"

func TestMessage(t *testing.T) {
	want := "from unstableapi.Message"
	if Message() != want {
		t.Errorf("got %q, want %q", Message(), want)
	}
}
