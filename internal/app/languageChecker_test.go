package app_test

import (
	"AhmadAbdelrazik/mark2right/internal/app"
	"testing"
)

func TestChecker(t *testing.T) {
	checker := app.NewChecker()

	t.Run("Valid note", func(t *testing.T) {
		given := `Hello, My Name is Ahmad. I am 22 years old.`
		got := checker.CheckSpelling(given)

		AssertEmptySlice(t, got)
	})

	t.Run("Invalid note", func(t *testing.T) {
		given := `Hello, My Name izz Ahmad. I am 22 yeras old.`
		got := checker.CheckSpelling(given)
		want := []string{"izz", "yeras"}

		AssertSlice(t, got, want)
	})
}
