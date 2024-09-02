package api_test

import (
	"AhmadAbdelrazik/mark2right/internal/api"
	"log"
	"os"
	"testing"
)

func TestChecker(t *testing.T) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ltime|log.Ldate|log.Lshortfile)
	app, err := api.NewApplication(infoLog, errorLog)

	if err != nil {
		t.Fatalf("app error: %v", err)
	}

	t.Run("Valid note", func(t *testing.T) {
		given := `Hello, My Name is Ahmad. I am 22 years old.`
		got := app.CheckText(given)

		AssertEmptySlice(t, got)
	})

	t.Run("Invalid note", func(t *testing.T) {
		given := `Hello, My Name izz Ahmad. I am 22 yeras old.`
		got := app.CheckText(given)
		want := []string{"izz", "yeras"}

		AssertSlice(t, got, want)
	})
}
