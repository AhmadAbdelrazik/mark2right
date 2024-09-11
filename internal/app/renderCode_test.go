package app_test

import (
	"AhmadAbdelrazik/mark2right/internal/app"
	"testing"
)

func TestCode(t *testing.T) {
	codeRenderer := app.NewCodeRenderer()

	t.Run("Test Inline Code Rendering", func(t *testing.T) {
		tests := []struct {
			name  string
			given string
			want  string
		}{
			{
				name:  "one inline code",
				given: "This is `var a = 3` code",
				want:  "This is <code>var a = 3</code> code",
			},
			{
				name:  "two inline codes",
				given: "this is `var a = 3` and `var b = 4`.",
				want:  "this is <code>var a = 3</code> and <code>var b = 4</code>.",
			},
			{
				name:  "one inline code",
				given: "This is ``var a = 3`` code",
				want:  "This is <code>var a = 3</code> code",
			},
			{
				name:  "two inline codes",
				given: "this is ``var a = 3`` and ``var b = 4``.",
				want:  "this is <code>var a = 3</code> and <code>var b = 4</code>.",
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				got := codeRenderer.Render(test.given)
				if got != test.want {
					t.Fatalf("\ngot:  %q\nwant: %q", got, test.want)
				}
			})
		}
	})

	t.Run("Test Multiline Code Rendering", func(t *testing.T) {
		tests := []struct {
			name  string
			given string
			want  string
		}{
			{
				name:  "Simple Inline",
				given: "```\nHello World\nthis is actually awesome\n```",
				want:  "<code>\nHello World\nthis is actually awesome\n</code>",
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				got := codeRenderer.Render(test.given)
				if got != test.want {
					t.Fatalf("\ngot:  %q\nwant: %q", got, test.want)
				}
			})
		}

	})
}
