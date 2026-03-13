package postgres

import (
	"context"
	"time"

	"real-holat/api/models"
	"real-holat/storage/repo"

	"gorm.io/gorm"
)

type verificationRepo struct {
	db *gorm.DB
}

func NewVerification(db *gorm.DB) repo.VerificationStorageI {
	db.AutoMigrate(&models.VerificationCode{})
	return &verificationRepo{db: db}
}

func (r *verificationRepo) Create(ctx context.Context, v *repo.VerificationModel) error {
	m := models.VerificationCode{
		TelegramID:     v.TelegramID,
		Phone:          v.Phone,
		Code:           v.Code,
		ExpiresAt:      v.ExpiresAt,
		CreatedAt:      v.CreatedAt,
		TgUserName:     v.TgUserName,
		TgFirstName:    v.TgFirstName,
		TgLanguageCode: v.TgLanguageCode,
	}
	return r.db.WithContext(ctx).Create(&m).Error
}

func (r *verificationRepo) DeleteByTelegramID(ctx context.Context, telegramID int64) error {
	return r.db.WithContext(ctx).Where("telegram_id = ?", telegramID).Delete(&models.VerificationCode{}).Error
}

func (r *verificationRepo) GetValidByTelegramID(ctx context.Context, telegramID int64, now time.Time) (*repo.VerificationModel, error) {
	var m models.VerificationCode
	res := r.db.WithContext(ctx).Where("telegram_id = ? AND expires_at > ?", telegramID, now).Order("created_at desc").First(&m)
	if res.Error != nil {
		return nil, res.Error
	}
	return &repo.VerificationModel{
		ID:             m.ID,
		TelegramID:     m.TelegramID,
		Phone:          m.Phone,
		Code:           m.Code,
		ExpiresAt:      m.ExpiresAt,
		CreatedAt:      m.CreatedAt,
		TgUserName:     m.TgUserName,
		TgFirstName:    m.TgFirstName,
		TgLanguageCode: m.TgLanguageCode,
	}, nil
}

func (r *verificationRepo) VerifyCode(ctx context.Context, telegramID int64, code string, now time.Time) (bool, error) {
	var m models.VerificationCode
	res := r.db.WithContext(ctx).Where("telegram_id = ? AND code = ? AND expires_at > ?", telegramID, code, now).First(&m)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, res.Error
	}
	// delete used code(s)
	if err := r.db.WithContext(ctx).Where("telegram_id = ?", telegramID).Delete(&models.VerificationCode{}).Error; err != nil {
		return false, err
	}
	return true, nil
}

// GetByCode returns the most recent valid verification record by code alone
func (r *verificationRepo) GetByCode(ctx context.Context, code string, now time.Time) (*repo.VerificationModel, error) {
	var m models.VerificationCode
	res := r.db.WithContext(ctx).Where("code = ? AND expires_at > ?", code, now).Order("created_at desc").First(&m)
	if res.Error != nil {
		return nil, res.Error
	}
	return &repo.VerificationModel{
		ID:             m.ID,
		TelegramID:     m.TelegramID,
		Phone:          m.Phone,
		Code:           m.Code,
		ExpiresAt:      m.ExpiresAt,
		CreatedAt:      m.CreatedAt,
		TgUserName:     m.TgUserName,
		TgFirstName:    m.TgFirstName,
		TgLanguageCode: m.TgLanguageCode,
	}, nil
}

func (r *verificationRepo) VerifyCodeByCode(ctx context.Context, code string, now time.Time) (bool, error) {
	var m models.VerificationCode
	res := r.db.WithContext(ctx).Where("code = ? AND expires_at > ?", code, now).First(&m)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, res.Error
	}
	// delete all codes for that user (cleanup)
	if err := r.db.WithContext(ctx).Where("telegram_id = ?", m.TelegramID).Delete(&models.VerificationCode{}).Error; err != nil {
		return false, err
	}
	return true, nil
}
