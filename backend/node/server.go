package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func newServer() *Server {
	app := fiber.New()
	client := &http.Client{Timeout: 2 * time.Second}

	server := &Server{
		app:    app,
		client: client,
	}
	server.routes()
	return server
}

func (s *Server) routes() {
	s.app.Get("/rating/:movie", func(c *fiber.Ctx) error {
		movie := c.Params("movie")
		req, err := http.NewRequest("GET", fmt.Sprintf("http://www.omdbapi.com/?apikey=%v&t=%v", API_KEY, movie), nil)
		if err != nil {
			return err
		}

		response, err := s.client.Do(req)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		fmt.Println("OMDb Raw Response:", string(body))
		var omdb OMBDResponse
		if err := json.Unmarshal(body, &omdb); err != nil {
			return err
		}
		rating, _ := strconv.ParseFloat(omdb.Rating, 64)
		res := &MovieRatingResponse{
			Rating: rating,
		}
		return c.JSON(res)
	})
	s.app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
}

func (s *Server) start(port int) {
	if err := s.app.Listen(fmt.Sprintf(":%v", port)); err != nil {
		panic(err)
	}
}
