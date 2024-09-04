package api_test

import (
	"AhmadAbdelrazik/mark2right/internal/api"
	"testing"
)

func TestRenderer(t *testing.T) {
	t.Run("Render h1", func(t *testing.T) {
		given := `# Hello World`
		got := api.Render(given)
		want := `<h1>Hello World<\h1>`

		AssertEquality(t, got, want)
	})
}
