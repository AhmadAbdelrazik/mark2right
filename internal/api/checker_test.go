package api_test

import (
	"AhmadAbdelrazik/mark2right/internal/api"
	"AhmadAbdelrazik/mark2right/internal/dictionary"
	"testing"
)

func TestChecker(t *testing.T) {
	d, _ := dictionary.NewDictionary()
	t.Run("Valid note", func(t *testing.T) {
		given := `Hello, My Name is Ahmad. I am 22 years old.`
		got := api.CheckText(given, d)

		AssertEmptySlice(t, got)
	})

	t.Run("Invalid note", func(t *testing.T) {
		given := `Hello, My Name izz Ahmad. I am 22 yeras old.`
		got := api.CheckText(given, d)
		want := []string{"izz", "yeras"}

		AssertSlice(t, got, want)
	})
}
