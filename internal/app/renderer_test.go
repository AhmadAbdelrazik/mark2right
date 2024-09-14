package app_test

import (
	"AhmadAbdelrazik/mark2right/internal/app"
	"testing"
)

func TestRenderer(t *testing.T) {
	r := app.NewRenderer()

	tests := []struct {
		name  string
		given string
		want  string
	}{
		{
			name: "Simple Test with endline",
			given: `# Hello World
This is really good.
- I like bullet points.
1. I like Numbered lines too.
I *Really* **Like** ***Markdown***

I like programming too.
`,
			want: `<h1>Hello World</h1>
This is really good.
<ul><li>I like bullet points.</li></ul>
<ol start="1"><li>I like Numbered lines too.</li></ol>
I <i>Really</i> <b>Like</b> <i><b>Markdown</b></i>

I like programming too.
`,
		},
		{
			name: "Simple Test",
			given: `# Hello World
This is really good.
- I like bullet points.
1. I like Numbered lines too.
I *Really* **Like** ***Markdown***

I like programming too.`,
			want: `<h1>Hello World</h1>
This is really good.
<ul><li>I like bullet points.</li></ul>
<ol start="1"><li>I like Numbered lines too.</li></ol>
I <i>Really</i> <b>Like</b> <i><b>Markdown</b></i>

I like programming too.`,
		},
		{
			name:  "Code block",
			given: "# Example Shell Script\n```shell\ngo mod tidy\n```",
			want: `<h1>Example Shell Script</h1>
<code>
go mod tidy
</code>`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := r.Render(test.given)
			if got != test.want {
				t.Fatalf("\ngot:  %q\nwant: %q", got, test.want)
			}
		})
	}
}
