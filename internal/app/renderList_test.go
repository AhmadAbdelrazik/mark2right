package app_test

import (
	"AhmadAbdelrazik/mark2right/internal/app"
	"testing"
)

func TestCalculateLevel(t *testing.T) {
	list := app.NewListRenderer()

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

func TestCleanseLine(t *testing.T) {
	list := app.NewListRenderer()

	tests := []struct {
		name  string
		given string
		want  string
	}{
		{
			name:  "level 1",
			given: "1. Hello Ian Long time no see!",
			want:  "Hello Ian Long time no see!",
		},
		{
			name:  "level 2",
			given: "  - First Important Point.",
			want:  "First Important Point.",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := list.CleanseLine(test.given)

			if got != test.want {
				t.Fatalf("\ngot %q\nwant %q", got, test.want)
			}
		})
	}
}

func TestList(t *testing.T) {
	list := app.NewListRenderer()
	t.Run("Test bullet points", func(t *testing.T) {
		tests := []struct {
			name  string
			given string
			want  string
		}{
			{
				name:  "level 0",
				given: "Hello Ian, Long time no see.",
				want:  `Hello Ian, Long time no see.`,
			},
			{
				name:  "level 1",
				given: "- Hello Ian, Long time no see.",
				want: `<ul><li>Hello Ian, Long time no see.</li></ul>
`,
			},
			{
				name:  "level 2",
				given: "  - Hello Ian, Long time no see.",
				want: `<ul>
<ul><li>Hello Ian, Long time no see.</li></ul>
</ul>
`,
			},
			{
				name: "level 3",
				given: `- Why Should we test?
  - Easier development
  - Good Behaviour
  - Improved Performance`,
				want: `<ul><li>Why Should we test?</li></ul>
<ul>
<ul><li>Easier development</li></ul>
</ul>
<ul>
<ul><li>Good Behaviour</li></ul>
</ul>
<ul>
<ul><li>Improved Performance</li></ul>
</ul>
`,
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				got := list.Render(test.given)
				if got != test.want {
					t.Fatalf("\ngot %q\nwant %q", got, test.want)
				}
			})
		}
	})

	t.Run("Test Ordered list", func(t *testing.T) {
		tests := []struct {
			name  string
			given string
			want  string
		}{
			{
				name:  "level 1",
				given: "1. Hello Ian, Long time no see.",
				want:  "<ol start=\"1\"><li>Hello Ian, Long time no see.</li></ol>\n",
			},
			{
				name:  "level 2",
				given: "  1. Hello Ian, Long time no see.",
				want:  "<ul>\n<ol start=\"1\"><li>Hello Ian, Long time no see.</li></ol>\n</ul>\n",
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				got := list.Render(test.given)
				if got != test.want {
					t.Fatalf("\ngot %q\nwant %q", got, test.want)
				}
			})
		}
	})
}
