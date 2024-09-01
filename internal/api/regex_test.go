package api_test

import (
	"AhmadAbdelrazik/mark2right/internal/api"
	"testing"
)

func TestRegex(t *testing.T) {
	t.Run("Compile Regex", func(t *testing.T) {
		_, err := api.CompileRegex()

		AssertNoError(t, err)
	})

	t.Run("Validate Regex", func(t *testing.T) {
		regexp, _ := api.CompileRegex()
		t.Run("email", func(t *testing.T) {
			given := `ahmadabdelrazik159@gmail.com`
			got := api.CheckRegex(regexp, given)

			AssertBool(t, got, true)
		})

		t.Run("url", func(t *testing.T) {
			given := `www.youtube.com/thisIsAwesome`
			got := api.CheckRegex(regexp, given)

			AssertBool(t, got, true)

		})

		t.Run("date", func(t *testing.T) {
			given := []string{
				`1918-11-11`,
				`11-11-1918`,
				`11/11/1918`,
				`11/Nov/1918`,
			}
			for _, d := range given {
				t.Run(d, func(t *testing.T) {
					got := api.CheckRegex(regexp, d)

					AssertBool(t, got, true)
				})
			}

		})
	})

}
