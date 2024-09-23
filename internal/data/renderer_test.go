package data_test

import (
	"AhmadAbdelrazik/mark2right/internal/data"
	"fmt"
	"testing"
)

func TestRenderer(t *testing.T) {
	tests := []struct {
		name  string
		given string
		want  string
	}{
		{
			name: "Simple Test",
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
			got := data.Render(test.given)
			if got != test.want {
				t.Fatalf("\ngot:  %q\nwant: %q", got, test.want)
			}
		})
	}
}

func TestCode(t *testing.T) {
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
			{
				name:  "no inline code",
				given: "this is nothing",
				want:  "this is nothing",
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				got := data.RenderCode(test.given)
				if got != test.want {
					t.Fatalf("\ngot:  %q\nwant: %q", got, test.want)
				}
			})
		}
	})
}

func TestFont(t *testing.T) {
	t.Run("1 star", func(t *testing.T) {
		given := []string{
			`*Hi*`,
			`*i*`,
			`My Name is *Ahmad*`,
			`I am an* Egyptian Student*`,
			`****`,
			`Hi!, I am *Ahmad Abdelrazik*
I am a student at *Suez Canal University*`,
			"hi",
		}

		want := []string{
			`<i>Hi</i>`,
			`<i>i</i>`,
			`My Name is <i>Ahmad</i>`,
			`I am an* Egyptian Student*`,
			`****`,
			`Hi!, I am <i>Ahmad Abdelrazik</i>
I am a student at <i>Suez Canal University</i>`,
			"hi",
		}

		for i, test := range given {
			t.Run(test, func(t *testing.T) {
				got := data.RenderFont(test)
				AssertStringEquality(t, got, want[i])
			})
		}

	})

	t.Run("2 stars", func(t *testing.T) {
		given := []string{
			`**Hi**`,
			`**i**`,
			`My Name is **Ahmad**`,
			`I am an** Egyptian Student**`,
			`****`,
		}

		want := []string{
			`<b>Hi</b>`,
			`<b>i</b>`,
			`My Name is <b>Ahmad</b>`,
			`I am an** Egyptian Student**`,
			`****`,
		}

		for i, test := range given {
			t.Run(test, func(t *testing.T) {
				got := data.RenderFont(test)
				AssertStringEquality(t, got, want[i])
			})
		}

	})

	t.Run("3 stars", func(t *testing.T) {
		given := []string{
			`***Hi***`,
			`***i***`,
			`My Name is ***Ahmad***`,
			`I am an*** Egyptian Student***`,
			`****`,
		}

		want := []string{
			`<i><b>Hi</b></i>`,
			`<i><b>i</b></i>`,
			`My Name is <i><b>Ahmad</b></i>`,
			`I am an*** Egyptian Student***`,
			`****`,
		}

		for i, test := range given {
			t.Run(test, func(t *testing.T) {
				got := data.RenderFont(test)
				AssertStringEquality(t, got, want[i])
			})
		}

	})

	t.Run("4 stars", func(t *testing.T) {
		given := []string{
			`****Hi****`,
			`****i****`,
			`My Name is ****Ahmad****`,
			`I am an**** Egyptian Student****`,
			`****`,
		}

		want := []string{
			`*<i><b>Hi</b></i>*`,
			`*<i><b>i</b></i>*`,
			`My Name is *<i><b>Ahmad</b></i>*`,
			`I am an**** Egyptian Student****`,
			`****`,
		}

		for i, test := range given {
			t.Run(test, func(t *testing.T) {
				got := data.RenderFont(test)
				AssertStringEquality(t, got, want[i])
			})
		}

	})
	t.Run("5 stars", func(t *testing.T) {
		given := []string{
			`*****Hi*****`,
			`*****i*****`,
			`My Name is *****Ahmad*****`,
			`I am an***** Egyptian Student*****`,
			`****`,
		}

		want := []string{
			`**<i><b>Hi</b></i>**`,
			`**<i><b>i</b></i>**`,
			`My Name is **<i><b>Ahmad</b></i>**`,
			`I am an***** Egyptian Student*****`,
			`****`,
		}

		for i, test := range given {
			t.Run(test, func(t *testing.T) {
				got := data.RenderFont(test)
				AssertStringEquality(t, got, want[i])
			})
		}

	})
}

func AssertStringEquality(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Fatalf("\ngot %q\nwant %q", got, want)
	}
}

func TestHeader(t *testing.T) {
	t.Run("Testing Headers", func(t *testing.T) {
		given := []string{
			"# Hello",
			"## My name is Ahmad",
			"### I live in Egypt",
			"#### I am an Egyptian",
			"##### I am a computer Engineer",
			"###### #EngineerRule",
			`# This is entertaining
## I like apples Too
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
<h2>I like apples Too</h2>
<h3>Nothing is bad in Banana Government</h3>`,
		}

		for i, test := range given {
			t.Run(fmt.Sprintf("%d hashtags", i+1), func(t *testing.T) {
				got := data.RenderHeader(test)
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
				got := data.RenderHeader(test)
				if got != want[i] {
					t.Fatalf("got %q, want %q", got, want[i])
				}
			})
		}

	})
}

func TestLinks(t *testing.T) {

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
			got := data.RenderLink(test.given)
			if got != test.want {
				t.Fatalf("\ngot:  %q\nwant: %q", got, test.want)
			}
		})
	}
}

func TestList(t *testing.T) {
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
				want:  `<ul><li>Hello Ian, Long time no see.</li></ul>`,
			},
			{
				name:  "level 2",
				given: "  - Hello Ian, Long time no see.",
				want: `<ul>
<ul><li>Hello Ian, Long time no see.</li></ul>
</ul>`,
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
</ul>`,
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				got := data.RenderList(test.given)
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
				want:  "<ol start=\"1\"><li>Hello Ian, Long time no see.</li></ol>",
			},
			{
				name:  "level 2",
				given: "  1. Hello Ian, Long time no see.",
				want:  "<ul>\n<ol start=\"1\"><li>Hello Ian, Long time no see.</li></ol>\n</ul>",
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				got := data.RenderList(test.given)
				if got != test.want {
					t.Fatalf("\ngot %q\nwant %q", got, test.want)
				}
			})
		}
	})
}
