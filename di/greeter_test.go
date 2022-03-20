package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Nik")

	got := buffer.String()
	want := "Hello, Nik"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
