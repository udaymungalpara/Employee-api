package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/udaymungalpara/employee-api/internal/config"
)

func main() {

	//loading config

	cfg := config.ConfigLoad()

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("employee are"))

	})

	//server setup

	server := http.Server{
		Addr:    cfg.Add,
		Handler: router,
	}
	fmt.Println("strated server")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start server")
	}

}
