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

func (r *reportRepo) Verify(ctx context.Context, id uuid.UUID) (*repo.Report, error) {
	var report repo.Report
	if err := r.db.WithContext(ctx).
		Table("reports").
		Where("id = ? AND deleted_at IS NULL", id).
		First(&report).Error; err != nil {
		return nil, err
	}

	// Increment verification count
	report.VerificationCount++
	if err := r.db.WithContext(ctx).
		Table("reports").
		Where("id = ?", id).
		Update("verification_count", report.VerificationCount).Error; err != nil {
		return nil, err
	}

	return &report, nil
}

func (r *reportRepo) ReportVerification(ctx context.Context, req repo.ReportVerification) (uuid.UUID, error) {
	if err := r.db.WithContext(ctx).Table("report_verifications").
		Create(&req).Error; err != nil {
		return uuid.Nil, err
	}

	return req.Id, nil
}

func (r *reportRepo) MainPageStats(ctx context.Context) (*repo.MainPageStats, error) {

	var stats repo.MainPageStats

	if err := r.db.WithContext(ctx).
		Table("reports").
		Where("deleted_at IS NULL").
		Count(&stats.TotalReportsCount).Error; err != nil {
		return nil, err
	}

	if err := r.db.WithContext(ctx).
		Table("reports").
		Where("deleted_at IS NULL AND verification_count > 0").
		Count(&stats.VerifiedReportsCount).Error; err != nil {
		return nil, err
	}

	if err := r.db.WithContext(ctx).
		Table("reports").
		Where("deleted_at IS NULL").
		Distinct("user_id").
		Count(&stats.ParticipatedUsersCount).Error; err != nil {
		return nil, err
	}

	if err := r.db.WithContext(ctx).
		Table("infrastructures").
		Where("deleted_at IS NULL").
		Count(&stats.AggregationReport.TotalInfrastructures).Error; err != nil {
		return nil, err
	}

	if err := r.db.WithContext(ctx).
		Table("reports").
		Where("deleted_at IS NULL").
		Distinct("infrastructure_id").
		Count(&stats.AggregationReport.TotalCheckedInfrastructures).Error; err != nil {
		return nil, err
	}

	if stats.AggregationReport.TotalInfrastructures > 0 {
		stats.AggregationReport.InfrastructuresCheckedPercentage = (float64(stats.AggregationReport.TotalCheckedInfrastructures) / float64(stats.AggregationReport.TotalInfrastructures)) * 100
	}

	return &stats, nil
}
