package dictionary

import "testing"

func TestDictionary(t *testing.T) {
	t.Run("Successful loading of dictionary", func(t *testing.T) {
		got, err := NewDictionary()
		if err != nil {
			t.Fatalf("got %v", err)
		}

		if len(got) == 0 {
			t.Fatalf("got empty dictionary")
		}

	})

}

func TestSearch(t *testing.T) {
	dictionary, err := NewDictionary()
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Run("Search an existing word", func(t *testing.T) {
		given := "Hello"
		got := dictionary.Search(given)
		want := true

		assertEquality(t, got, want)
	})

	t.Run("Search non existing word", func(t *testing.T) {
		given := "walyta"
		got := dictionary.Search(given)
		want := false

		assertEquality(t, got, want)

	})
}

func assertEquality(t testing.TB, got, want any) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}
