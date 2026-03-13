package postgres

import (
	"context"

	"real-holat/storage/repo"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type infrastructureTypeRepo struct {
	db *gorm.DB
}

func NewInfrastructureType(db *gorm.DB) repo.InfrastructureTypeI {
	return &infrastructureTypeRepo{
		db: db,
	}
}

func (r *infrastructureTypeRepo) Create(ctx context.Context, infraType repo.InfrastructureType) (*repo.InfrastructureType, error) {
	if err := r.db.WithContext(ctx).
		Table("infrastructure_types").
		Create(&infraType).
		Error; err != nil {
		return nil, err
	}

	return &infraType, nil
}

func (r *infrastructureTypeRepo) Update(ctx context.Context, infraType repo.InfrastructureType) (*repo.InfrastructureType, error) {
	if err := r.db.WithContext(ctx).
		Table("infrastructure_types").
		Where("id = ?", infraType.Id).
		Updates(&infraType).
		Error; err != nil {
		return nil, err
	}

	return &infraType, nil
}

func (r *infrastructureTypeRepo) GetById(ctx context.Context, id uuid.UUID) (*repo.InfrastructureType, error) {
	var infraType repo.InfrastructureType

	if err := r.db.WithContext(ctx).
		Table("infrastructure_types").
		Where("id = ? AND deleted_at IS NULL", id).
		First(&infraType).
		Error; err != nil {
		return nil, err
	}

	return &infraType, nil
}

func (r *infrastructureTypeRepo) GetListInfrastructureTypes(ctx context.Context, req repo.GetAllInfrastructureTypesReq) (repo.GetAllInfrastructureTypesResp, error) {
	var infraTypes []repo.InfrastructureType

	req.Query = "%" + req.Query + "%"
	var count int64

	tx := r.db.WithContext(ctx).
		Table("infrastructure_types").
		Where("deleted_at IS NULL").
		Where("name ILIKE ?", "%"+req.Query+"%")

	err := tx.Count(&count).Error
	if err != nil {
		return repo.GetAllInfrastructureTypesResp{}, err
	}

	if req.Page > 0 && req.Limit > 0 {
		offset := (req.Page - 1) * req.Limit
		tx = tx.Offset(int(offset)).
			Limit(int(req.Limit))
	} else if req.Limit > 0 {
		tx = tx.Limit(int(req.Limit))
	}

	if err := tx.Find(&infraTypes).Error; err != nil {
		return repo.GetAllInfrastructureTypesResp{}, err
	}

	return repo.GetAllInfrastructureTypesResp{
		InfrastructureTypes: infraTypes,
		Count:               count,
	}, nil
}

func (r *infrastructureTypeRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).
		Table("infrastructure_types").
		Where("id = ?", id).
		Update("deleted_at", "NOW()").
		Error
}
