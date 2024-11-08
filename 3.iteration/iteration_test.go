package iteration

import "testing"

func TestIteraton(t *testing.T) {

	t.Run("Repeat a char n times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Repeat empty char n times", func(t *testing.T) {
		got := Repeat("", 5)
		want := ""

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
