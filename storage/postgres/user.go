package postgres

import (
	"context"
	"real-holat/storage/repo"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) repo.UserI {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) Create(ctx context.Context, req repo.User) (*repo.User, error) {
	if err := u.db.WithContext(ctx).
		Table("users").
		Create(&req).
		Error; err != nil {
		return nil, err
	}

	return &req, nil
}

func (u *userRepo) GetByEmail(ctx context.Context, email string) (*repo.User, error) {
	var user repo.User

	if err := u.db.WithContext(ctx).
		Table("users").
		Where("phone_number = ? AND deleted_at IS NULL", email). // This is a temporary mapping
		First(&user).
		Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepo) GetByPhone(ctx context.Context, phone string) (*repo.User, error) {
	var user repo.User

	if err := u.db.WithContext(ctx).
		Table("users").
		Where("phone_number = ? AND deleted_at IS NULL", phone).
		First(&user).
		Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepo) GetByTelegramID(ctx context.Context, telegramID int64) (*repo.User, error) {
	// This might need a telegram_id column in users table, or we can look it up via phone
	// For now, we'll implement it as a placeholder - might need to add telegram_id to users table
	var user repo.User

	if err := u.db.WithContext(ctx).
		Table("users").
		Where("phone_number = (SELECT phone FROM verification_codes WHERE telegram_id = ? ORDER BY created_at DESC LIMIT 1) AND deleted_at IS NULL", telegramID).
		First(&user).
		Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepo) Update(ctx context.Context, req repo.User) (*repo.User, error) {
	if err := u.db.WithContext(ctx).
		Table("users").
		Where("id = ?", req.Id).
		Updates(&req).
		Error; err != nil {
		return nil, err
	}

	return &req, nil
}
