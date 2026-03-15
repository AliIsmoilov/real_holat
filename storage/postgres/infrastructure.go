package postgres

import (
	"context"
	"real-holat/storage/repo"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type infrastructureRepo struct {
	db *gorm.DB
}

func NewInfrastructure(db *gorm.DB) repo.InfrastructureI {
	return &infrastructureRepo{
		db: db,
	}
}

func (r *infrastructureRepo) Create(ctx context.Context, req repo.Infrastructure) (*repo.Infrastructure, error) {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Create infrastructure
	if err := tx.Table("infrastructures").Create(&req).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// // Create check items if any
	// if len(req.CheckItems) > 0 {
	// 	for _, item := range req.CheckItems {
	// 		item.Id = uuid.New()
	// 		item.InfrastructureId = req.Id
	// 	}
	// 	if err := tx.Table("infrastructure_check_items").Create(&req.CheckItems).Error; err != nil {
	// 		tx.Rollback()
	// 		return nil, err
	// 	}
	// }

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &req, nil
}

func (r *infrastructureRepo) GetByID(ctx context.Context, id uuid.UUID) (*repo.Infrastructure, error) {
	var infrastructure repo.Infrastructure

	if err := r.db.WithContext(ctx).
		Table("infrastructures").
		Preload("CheckItems", "deleted_at IS NULL").
		Where("id = ? AND deleted_at IS NULL", id).
		First(&infrastructure).
		Error; err != nil {
		return nil, err
	}

	return &infrastructure, nil
}

func (r *infrastructureRepo) GetAll(ctx context.Context, req repo.GetAllInfrastructuresReq) (*repo.GetAllInfrastructuresResp, error) {
	var infrastructures []*repo.Infrastructure
	var count int64

	if err := r.db.WithContext(ctx).
		Table("infrastructures").
		Where("deleted_at IS NULL").
		Count(&count).
		Error; err != nil {
		return nil, err
	}

	query := r.db.WithContext(ctx).
		Table("infrastructures").
		Preload("CheckItems", "deleted_at IS NULL").
		Where("deleted_at IS NULL")

	if req.Query != "" {
		query = query.Where("name ILIKE ? OR description ILIKE ? OR address ILIKE ?",
			"%"+req.Query+"%", "%"+req.Query+"%", "%"+req.Query+"%")
	}

	orderBy := "created_at DESC"
	if req.Condition == "worst" {
		orderBy = "overall_rating ASC"
	} else if req.Condition == "best" {
		orderBy = "overall_rating DESC"
	}

	query = query.Order(orderBy)

	if req.Tops > 0 {
		query = query.Limit(int(req.Tops))
	} else if req.Page > 0 && req.Limit > 0 {
		offset := (req.Page - 1) * req.Limit
		query = query.Offset(int(offset)).Limit(int(req.Limit))
	} else if req.Limit > 0 {
		query = query.Limit(int(req.Limit))
	}

	if err := query.Find(&infrastructures).Error; err != nil {
		return nil, err
	}

	return &repo.GetAllInfrastructuresResp{
		Infrastructures: infrastructures,
		Count:           count,
	}, nil
}

func (r *infrastructureRepo) Update(ctx context.Context, req repo.Infrastructure) (*repo.Infrastructure, error) {
	if err := r.db.WithContext(ctx).
		Table("infrastructures").
		Where("id = ?", req.Id).
		Updates(&req).
		Error; err != nil {
		return nil, err
	}

	return &req, nil
}

func (r *infrastructureRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).
		Table("infrastructures").
		Where("id = ?", id).
		Update("deleted_at", "NOW()").
		Error
}
