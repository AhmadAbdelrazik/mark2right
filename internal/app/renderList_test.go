package app_test

import (
	"AhmadAbdelrazik/mark2right/internal/app"
	"testing"
)

func TestCalculateLevel(t *testing.T) {
	list := &app.ListRenderer{}

	given := []string{
		"1. Hi",
		"3. Hi",
		"  - Hi",
		"    - Hi",
		"  1. Hi",
	}

	want := []int{
		1,
		1,
		2,
		3,
		2,
	}

	for i, test := range given {
		t.Run(test, func(t *testing.T) {
			got := list.CalculateListLevel(test)
			AssertEquality(t, got, want[i])
		})
	}
}

func TestList(t *testing.T) {
	t.Run("Test bullet points", func(t *testing.T) {
	})
}
