package models

import "github.com/google/uuid"

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
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
