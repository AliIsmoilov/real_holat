package service

import (
	"context"
	"errors"
	"real-holat/pkg/constants"
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
	Verify(ctx context.Context, req *repo.VerifyReportReq) (*repo.VerifyReportResponse, error)
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

func (s *reportService) Verify(ctx context.Context, req *repo.VerifyReportReq) (*repo.VerifyReportResponse, error) {

	report, err := s.strg.Report().GetByID(ctx, req.ReportId)
	if err != nil {
		return nil, err
	}

	// Prevent self-verification
	if report.UserId != nil && *report.UserId == req.UserId {
		return nil, errors.New("users cannot verify their own reports")
	}

	_, err = s.strg.Report().ReportVerification(ctx, repo.ReportVerification{
		Id:       uuid.New(),
		ReportId: req.ReportId,
		UserId:   req.UserId,
	})
	if err != nil {
		return nil, err
	}

	_, err = s.strg.Report().Verify(ctx, req.ReportId)
	if err != nil {
		return nil, err
	}

	if err := s.strg.User().AddCoins(ctx, req.UserId, constants.CoinsForReportVerification); err != nil {
		// Log error but don't fail the request
	}

	user, err := s.strg.User().GetByID(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &repo.VerifyReportResponse{
		GivenCoins: constants.CoinsForReportVerification,
		TotalCoins: user.Coins + constants.CoinsForReportVerification,
	}, nil
}
