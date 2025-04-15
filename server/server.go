package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func NewServer() *Server {
	s := Server{
		engine: gin.Default(),
	}
	s.AddRoutes()
	return &s
}

func (s *Server) Serve(port uint) error {
	if err := s.engine.Run(fmt.Sprintf(":%d", port)); err != nil {
		return fmt.Errorf("failed to run server: %v", err)
	}
	return nil
}
