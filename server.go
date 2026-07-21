package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Em1rGuven/loadBalancer/services"
	"github.com/Em1rGuven/loadBalancer/storage"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app              *fiber.App
	serviceContainer *services.ServiceContainer
	client           *http.Client
	healthChecker    *http.Client
	backends         []*storage.Backend
}

func newServer() *Server {
	app := fiber.New()
	client := &http.Client{Timeout: 4 * time.Second}
	healthChecker := &http.Client{Timeout: 1 * time.Second}
	backends := []*storage.Backend{
		{Port: 8081},
		{Port: 8082},
		{Port: 8083},
	}

	server := &Server{
		app:              app,
		serviceContainer: services.NewServiceContainer(backends, client),
		client:           client,
		healthChecker:    healthChecker,
		backends:         backends,
	}
	server.routes()
	return server
}

func (s *Server) start(port int) {
	go s.healthCheck()
	if err := s.app.Listen(fmt.Sprintf(":%v", port)); err != nil {
		panic(err)
	}
}

func (s *Server) routes() {
	s.app.Get("/score/:movie", func(c *fiber.Ctx) error {
		return s.serviceContainer.GetMovieScore(c)
	})
}
