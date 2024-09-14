package app_test

import (
	"AhmadAbdelrazik/mark2right/internal/app"
	"testing"
)

func TestLinks(t *testing.T) {
	linkRenderer := app.NewLinkRenderer()

	tests := []struct {
		name  string
		given string
		want  string
	}{
		{
			name:  "Simple Link",
			given: "go to facebook from [here](https://www.facebook.com)",
			want:  `go to facebook from <a href="https://www.facebook.com">here</a>`,
		},
		{
			name:  "two Links",
			given: "go to facebook from [here](https://www.facebook.com) or twitter from [here](https://www.x.com)",
			want:  `go to facebook from <a href="https://www.facebook.com">here</a> or twitter from <a href="https://www.x.com">here</a>`,
		},
		{
			name:  "no Links",
			given: "go to nowhere",
			want:  `go to nowhere`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := linkRenderer.Render(test.given)
			if got != test.want {
				t.Fatalf("\ngot:  %q\nwant: %q", got, test.want)
			}
		})
	}
}
