package storage

import (
	"sync/atomic"
)

type Backend struct {
	Port    int
	IsAlive atomic.Bool
}

type MovieScoreResponse struct {
	Artist string `json:"artist"`
	Song   string `json:"song"`
	BPM    int    `json:"bpm"`
}
