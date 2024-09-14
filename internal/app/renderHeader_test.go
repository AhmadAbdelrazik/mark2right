package app_test

import (
	"AhmadAbdelrazik/mark2right/internal/app"
	"fmt"
	"testing"
)

func TestHeader(t *testing.T) {
	header := app.NewHeaderRenderer()
	t.Run("Testing Headers", func(t *testing.T) {
		given := []string{
			"# Hello",
			"## My name is Ahmad",
			"### I live in Egypt",
			"#### I am an Egyptian",
			"##### I am a computer Engineer",
			"###### #EngineerRule",
			`# This is entertaining
## I like Apples Too
### Nothing is bad in Banana Government`,
		}

		want := []string{
			"<h1>Hello</h1>",
			"<h2>My name is Ahmad</h2>",
			"<h3>I live in Egypt</h3>",
			"<h4>I am an Egyptian</h4>",
			"<h5>I am a computer Engineer</h5>",
			"<h6>#EngineerRule</h6>",
			`<h1>This is entertaining</h1>
<h2>I like Apples Too</h2>
<h3>Nothing is bad in Banana Government</h3>`,
		}

		for i, test := range given {
			t.Run(fmt.Sprintf("%d hashtags", i+1), func(t *testing.T) {
				got := header.Render(test)
				if got != want[i] {
					t.Fatalf("got %q, want %q", got, want[i])
				}
			})
		}
	})

	t.Run("Bad Headers", func(t *testing.T) {
		given := []string{
			"#Hello",
			" ## My name is Ahmad",
			"###I live in Egypt",
			"I am an Egyptian",
			"###_## I am a computer Engineer",
			"####### #EngineerRule",
			"Hi",
		}

		want := []string{
			"#Hello",
			" ## My name is Ahmad",
			"###I live in Egypt",
			"I am an Egyptian",
			"###_## I am a computer Engineer",
			"####### #EngineerRule",
			"Hi",
		}

		for i, test := range given {
			t.Run(test, func(t *testing.T) {
				got := header.Render(test)
				if got != want[i] {
					t.Fatalf("got %q, want %q", got, want[i])
				}
			})
		}

	})
}
