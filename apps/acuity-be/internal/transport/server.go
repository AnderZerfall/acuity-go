package transport

import (
	"Acuity/gen/api/analyzer/v1/analyzerconnect"
	analyzer "Acuity/internal/domain"
	"Acuity/internal/domain/emotion"
	analyzerhandler "Acuity/internal/transport/handlers"
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"Acuity/internal/infrastructure"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo *echo.Echo
}

func NewServer(emotionService emotion.EmotionClassifierService) *Server {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	path, handler := analyzerconnect.NewAnalyzerServiceHandler(&analyzerhandler.AnalyzeServiceServer{
		Analyzer: analyzer.NewPostAnalyzer(emotionService, infrastructure.NewGoogleService(context.Background())),
	})

	e.Any(path+"*", echo.WrapHandler(handler))

	return &Server{echo: e}
}

func (s *Server) Start() {
	go s.handleShutdown()

	s.echo.Logger.Print("Starting server on :8080")
	if err := s.echo.Start(":8080"); err != nil && err != http.ErrServerClosed {
		s.echo.Logger.Fatal(err)
	}

	s.echo.Logger.Print("Graceful shutdown complete.")
}

func (s *Server) handleShutdown() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.echo.Shutdown(shutdownCtx); err != nil {
		s.echo.Logger.Fatal(err)
	}
}
