package repo

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type InfrastructureI interface {
	Create(ctx context.Context, req Infrastructure) (*Infrastructure, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Infrastructure, error)
	GetAll(ctx context.Context, req GetAllInfrastructuresReq) (*GetAllInfrastructuresResp, error)
	Update(ctx context.Context, req Infrastructure) (*Infrastructure, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type InfrastructureCheckItem struct {
	Id               uuid.UUID  `gorm:"column:id"`
	InfrastructureId uuid.UUID  `gorm:"column:infrastructure_id"`
	Category         string     `gorm:"column:category"`
	Question         string     `gorm:"column:question"`
	IsActive         bool       `gorm:"column:is_active"`
	CreatedAt        time.Time  `gorm:"column:created_at"`
	UpdatedAt        time.Time  `gorm:"column:updated_at"`
	DeletedAt        *time.Time `gorm:"column:deleted_at"`
}

type Infrastructure struct {
	Id             uuid.UUID                  `gorm:"column:id"`
	TypeId         uuid.UUID                  `gorm:"column:type_id"`
	Name           string                     `gorm:"column:name"`
	Description    string                     `gorm:"column:description"`
	Address        string                     `gorm:"column:address"`
	Latitude       float64                    `gorm:"column:latitude"`
	Longitude      float64                    `gorm:"column:longitude"`
	Status         string                     `gorm:"column:status"`
	OverallRating  int                        `gorm:"column:overall_rating"`
	ContractorName string                     `gorm:"column:contractor_name"`
	CheckItems     []*InfrastructureCheckItem `gorm:"foreignKey:InfrastructureId"`
	CreatedAt      time.Time                  `gorm:"column:created_at"`
	UpdatedAt      time.Time                  `gorm:"column:updated_at"`
	DeletedAt      *time.Time                 `gorm:"column:deleted_at"`
}

type GetAllInfrastructuresReq struct {
	Limit int32
	Page  int32
	Query string
}

type GetAllInfrastructuresResp struct {
	Infrastructures []*Infrastructure
	Count           int64
}
