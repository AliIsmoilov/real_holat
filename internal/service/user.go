package service

import (
	"context"
	"real-holat/storage"
	"real-holat/storage/repo"
)

type UserServiceI interface {
	Create(ctx context.Context, req *repo.User) (*repo.User, error)
	GetByEmail(ctx context.Context, email string) (*repo.User, error)
}

type userService struct {
	strg storage.StorageI
}

func NewUserService(strg storage.StorageI) UserServiceI {
	return &userService{
		strg: strg,
	}
}

func (u *userService) Create(ctx context.Context, req *repo.User) (*repo.User, error) {
	// For now, just call storage. Later can add business logic.
	return u.strg.User().Create(ctx, *req)
}

func (u *userService) GetByEmail(ctx context.Context, email string) (*repo.User, error) {
	return u.strg.User().GetByEmail(ctx, email)
}
