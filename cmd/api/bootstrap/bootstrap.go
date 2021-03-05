package bootstrap

import (
	"interview1-assessment/internal/platform/server"
	counterMemoryHll "interview1-assessment/internal/platform/storage/memory/hll"
)

// "github.com/CodelyTV/go-hexagonal_http_api-course/04-01-application-service/internal/platform/storage/mysql"

const (
	host = "localhost"
	port = 8080
)

func Run() error {
	counter := counterMemoryHll.NewCounterRepository()

	srv := server.New(host, port, counter)

	return srv.Run()
}
