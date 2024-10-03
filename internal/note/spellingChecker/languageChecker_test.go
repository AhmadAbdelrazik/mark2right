package spellingchecker

import (
	"reflect"
	"testing"
)

func TestChecker(t *testing.T) {
	t.Run("Valid note", func(t *testing.T) {
		given := `Hello, My Name is Ahmad. I am 22 years old.`
		got := CheckSpelling(given)

		assertEmptySlice(t, got)
	})

	t.Run("Invalid note", func(t *testing.T) {
		given := `Hello, My Name izz Ahmad. I am 22 yeras old.`
		got := CheckSpelling(given)
		want := []string{"izz", "yeras"}

		assertSlice(t, got, want)
	})
}

func assertEmptySlice(t testing.TB, got []string) {
	t.Helper()
	if len(got) > 0 {
		t.Fatalf("got %v, want %v", got, []string{})
	}
}

func assertSlice(t testing.TB, got, want []string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("\ngot %v\nwant %v", got, want)
	}
}
