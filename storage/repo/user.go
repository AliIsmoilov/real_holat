package repo

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type UserI interface {
	Create(ctx context.Context, req User) (*User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetByPhone(ctx context.Context, phone string) (*User, error)
	GetByTelegramID(ctx context.Context, telegramID int64) (*User, error)
	GetAll(ctx context.Context, req GetAllUsersReq) (*GetAllUsersResp, error)
	Update(ctx context.Context, req User) (*User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type User struct {
	Id          uuid.UUID  `gorm:"column:id"`
	FullName    string     `gorm:"column:full_name"`
	PhoneNumber string     `gorm:"column:phone_number"`
	Role        string     `gorm:"column:role"`
	TgID        int64      `gorm:"column:tg_id"`
	TgUserName  string     `gorm:"column:tg_user_name"`
	CreatedAt   time.Time  `gorm:"column:created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at"`
}

type GetAllUsersReq struct {
	Limit int32
	Page  int32
	Query string
}

type GetAllUsersResp struct {
	Users []*User
	Count int64
}
