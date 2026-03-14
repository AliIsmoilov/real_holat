package api

import (
	"real-holat/api/models"
	"real-holat/storage/repo"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func ParseLoginWithPhoneToResponse(token string, user *repo.User) models.LoginResponseWithPhone {
	return models.LoginResponseWithPhone{
		AccessToken: token,
		User: models.LoginWithPhoneUserResp{
			Id:          user.Id,
			FullName:    user.FullName,
			PhoneNumber: user.PhoneNumber,
			Role:        user.Role,
		},
	}
}

func ParseUserToResponse(user *repo.User) models.UserResponse {
	return models.UserResponse{
		Id:          user.Id,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
		Role:        user.Role,
		TgID:        &user.TgID,
		TgUserName:  user.TgUserName,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}

func ParseUsersToListResponse(data *repo.GetAllUsersResp) models.UserListResponse {
	userResponses := make([]*models.UserResponse, len(data.Users))
	for i, user := range data.Users {
		resp := ParseUserToResponse(user)
		userResponses[i] = &resp
	}
	return models.UserListResponse{
		Users: userResponses,
		Count: data.Count,
	}
}

func ToUserAllReqFromQueryParams(limit, page int32, query string) repo.GetAllUsersReq {
	return repo.GetAllUsersReq{
		Limit: limit,
		Page:  page,
		Query: query,
	}
}

func ParseUserCreateRequestToUser(req models.UserCreateRequest) repo.User {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	return repo.User{
		Id:           uuid.New(),
		FullName:     req.FullName,
		PhoneNumber:  req.PhoneNumber,
		PasswordHash: string(hashedPassword),
		Role:         req.Role,
	}
}

func ParseUserUpdateRequestToUser(req models.UserUpdateRequest, existingUser *repo.User) repo.User {
	if req.FullName != "" {
		existingUser.FullName = req.FullName
	}
	if req.PhoneNumber != "" {
		existingUser.PhoneNumber = req.PhoneNumber
	}
	if req.Password != "" {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		existingUser.PasswordHash = string(hashedPassword)
	}
	if req.Role != "" {
		existingUser.Role = req.Role
	}
	return *existingUser
}
