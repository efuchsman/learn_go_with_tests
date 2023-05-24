package hello_world

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, World"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloWithName(t *testing.T) {
	got := HelloWithName("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
