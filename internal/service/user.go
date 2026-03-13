package service

import (
	"context"
	"real-holat/storage"
	"real-holat/storage/repo"

	"github.com/google/uuid"
)

type UserServiceI interface {
	Create(ctx context.Context, req *repo.User) (*repo.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*repo.User, error)
	GetByPhone(ctx context.Context, phone string) (*repo.User, error)
	GetByTelegramID(ctx context.Context, telegramID int64) (*repo.User, error)
	GetAll(ctx context.Context, req repo.GetAllUsersReq) (*repo.GetAllUsersResp, error)
	Update(ctx context.Context, req *repo.User) (*repo.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
	CreateOrUpdateFromVerification(ctx context.Context, verification *repo.VerificationModel) (*repo.User, error)
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
	return u.strg.User().Create(ctx, *req)
}

func (u *userService) GetByPhone(ctx context.Context, phone string) (*repo.User, error) {
	return u.strg.User().GetByPhone(ctx, phone)
}

func (u *userService) GetByTelegramID(ctx context.Context, telegramID int64) (*repo.User, error) {
	return u.strg.User().GetByTelegramID(ctx, telegramID)
}

func (u *userService) Update(ctx context.Context, req *repo.User) (*repo.User, error) {
	return u.strg.User().Update(ctx, *req)
}

func (u *userService) GetByID(ctx context.Context, id uuid.UUID) (*repo.User, error) {
	return u.strg.User().GetByID(ctx, id)
}

func (u *userService) GetAll(ctx context.Context, req repo.GetAllUsersReq) (*repo.GetAllUsersResp, error) {
	return u.strg.User().GetAll(ctx, req)
}

func (u *userService) Delete(ctx context.Context, id uuid.UUID) error {
	return u.strg.User().Delete(ctx, id)
}

func (u *userService) CreateOrUpdateFromVerification(ctx context.Context, verification *repo.VerificationModel) (*repo.User, error) {

	user, err := u.strg.User().GetByPhone(ctx, verification.Phone)
	if err != nil && err.Error() != "record not found" {
		return nil, err
	}

	if user != nil {
		if user.TgID == 0 {
			user.TgID = verification.TelegramID
		}
		if user.TgUserName == "" {
			user.TgUserName = verification.TgUserName
		}
		if user.FullName == "" {
			user.FullName = verification.TgFirstName
		}
		return u.strg.User().Update(ctx, *user)
	}

	// User doesn't exist, create new one
	newUser := repo.User{
		Id:          uuid.New(),
		FullName:    verification.TgFirstName,
		PhoneNumber: verification.Phone,
		Role:        "citizen",
		TgID:        verification.TelegramID,
		TgUserName:  verification.TgUserName,
	}

	return u.strg.User().Create(ctx, newUser)
}
