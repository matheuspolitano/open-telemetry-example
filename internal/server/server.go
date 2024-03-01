package server

import (
	"github.com/google/logger"
	"github.com/labstack/echo"
)

type Server struct {
	echo   *echo.Echo
	logger logger.Logger
}

func (s *Server) Run() error {

}
