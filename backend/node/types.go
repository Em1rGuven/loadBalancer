package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app    *fiber.App
	client *http.Client
}

type MovieRatingResponse struct {
	Rating float64 `json:"rating"`
}

type OMBDResponse struct {
	Rating string `json:"imdbRating"`
}
