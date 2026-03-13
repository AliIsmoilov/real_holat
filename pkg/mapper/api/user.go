package api

import (
	"real-holat/api/models"
	"real-holat/storage/repo"
)

func ParseLoginWithTgOtpToResponse(token string, user *repo.User, verification *repo.VerificationModel) models.LoginWithTgOtpResponse {
	return models.LoginWithTgOtpResponse{
		AccessToken: token,
		User: models.LoginWithTgOtpUserResp{
			Id:          user.Id,
			FullName:    user.FullName,
			PhoneNumber: user.PhoneNumber,
			Role:        user.Role,
			TgID:        &user.TgID,
			TgUserName:  user.TgUserName,
		},
		TelegramInfo: models.LoginWithTgOtpTelegramInfo{
			TgUserName:     verification.TgUserName,
			TgFirstName:    verification.TgFirstName,
			TgLanguageCode: verification.TgLanguageCode,
		},
	}
}
