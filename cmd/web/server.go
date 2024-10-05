package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (app *application) serve() error {
	srv := http.Server{
		Addr:         fmt.Sprintf("localhost:%d", app.cfg.port),
		ErrorLog:     app.errorLogger,
		Handler:      app.Routes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  3 * time.Minute,
	}

	shutdownError := make(chan error, 1)

	go func() {
		quitSignal := make(chan os.Signal, 1)
		signal.Notify(quitSignal, syscall.SIGTERM, syscall.SIGINT)

		s := <-quitSignal
		app.infoLogger.Printf("recieved %s, shutting down the server", s.String())

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			shutdownError <- err
			return
		}

		app.infoLogger.Printf("completing background tasks")

		app.wg.Wait()

		shutdownError <- nil
	}()

	app.infoLogger.Printf("starting server at port %d", app.cfg.port)

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	app.infoLogger.Printf("server has stopped")

	return nil
}
