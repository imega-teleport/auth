package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/imega-teleport/auth/handler"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.WithField("channel", "server-auth")
	logrus.Info("Starting server")
	mux := http.NewServeMux()
	mux.Handle("/", handler.Handler())
	s := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      mux,
		ReadTimeout:  time.Duration(1 * time.Second),
		WriteTimeout: time.Duration(1 * time.Second),
	}

	sigchan := make(chan os.Signal)

	signal.Notify(
		sigchan,
		syscall.SIGTERM,
		syscall.SIGINT,
	)

	go s.ListenAndServe()

	<-sigchan

	logrus.Info("Stopping...")
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	s.Shutdown(ctx)
	logrus.Info("Stopped server...")
}
