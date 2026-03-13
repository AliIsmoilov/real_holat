package models

import (
	"time"

	"github.com/google/uuid"
)

type Stadium struct {
	Name         string   `json:"name"`
	City         string   `json:"city" binding:"required,min=2,max=100"`
	District     string   `json:"district,omitempty"`
	Address      string   `json:"address,omitempty"`
	Capacity     int      `json:"capacity"`
	PricePerHour int64    `json:"price_per_hour" binding:"required,gte=0"`
	PhoneNumber  string   `json:"phone_number" binding:"required"`
	SurfaceType  string   `json:"surface_type" binding:"required"`
	StadiumType  string   `json:"stadium_type" binding:"required"`
	Latitude     float64  `json:"latitude"`
	Longitude    float64  `json:"longitude"`
	Photos       []string `json:"photos,omitempty"`
}

type StadiumResponse struct {
	Id           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	City         string    `json:"city"`
	District     string    `json:"district"`
	Address      string    `json:"address"`
	Capacity     int       `json:"capacity"`
	PricePerHour int64     `json:"price_per_hour"`
	PhoneNumber  string    `json:"phone_number"`
	SurfaceType  string    `json:"surface_type"`
	StadiumType  string    `json:"stadium_type"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	Photos       []string  `json:"photos,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
