package repo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type ReportI interface {
	Create(ctx context.Context, req Report) (*Report, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Report, error)
	GetByInfrastructureID(ctx context.Context, infrastructureId uuid.UUID, req GetReportsByInfrastructureReq) (*GetReportsByInfrastructureResp, error)
	Update(ctx context.Context, req Report) (*Report, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Verify(ctx context.Context, id uuid.UUID) (*Report, error)
	ReportVerification(ctx context.Context, req ReportVerification) (uuid.UUID, error)

	MainPageStats(ctx context.Context) (*MainPageStats, error)
}

type Report struct {
	Id                uuid.UUID      `gorm:"column:id"`
	UserId            *uuid.UUID     `gorm:"column:user_id"`
	InfrastructureId  uuid.UUID      `gorm:"column:infrastructure_id"`
	PhotoUrl          pq.StringArray `gorm:"type:text[]"`
	Comment           string         `gorm:"column:comment"`
	LatAtSubmission   float64        `gorm:"column:lat_at_submission"`
	LongAtSubmission  float64        `gorm:"column:long_at_submission"`
	VerificationCount int            `gorm:"column:verification_count;default:0"`
	IsPublic          bool           `gorm:"column:is_public;default:true"`
	GroupName         string         `gorm:"column:group_name"`
	OrganizationName  string         `gorm:"column:organization_name"`
	CreatedAt         time.Time      `gorm:"column:created_at"`
	UpdatedAt         time.Time      `gorm:"column:updated_at"`
	DeletedAt         *time.Time     `gorm:"column:deleted_at"`
}

type VerifyReportReq struct {
	ReportId uuid.UUID `json:"report_id"`
	UserId   uuid.UUID `json:"user_id"`
}

type VerifyReportResponse struct {
	GivenCoins int `json:"given_coins"`
	TotalCoins int `json:"total_coins"`
}

type MainPageStats struct {
	TotalReportsCount      int64             `json:"total_reports"`
	VerifiedReportsCount   int64             `json:"verified_reports"`
	ParticipatedUsersCount int64             `json:"participated_users"`
	AggregationReport      AggregationReport `json:"aggregation_report"`
}

type AggregationReport struct {
	TotalInfrastructures             int64   `json:"total_infrastructures"`
	TotalCheckedInfrastructures      int64   `json:"total_checked_infrastructures"`
	InfrastructuresCheckedPercentage float64 `json:"infrastructures_checked_percentage"`
}

type ReportVerification struct {
	Id        uuid.UUID `gorm:"column:id"`
	ReportId  uuid.UUID `gorm:"column:report_id"`
	UserId    uuid.UUID `gorm:"column:user_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

type GetReportsByInfrastructureReq struct {
	Limit int32
	Page  int32
}

type GetReportsByInfrastructureResp struct {
	Reports                []*Report
	Count                  int64
	ParticipatedUsersCount int64
	TotalReportsCount      int64
	VerifiedReportsCount   int64
	InfrastructureRating   float64
}
