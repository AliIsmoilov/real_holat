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
		Where("email = ?", email).
		First(&user).
		Error; err != nil {
		return nil, err
	}

	return &user, nil
}
