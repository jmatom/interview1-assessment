package server

import (
	"fmt"
	"interview1-assessment/internal/platform/server/handler/health"
	events "interview1-assessment/internal/platform/server/handler/tracking/events"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
}

func New(host string, port uint) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())

	s.engine.POST("/tracking/events", events.CreateEventHandler())
}
