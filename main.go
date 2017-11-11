package main

import (
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/imega-teleport/auth/api"
	"github.com/imega-teleport/auth/server"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logger := logrus.WithField("channel", "server-auth")
	logrus.Info("Starting server")

	rOpts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(server.RecoveryHandler),
	}
	grpcSrv := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_logrus.UnaryServerInterceptor(logger),
			grpc_recovery.UnaryServerInterceptor(rOpts...),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(rOpts...),
		),
	)

	auth.RegisterAuthBasicServer(grpcSrv, server.NewServer())
	l, _ := net.Listen("tcp", "0.0.0.0:9000")
	go grpcSrv.Serve(l)

	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := auth.RegisterAuthBasicHandlerFromEndpoint(context.Background(), gwmux, "0.0.0.0:9000", opts)
	if err != nil {
		logrus.Errorf("Error on startup %s", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
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
