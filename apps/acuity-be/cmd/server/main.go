package main

import (
	"Acuity/gen/api/analyzer/v1/analyzerconnect"
	"Acuity/internal/analyzer"
	"Acuity/internal/text-analyzer/features"
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func gracefulShutdown(apiServer *http.Server, done chan bool) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")
	done <- true
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	path, handler := analyzerconnect.NewAnalyzerServiceHandler(&analyzer.AnalyzeServiceServer{})

	e.Any(path+"*", echo.WrapHandler(handler))

	apiServer := &http.Server{
		Addr:    ":8080",
		Handler: e,
	}

	features.RunLanguageModel()

	done := make(chan bool, 1)

	go gracefulShutdown(apiServer, done)

	log.Println("Starting server on :8080")
	err := apiServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	<-done
	log.Println("Graceful shutdown complete.")
}
