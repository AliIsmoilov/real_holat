package postgres

import (
	"context"
	"real-holat/storage/repo"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type reportRepo struct {
	db *gorm.DB
}

func NewReport(db *gorm.DB) repo.ReportI {
	return &reportRepo{
		db: db,
	}
}

func (r *reportRepo) Create(ctx context.Context, req repo.Report) (*repo.Report, error) {
	if err := r.db.WithContext(ctx).Table("reports").Create(&req).Error; err != nil {
		return nil, err
	}
	return &req, nil
}

func (r *reportRepo) GetByID(ctx context.Context, id uuid.UUID) (*repo.Report, error) {
	var report repo.Report
	if err := r.db.WithContext(ctx).Table("reports").Where("id = ? AND deleted_at IS NULL", id).First(&report).Error; err != nil {
		return nil, err
	}
	return &report, nil
}

func (r *reportRepo) GetByInfrastructureID(ctx context.Context, infrastructureId uuid.UUID, req repo.GetReportsByInfrastructureReq) (*repo.GetReportsByInfrastructureResp, error) {
	var reports []*repo.Report
	var count int64

	query := r.db.WithContext(ctx).Table("reports").Where("infrastructure_id = ? AND deleted_at IS NULL", infrastructureId)

	if err := query.Model(&repo.Report{}).Count(&count).Error; err != nil {
		return nil, err
	}

	offset := (req.Page - 1) * req.Limit
	if err := query.Order("created_at DESC").Limit(int(req.Limit)).Offset(int(offset)).Find(&reports).Error; err != nil {
		return nil, err
	}

	return &repo.GetReportsByInfrastructureResp{
		Reports: reports,
		Count:   count,
	}, nil
}

func (r *reportRepo) Update(ctx context.Context, req repo.Report) (*repo.Report, error) {
	if err := r.db.WithContext(ctx).Table("reports").Where("id = ?", req.Id).Updates(&req).Error; err != nil {
		return nil, err
	}
	return &req, nil
}

func (r *reportRepo) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).Table("reports").Where("id = ?", id).Update("deleted_at", "NOW()").Error; err != nil {
		return err
	}
	return nil
}
