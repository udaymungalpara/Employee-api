package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/udaymungalpara/employee-api/internal/config"
	"github.com/udaymungalpara/employee-api/internal/handlers/employee"
)

func main() {

	//loading config

	cfg := config.ConfigLoad()

	router := http.NewServeMux()

	//Handler

	router.HandleFunc("GET /api/employee", employee.New())

	//server setup

	server := http.Server{
		Addr:    cfg.Add,
		Handler: router,
	}

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		fmt.Println("strated server")
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server")
		}
	}()

	<-done
	slog.Info("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Info("server not shutdown")
	}

	slog.Info("server shutdown")
}
