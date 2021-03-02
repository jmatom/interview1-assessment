package bootstrap

import "interview1-assessment/internal/platform/server"

const (
	host = "localhost"
	port = 8080
)

func Run() error {
	srv := server.New(host, port)

	return srv.Run()
}
