package main

import (
	"AhmadAbdelrazik/mark2right/internal/note"
	"database/sql"
	"flag"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type config struct {
	env  string
	port int
	db   struct {
		dsn string
	}
}

type application struct {
	cfg         config
	models      *note.Models
	errorLogger *log.Logger
	infoLogger  *log.Logger
}

func main() {
	var cfg config

	flag.StringVar(&cfg.env, "environment", "development", "Environment{development - testing - production}")
	flag.IntVar(&cfg.port, "port", 4000, "Port Number")

	flag.StringVar(&cfg.db.dsn, "db-dsn", "", "Database DNS")

	flag.Parse()

	infoLogger := log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate)
	errorLogger := log.New(os.Stderr, "ERROR\t", log.Ltime|log.Ldate|log.Lshortfile)

	db, err := openDB(cfg.db.dsn)
	if err != nil {
		errorLogger.Fatal(err)
	}

	models := note.NewModels(db)

	app := &application{
		cfg:         cfg,
		models:      models,
		errorLogger: errorLogger,
		infoLogger:  infoLogger,
	}

	err = app.serve()
	if err != nil {
		app.errorLogger.Fatal(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
