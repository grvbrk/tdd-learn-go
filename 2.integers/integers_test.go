package integers

import "testing"

func TestIntegers(t *testing.T) {
	t.Run("Sum 2 numbers", func(t *testing.T) {
		want := 3
		got := Add(1, 2)

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
