package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/udaymungalpara/employee-api/handlers/employee"
	"github.com/udaymungalpara/employee-api/internal/config"
	"github.com/udaymungalpara/employee-api/internal/storage/sqlite"
)

func main() {

	//loading config

	cfg := config.ConfigLoad()

	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal("database error", err)
	}

	router := http.NewServeMux()

	//Handler

	router.HandleFunc("POST /api/employee", employee.New(storage))

	router.HandleFunc("GET /api/employee/{id}", employee.GetId(storage))

	router.HandleFunc("GET /api/employees", employee.GetList(storage))

	router.HandleFunc("DELETE /api/employee/{id}", employee.DeleteById(storage))

	//server setup

	server := http.Server{
		Addr:    cfg.Add,
		Handler: router,
	}

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		slog.Info("Started server", "addr", cfg.Add)
		err := server.ListenAndServe()
		if err != nil {
			slog.Error("Failed to start server", "error", err)
			log.Fatal("Failed to start server")
		}
	}()

	<-done
	slog.Info("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		slog.Info("server not shutdown")
	}

	slog.Info("server shutdown")
}
