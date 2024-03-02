package server

import (
	"context"
	"net/http"
	"time"

	"github.com/google/logger"
	"github.com/labstack/echo/v4"
	"github.com/matheuspolitano/open-telemetry-example/config"
)

type Server struct {
	echo   *echo.Echo
	logger logger.Logger
	cfg    config.Config
	svc    ServerServices
}

func NewServer(logger logger.Logger, cfg config.Config, svc ServerServices) *Server {
	return &Server{
		echo:   echo.New(),
		logger: logger,
		cfg:    cfg,
		svc:    svc,
	}
}

func (s *Server) Run() error {
	s.echo.Server.ReadTimeout = time.Second * s.cfg.Server.ReadTimeout
	s.echo.Server.WriteTimeout = time.Second * s.cfg.Server.WriteTimeout
	httpServer := &http.Server{
		Addr:         s.cfg.Server.Port,
		ReadTimeout:  time.Second * s.cfg.Server.ReadTimeout,
		WriteTimeout: time.Second * s.cfg.Server.WriteTimeout,
	}

	go func() {
		s.logger.Info("Server is listening on PORT: %s", s.cfg.Server.Port)
		if err := s.echo.StartServer(httpServer); err != nil {
			s.logger.Fatal("Wrrot stating Server", err)
		}
	}()
	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
	return s.echo.Server.Shutdown(ctx)
}
