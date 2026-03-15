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
	IsPublic          bool       `json:"is_public"`
	GroupName         string     `json:"group_name"`
	OrganizationName  string     `json:"organization_name"`
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
	IsPublic         bool       `json:"is_public"`
	GroupName        string     `json:"group_name"`
	OrganizationName string     `json:"organization_name"`
}

type ReportUpdateReq struct {
	PhotoUrl         []string `json:"photo_url"`
	Comment          string   `json:"comment"`
	LatAtSubmission  float64  `json:"lat_at_submission"`
	LongAtSubmission float64  `json:"long_at_submission"`
	IsPublic         bool     `json:"is_public"`
	GroupName        string   `json:"group_name"`
	OrganizationName string   `json:"organization_name"`
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
	IsPublic          bool       `json:"is_public"`
	GroupName         string     `json:"group_name"`
	OrganizationName  string     `json:"organization_name"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

type ReportListResponse struct {
	Reports                []*Report `json:"reports"`
	Count                  int64     `json:"count"`
	ParticipatedUsersCount int64     `json:"participated_users_count"`
	TotalReportsCount      int64     `json:"total_reports_count"`
	VerifiedReportsCount   int64     `json:"verified_reports_count"`
	InfrastructureRating   float64   `json:"infrastructure_rating"`
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
