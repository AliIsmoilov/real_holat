package repo

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type InfrastructureTypeI interface {
	Create(context.Context, InfrastructureType) (*InfrastructureType, error)
	Update(context.Context, InfrastructureType) (*InfrastructureType, error)
	GetById(context.Context, uuid.UUID) (*InfrastructureType, error)
	GetListInfrastructureTypes(ctx context.Context, req GetAllInfrastructureTypesReq) (GetAllInfrastructureTypesResp, error)
	Delete(context.Context, uuid.UUID) error
}

type GetAllInfrastructureTypesReq struct {
	Limit int32
	Page  int32
	Query string
}

type GetAllInfrastructureTypesResp struct {
	InfrastructureTypes []InfrastructureType
	Count               int64
}

type InfrastructureType struct {
	Id        uuid.UUID
	Name      string
	IconUrl   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
