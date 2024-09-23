package data_test

import (
	"AhmadAbdelrazik/mark2right/internal/data"
	"testing"
)

func TestChecker(t *testing.T) {
	t.Run("Valid note", func(t *testing.T) {
		given := `Hello, My Name is Ahmad. I am 22 years old.`
		got := data.CheckSpelling(given)

		AssertEmptySlice(t, got)
	})

	t.Run("Invalid note", func(t *testing.T) {
		given := `Hello, My Name izz Ahmad. I am 22 yeras old.`
		got := data.CheckSpelling(given)
		want := []string{"izz", "yeras"}

		AssertSlice(t, got, want)
	})
}
