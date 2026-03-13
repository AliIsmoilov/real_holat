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
}

type Report struct {
	Id               uuid.UUID      `gorm:"column:id"`
	UserId           *uuid.UUID     `gorm:"column:user_id"`
	InfrastructureId uuid.UUID      `gorm:"column:infrastructure_id"`
	PhotoUrl         pq.StringArray `gorm:"type:text[]"`
	Comment          string         `gorm:"column:comment"`
	LatAtSubmission  float64        `gorm:"column:lat_at_submission"`
	LongAtSubmission float64        `gorm:"column:long_at_submission"`
	CreatedAt        time.Time      `gorm:"column:created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at"`
	DeletedAt        *time.Time     `gorm:"column:deleted_at"`
}

type GetReportsByInfrastructureReq struct {
	Limit int32
	Page  int32
}

type GetReportsByInfrastructureResp struct {
	Reports []*Report
	Count   int64
}
