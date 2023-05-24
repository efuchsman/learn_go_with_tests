package hello_world

import "testing"

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, World"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloWithName(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := HelloWithName("Chris")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)

	})
	t.Run("empty string defaults to world", func(t *testing.T) {
		got := HelloWithName("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)

	})
}
