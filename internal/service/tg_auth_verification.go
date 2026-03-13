package service

import (
	"context"
	"time"

	"real-holat/storage"
	"real-holat/storage/repo"
)

type VerificationServiceI interface {
	Create(ctx context.Context, req *repo.VerificationModel) error
	GetValid(ctx context.Context, telegramID int64) (*repo.VerificationModel, error)
	GetByCode(ctx context.Context, code string, now time.Time) (*repo.VerificationModel, error)
	Verify(ctx context.Context, telegramID int64, code string) (bool, error)
	VerifyByCode(ctx context.Context, code string) (bool, error)
}

type VerificationService struct {
	strg storage.StorageI
}

func NewVerificationService(strg storage.StorageI) *VerificationService {
	return &VerificationService{strg: strg}
}

func (s *VerificationService) Create(ctx context.Context, req *repo.VerificationModel) error {
	// delete old
	if err := s.strg.Verification().DeleteByTelegramID(ctx, req.TelegramID); err != nil {
		return err
	}
	return s.strg.Verification().Create(ctx, req)
}

func (s *VerificationService) GetValid(ctx context.Context, telegramID int64) (*repo.VerificationModel, error) {
	return s.strg.Verification().GetValidByTelegramID(ctx, telegramID, time.Now())
}

func (s *VerificationService) GetByCode(ctx context.Context, code string, now time.Time) (*repo.VerificationModel, error) {
	return s.strg.Verification().GetByCode(ctx, code, now)
}

func (s *VerificationService) Verify(ctx context.Context, telegramID int64, code string) (bool, error) {
	return s.strg.Verification().VerifyCode(ctx, telegramID, code, time.Now())
}

func (s *VerificationService) VerifyByCode(ctx context.Context, code string) (bool, error) {
	return s.strg.Verification().VerifyCodeByCode(ctx, code, time.Now())
}
