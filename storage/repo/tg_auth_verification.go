package repo

import (
	"context"
	"time"
)

type VerificationStorageI interface {
	Create(ctx context.Context, v *VerificationModel) error
	DeleteByTelegramID(ctx context.Context, telegramID int64) error
	GetValidByTelegramID(ctx context.Context, telegramID int64, now time.Time) (*VerificationModel, error)
	VerifyCode(ctx context.Context, telegramID int64, code string, now time.Time) (bool, error)

	GetByCode(ctx context.Context, code string, now time.Time) (*VerificationModel, error)
	VerifyCodeByCode(ctx context.Context, code string, now time.Time) (bool, error)
}

type VerificationModel struct {
	ID             uint
	TelegramID     int64
	Phone          string
	Code           string
	ExpiresAt      time.Time
	CreatedAt      time.Time
	TgUserName     string
	TgFirstName    string
	TgLanguageCode string
}
