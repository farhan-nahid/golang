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
	student "github.com/farhan-nahid/golang/students-api/internal/http/handlers"
	"github.com/farhan-nahid/golang/students-api/internal/storage/sqlite"
)

func main(){
	// load config 
	cfg := config.MustLoadConfig()

	// database connection
	 storage, err := sqlite.New(*cfg)

	 if err != nil {
		 slog.Error("Failed to connect to database", slog.String("error", err.Error()))
		 os.Exit(1)
	 }

	 slog.Info("Storage Initialize", slog.String("Env", cfg.Env))


	// set router
	router := http.NewServeMux()

	// set routes
	router.HandleFunc("GET /health-check", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the students API"))
	})


	router.HandleFunc("POST /student", student.New(storage))
	router.HandleFunc("GET /student", student.GetList(storage))
	router.HandleFunc("GET /student/{studentID}", student.GetByID(storage))
	router.HandleFunc("PUT /student/{studentID}", student.UpdateByID(storage))
	router.HandleFunc("DELETE /student/{studentID}", student.DeleteByID(storage))

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