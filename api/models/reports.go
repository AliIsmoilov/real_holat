package models

import (
	"time"

	"github.com/google/uuid"
)

type Report struct {
	Id                uuid.UUID  `json:"id"`
	UserId            *uuid.UUID `json:"user_id"`
	InfrastructureId  uuid.UUID  `json:"infrastructure_id"`
	PhotoUrl          []string   `json:"photo_url"`
	Comment           string     `json:"comment"`
	LatAtSubmission   float64    `json:"lat_at_submission"`
	LongAtSubmission  float64    `json:"long_at_submission"`
	VerificationCount int        `json:"verification_count"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at"`
}

type ReportCreateReq struct {
	UserId           *uuid.UUID `json:"user_id"`
	InfrastructureId uuid.UUID  `json:"infrastructure_id" binding:"required"`
	PhotoUrl         []string   `json:"photo_url" binding:"required"`
	Comment          string     `json:"comment"`
	LatAtSubmission  float64    `json:"lat_at_submission" binding:"required"`
	LongAtSubmission float64    `json:"long_at_submission" binding:"required"`
}

type ReportUpdateReq struct {
	PhotoUrl         []string `json:"photo_url"`
	Comment          string   `json:"comment"`
	LatAtSubmission  float64  `json:"lat_at_submission"`
	LongAtSubmission float64  `json:"long_at_submission"`
}

type ReportResponse struct {
	Id                uuid.UUID  `json:"id"`
	UserId            *uuid.UUID `json:"user_id"`
	InfrastructureId  uuid.UUID  `json:"infrastructure_id"`
	PhotoUrl          []string   `json:"photo_url"`
	Comment           string     `json:"comment"`
	LatAtSubmission   float64    `json:"lat_at_submission"`
	LongAtSubmission  float64    `json:"long_at_submission"`
	VerificationCount int        `json:"verification_count"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

type ReportListResponse struct {
	Reports []*Report `json:"reports"`
	Count   int64     `json:"count"`
}

type CreateReportResponse struct {
	Report     *ReportResponse `json:"report"`
	GivenCoins int             `json:"given_coins"`
	TotalCoins int             `json:"total_coins"`
}

type VerifyReportResponse struct {
	Report     *ReportResponse `json:"report"`
	GivenCoins int             `json:"given_coins"`
	TotalCoins int             `json:"total_coins"`
}
