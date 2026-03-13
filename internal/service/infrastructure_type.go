package service

import (
	"context"
	"real-holat/storage"
	"real-holat/storage/repo"

	"github.com/google/uuid"
)

type InfrastructureTypeServiceI interface {
	Create(ctx context.Context, req *repo.InfrastructureType) (*repo.InfrastructureType, error)
	Update(ctx context.Context, req *repo.InfrastructureType) (*repo.InfrastructureType, error)
	GetById(ctx context.Context, id string) (*repo.InfrastructureType, error)
	GetListInfrastructureTypes(ctx context.Context, req repo.GetAllInfrastructureTypesReq) (repo.GetAllInfrastructureTypesResp, error)
	Delete(ctx context.Context, id string) error
}

type infrastructureTypeService struct {
	strg storage.StorageI
}

func NewInfrastructureTypeService(strg storage.StorageI) InfrastructureTypeServiceI {
	return &infrastructureTypeService{
		strg: strg,
	}
}

func (s *infrastructureTypeService) Create(ctx context.Context, req *repo.InfrastructureType) (*repo.InfrastructureType, error) {
	return s.strg.InfrastructureType().Create(ctx, *req)
}

func (s *infrastructureTypeService) Update(ctx context.Context, req *repo.InfrastructureType) (*repo.InfrastructureType, error) {
	return s.strg.InfrastructureType().Update(ctx, *req)
}

func (s *infrastructureTypeService) GetById(ctx context.Context, id string) (*repo.InfrastructureType, error) {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.strg.InfrastructureType().GetById(ctx, parsedId)
}

func (s *infrastructureTypeService) GetListInfrastructureTypes(ctx context.Context, req repo.GetAllInfrastructureTypesReq) (repo.GetAllInfrastructureTypesResp, error) {
	return s.strg.InfrastructureType().GetListInfrastructureTypes(ctx, req)
}

func (s *infrastructureTypeService) Delete(ctx context.Context, id string) error {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.strg.InfrastructureType().Delete(ctx, parsedId)
}
