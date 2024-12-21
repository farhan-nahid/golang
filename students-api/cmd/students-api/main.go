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

	"github.com/farhan-nahid/golang/students-api/internal/config"
)

func main(){
	// load config 
	cfg := config.MustLoadConfig()

	// set router
	router := http.NewServeMux()

	// set routes
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the students API"))
	})

	// start server
	server := http.Server {
		Addr: cfg.HTTPServer.Address,
		Handler: router,
	}
	slog.Info("Server is running on", slog.String("", cfg.HTTPServer.Address))


	// Graceful shutdown
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Error starting server: %v\n", err.Error())
		}
	}()

	<-done

	slog.Info("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	if err:= server.Shutdown(ctx); err != nil {
		slog.Error("Server Shutdown Failed:", slog.String("error", err.Error()))
	}
	
	slog.Info("Server Exited Properly")
}