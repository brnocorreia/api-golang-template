package sampleentities

import "time"

type CreateSamplePayload struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
}

type Sample struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}