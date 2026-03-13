package models

import (
	"time"

	"github.com/google/uuid"
)

type InfrastructureCheckItem struct {
	Id               uuid.UUID  `json:"id"`
	InfrastructureId uuid.UUID  `json:"infrastructure_id"`
	Category         string     `json:"category"`
	Question         string     `json:"question"`
	IsActive         bool       `json:"is_active"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at"`
}

type InfrastructureCheckItemCreateReq struct {
	Category string `json:"category"`
	Question string `json:"question" binding:"required"`
	IsActive bool   `json:"is_active"`
}

type Infrastructure struct {
	Id             uuid.UUID                  `json:"id"`
	TypeId         uuid.UUID                  `json:"type_id"`
	Name           string                     `json:"name"`
	Description    string                     `json:"description"`
	Address        string                     `json:"address"`
	Latitude       float64                    `json:"latitude"`
	Longitude      float64                    `json:"longitude"`
	Status         string                     `json:"status"`
	OverallRating  int                        `json:"overall_rating"`
	ContractorName string                     `json:"contractor_name"`
	CheckItems     []*InfrastructureCheckItem `json:"check_items"`
	CreatedAt      time.Time                  `json:"created_at"`
	UpdatedAt      time.Time                  `json:"updated_at"`
	DeletedAt      *time.Time                 `json:"deleted_at"`
}

type InfrastructureCreateReq struct {
	TypeId         uuid.UUID                           `json:"type_id" binding:"required"`
	Name           string                              `json:"name" binding:"required"`
	Description    string                              `json:"description"`
	Address        string                              `json:"address" binding:"required"`
	Latitude       float64                             `json:"latitude" binding:"required"`
	Longitude      float64                             `json:"longitude" binding:"required"`
	Status         string                              `json:"status"`
	OverallRating  int                                 `json:"overall_rating"`
	ContractorName string                              `json:"contractor_name"`
	CheckItems     []*InfrastructureCheckItemCreateReq `json:"check_items"`
}

type InfrastructureUpdateReq struct {
	TypeId         uuid.UUID `json:"type_id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Address        string    `json:"address"`
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
	Status         string    `json:"status"`
	OverallRating  int       `json:"overall_rating"`
	ContractorName string    `json:"contractor_name"`
}

type InfrastructureResponse struct {
	Id             uuid.UUID `json:"id"`
	TypeId         uuid.UUID `json:"type_id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Address        string    `json:"address"`
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
	Status         string    `json:"status"`
	OverallRating  int       `json:"overall_rating"`
	ContractorName string    `json:"contractor_name"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type InfrastructureListResponse struct {
	Infrastructures []*Infrastructure `json:"infrastructures"`
	Count           int64             `json:"count"`
}
