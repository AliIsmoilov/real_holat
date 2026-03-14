package models

import (
	"time"

	"github.com/google/uuid"
)

type InfrastructureTypeModelResp struct {
	Id      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	IconUrl string    `json:"icon_url"`
}

// type UpdateInfrastructureTypeReq struct {
// 	Id      uuid.UUID `json:"id"`
// 	Name    string    `json:"name"`
// 	IconUrl string    `json:"icon_url"`
// }

type GetInfrastructureTypesListResp struct {
	InfrastructureTypes []*InfrastructureTypeModelResp `json:"infrastructure_types"`
	Count               int64                          `json:"count"`
}

type InfrastructureType struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	IconUrl   string    `json:"icon_url"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	DeletedAt string    `json:"deleted_at"`
}

type InfrastructureTypeCreateReq struct {
	Name    string `json:"name" binding:"required"`
	IconUrl string `json:"icon_url"`
}

type InfrastructureTypeUpdateReq struct {
	Name    string `json:"name"`
	IconUrl string `json:"icon_url"`
}

type LoginRequest struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type VerificationCode struct {
	ID             uint  `gorm:"primaryKey"`
	TelegramID     int64 `gorm:"index"`
	Phone          string
	Code           string `gorm:"size:6"`
	ExpiresAt      time.Time
	CreatedAt      time.Time
	TgUserName     string `gorm:"column:tg_user_name"`
	TgFirstName    string `gorm:"column:tg_first_name"`
	TgLanguageCode string `gorm:"column:tg_language_code"`
}

type VerifyOTPRequest struct {
	Code string `json:"code" binding:"required,len=6"`
}

type LoginWithTgOtpResponse struct {
	AccessToken  string                     `json:"access_token"`
	User         LoginWithTgOtpUserResp     `json:"user"`
	TelegramInfo LoginWithTgOtpTelegramInfo `json:"telegram_info"`
}

type LoginWithTgOtpUserResp struct {
	Id          uuid.UUID `json:"id"`
	FullName    string    `json:"full_name"`
	PhoneNumber string    `json:"phone_number"`
	Role        string    `json:"role"`
	TgID        *int64    `json:"tg_id,omitempty"`
	TgUserName  string    `json:"tg_user_name"`
}

type LoginWithPhoneUserResp struct {
	Id          uuid.UUID `json:"id"`
	FullName    string    `json:"full_name"`
	PhoneNumber string    `json:"phone_number"`
	Role        string    `json:"role"`
}

type LoginResponseWithPhone struct {
	User        LoginWithPhoneUserResp `json:"user"`
	AccessToken string                 `json:"access_token"`
}

type LoginWithTgOtpTelegramInfo struct {
	TgUserName     string `json:"tg_user_name"`
	TgFirstName    string `json:"tg_first_name"`
	TgLanguageCode string `json:"tg_language_code"`
}

type UserCreateRequest struct {
	FullName    string `json:"full_name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Role        string `json:"role"`
}

type UserUpdateRequest struct {
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	Role        string `json:"role"`
}

type UserResponse struct {
	Id          uuid.UUID `json:"id"`
	FullName    string    `json:"full_name"`
	PhoneNumber string    `json:"phone_number"`
	Role        string    `json:"role"`
	TgID        *int64    `json:"tg_id,omitempty"`
	TgUserName  string    `json:"tg_user_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserListResponse struct {
	Users []*UserResponse `json:"users"`
	Count int64           `json:"count"`
}
