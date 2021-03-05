package server

import (
	"fmt"
	"interview1-assessment/internal/platform/server/handler/health"
	handler "interview1-assessment/internal/platform/server/handler/tracking"
	"interview1-assessment/internal/tracking"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	counter tracking.CounterRepository
}

func New(host string, port uint, counter tracking.CounterRepository) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		counter: counter,
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

	s.engine.POST("/tracking/events", handler.CreateEventHandler(s.counter))
	s.engine.GET("/tracking/metrics", handler.CreateEventHandlerGetVisits(s.counter))
}
