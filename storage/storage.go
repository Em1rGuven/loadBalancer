package storage

import (
	"sync/atomic"
)

type Backend struct {
	Port    int
	IsAlive atomic.Bool
}

type SongBPMResponse struct {
	Artist string `json:"artist"`
	Song   string `json:"song"`
	BPM    int    `json:"bpm"`
}
