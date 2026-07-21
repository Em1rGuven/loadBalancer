package storage

import (
	"sync/atomic"
)

type Backend struct {
	Port    int
	IsAlive atomic.Bool
}

type MovieRatingResponse struct {
	Rating float64 `json:"rating"`
}
