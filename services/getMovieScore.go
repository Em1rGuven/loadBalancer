package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync/atomic"

	"github.com/Em1rGuven/loadBalancer/storage"
	"github.com/gofiber/fiber/v2"
)

type ServiceContainer struct {
	backends []*storage.Backend
	client   *http.Client
	counter  atomic.Uint32
}

func (s *ServiceContainer) GetMovieScore(c *fiber.Ctx) error {
	port, err := s.RoundRobin()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:%v/bpm/%v", port, c.Params("song")), nil)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"err": err.Error(),
		})
	}
	response, err := s.client.Do(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"err": err.Error(),
		})
	} else if response.StatusCode >= 300 {
		return c.Status(response.StatusCode).JSON(fiber.Map{
			"err": "unknown error",
		})
	}
	defer response.Body.Close()

	payload, err := io.ReadAll(response.Body)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	res := &storage.MovieScoreResponse{}
	if err := json.Unmarshal(payload, res); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"err": err.Error(),
		})
	}
	return c.JSON(res)
}

func (s *ServiceContainer) RoundRobin() (int, error) {
	n := len(s.backends)
	if n == 0 {
		return 0, fmt.Errorf("no backends configured")
	}

	for i := 0; i < n; i++ {
		idx := s.counter.Add(1)
		backend := s.backends[idx%uint32(n)]

		if backend.IsAlive.Load() {
			return backend.Port, nil
		}
	}
	return 0, fmt.Errorf("no active backends available")
}
