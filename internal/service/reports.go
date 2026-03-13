package service

import (
	"context"
	"real-holat/storage"
	"real-holat/storage/repo"

	"github.com/google/uuid"
)

type ReportServiceI interface {
	Create(ctx context.Context, req *repo.Report) (*repo.Report, error)
	GetByID(ctx context.Context, id string) (*repo.Report, error)
	GetByInfrastructureID(ctx context.Context, infrastructureId string, req repo.GetReportsByInfrastructureReq) (*repo.GetReportsByInfrastructureResp, error)
	Update(ctx context.Context, req *repo.Report) (*repo.Report, error)
	Delete(ctx context.Context, id string) error
}

type reportService struct {
	strg storage.StorageI
}

func NewReportService(strg storage.StorageI) ReportServiceI {
	return &reportService{
		strg: strg,
	}
}

func (s *reportService) Create(ctx context.Context, req *repo.Report) (*repo.Report, error) {
	return s.strg.Report().Create(ctx, *req)
}

func (s *reportService) GetByID(ctx context.Context, id string) (*repo.Report, error) {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.strg.Report().GetByID(ctx, parsedId)
}

func (s *reportService) GetByInfrastructureID(ctx context.Context, infrastructureId string, req repo.GetReportsByInfrastructureReq) (*repo.GetReportsByInfrastructureResp, error) {
	parsedId, err := uuid.Parse(infrastructureId)
	if err != nil {
		return nil, err
	}
	return s.strg.Report().GetByInfrastructureID(ctx, parsedId, req)
}

func (s *reportService) Update(ctx context.Context, req *repo.Report) (*repo.Report, error) {
	return s.strg.Report().Update(ctx, *req)
}

func (s *reportService) Delete(ctx context.Context, id string) error {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.strg.Report().Delete(ctx, parsedId)
}
