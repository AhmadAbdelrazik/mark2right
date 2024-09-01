package api_test

import (
	"reflect"
	"testing"
)

func AssertEmptySlice(t testing.TB, got []string) {
	t.Helper()
	if len(got) > 0 {
		t.Fatalf("got %v, want %v", got, []string{})
	}
}

func AssertSlice(t testing.TB, got, want []string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("\ngot %v\nwant %v", got, want)
	}
}

func AssertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatalf("got %v, want nil", got)
	}
}

func AssertBool(t testing.TB, got, want bool) {
	t.Helper()
	if got != want {
		t.Fatalf(`got "%v", want "%v"`, got, want)
	}
}
