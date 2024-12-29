package server

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int
}

func NewServer() *http.Server {
	// port, _ := strconv.Atoi(os.Getenv("PORT"))
	port := 8080
	NewServer := &Server{
		port: port,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  20 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return server
}
