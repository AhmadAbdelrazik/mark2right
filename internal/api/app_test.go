package api_test

import (
	"AhmadAbdelrazik/mark2right/internal/api"
	"log"
	"os"
	"testing"
)

func TestApp(t *testing.T) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ltime|log.Ldate|log.Lshortfile)
	_, err := api.NewApplication(infoLog, errorLog)

	AssertNoError(t, err)
}
