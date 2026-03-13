package postgres

import (
	"context"
	"real-holat/storage/repo"

	"github.com/google/uuid"
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

func (u *userRepo) GetByID(ctx context.Context, id uuid.UUID) (*repo.User, error) {
	var user repo.User

	if err := u.db.WithContext(ctx).
		Table("users").
		Where("id = ? AND deleted_at IS NULL", id).
		First(&user).
		Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepo) GetAll(ctx context.Context, req repo.GetAllUsersReq) (*repo.GetAllUsersResp, error) {
	var users []*repo.User
	var count int64

	// Get total count
	if err := u.db.WithContext(ctx).
		Table("users").
		Where("deleted_at IS NULL").
		Count(&count).
		Error; err != nil {
		return nil, err
	}

	// Get paginated results
	if err := u.db.WithContext(ctx).
		Table("users").
		Where("deleted_at IS NULL").
		Limit(int(req.Limit)).
		Offset((int(req.Page) - 1) * int(req.Limit)).
		Order("created_at DESC").
		Find(&users).
		Error; err != nil {
		return nil, err
	}

	return &repo.GetAllUsersResp{
		Users: users,
		Count: count,
	}, nil
}

func (u *userRepo) Delete(ctx context.Context, id uuid.UUID) error {
	// Soft delete by setting deleted_at
	return u.db.WithContext(ctx).
		Table("users").
		Where("id = ?", id).
		Update("deleted_at", "NOW()").
		Error
}
