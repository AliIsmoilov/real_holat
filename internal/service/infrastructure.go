package service

import (
	"context"
	"real-holat/storage"
	"real-holat/storage/repo"

	"github.com/google/uuid"
)

type InfrastructureServiceI interface {
	Create(ctx context.Context, req *repo.Infrastructure) (*repo.Infrastructure, error)
	GetByID(ctx context.Context, id string) (*repo.Infrastructure, error)
	GetAll(ctx context.Context, req repo.GetAllInfrastructuresReq) (*repo.GetAllInfrastructuresResp, error)
	Update(ctx context.Context, req *repo.Infrastructure) (*repo.Infrastructure, error)
	Delete(ctx context.Context, id string) error
}

type infrastructureService struct {
	strg storage.StorageI
}

func NewInfrastructureService(strg storage.StorageI) InfrastructureServiceI {
	return &infrastructureService{
		strg: strg,
	}
}

func (s *infrastructureService) Create(ctx context.Context, req *repo.Infrastructure) (*repo.Infrastructure, error) {
	return s.strg.Infrastructure().Create(ctx, *req)
}

func (s *infrastructureService) GetByID(ctx context.Context, id string) (*repo.Infrastructure, error) {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.strg.Infrastructure().GetByID(ctx, parsedId)
}

func (s *infrastructureService) GetAll(ctx context.Context, req repo.GetAllInfrastructuresReq) (*repo.GetAllInfrastructuresResp, error) {
	return s.strg.Infrastructure().GetAll(ctx, req)
}

func (s *infrastructureService) Update(ctx context.Context, req *repo.Infrastructure) (*repo.Infrastructure, error) {
	return s.strg.Infrastructure().Update(ctx, *req)
}

func (s *infrastructureService) Delete(ctx context.Context, id string) error {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.strg.Infrastructure().Delete(ctx, parsedId)
}
