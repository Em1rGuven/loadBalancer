package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Em1rGuven/loadBalancer/storage"
)

func (s *Server) healthCheck() {
	for _, backend := range s.backends {
		backend.IsAlive.Store(true)
	}
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		for _, backend := range s.backends {
			go func(backend *storage.Backend) {
				req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:%v/ping", backend.Port), nil)
				if err != nil {
					return
				}

				response, err := s.healthChecker.Do(req)
				if err != nil {
					backend.IsAlive.Store(false)
				} else {
					backend.IsAlive.Store(response.StatusCode == 200)
					response.Body.Close()
				}
			}(backend)
		}
	}
}
