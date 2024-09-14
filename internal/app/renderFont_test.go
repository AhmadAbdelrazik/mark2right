package app_test

import (
	"AhmadAbdelrazik/mark2right/internal/app"
	"testing"
)

func TestFont(t *testing.T) {
	font := app.NewFontRenderer()

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
				got := font.Render(test)
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
				got := font.Render(test)
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
				got := font.Render(test)
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
				got := font.Render(test)
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
				got := font.Render(test)
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
