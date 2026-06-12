package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Grimmjow06100/helloAi/backend-go/internal/config"
	"github.com/Grimmjow06100/helloAi/backend-go/internal/envfile"
	"github.com/Grimmjow06100/helloAi/backend-go/internal/httpapi"
	"github.com/Grimmjow06100/helloAi/backend-go/internal/prompts"
)

func main() {
	if err := envfile.Load(".env"); err != nil {
		log.Printf("env file not loaded: %v", err)
	}

	cfg := config.Load()

	promptStore, err := prompts.Load(cfg.PromptsDir)
	if err != nil {
		log.Fatalf("load prompts: %v", err)
	}

	server := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           httpapi.NewRouter(cfg, promptStore),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("course-ai backend-go listening on %s", cfg.HTTPAddr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("http server: %v", err)
	}
}
