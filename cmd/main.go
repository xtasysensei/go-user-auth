package main

import (
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/xtasysensei/go-poll/internal/config"
	"github.com/xtasysensei/go-poll/pkg/database"
	"github.com/xtasysensei/go-poll/pkg/server"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	database.Init(cfg)

	srv, _ := server.StartServer()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", srv.Addr, err)
		}
	}()

	log.Printf("%s is ready to handle requests at %s%s", color.BlueString("Server"), color.GreenString("localhost"), color.GreenString(srv.Addr))
	server.GracefulShutdown(srv)
}
