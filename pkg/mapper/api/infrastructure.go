package api

import (
	"real-holat/api/models"
	"real-holat/storage/repo"

	"github.com/google/uuid"
)

func ParseInfrastructureRepoToApi(c *repo.Infrastructure) models.Infrastructure {
	checkItems := make([]*models.InfrastructureCheckItem, 0, len(c.CheckItems))
	for _, item := range c.CheckItems {
		checkItems = append(checkItems, &models.InfrastructureCheckItem{
			Id:               item.Id,
			InfrastructureId: item.InfrastructureId,
			Category:         item.Category,
			Question:         item.Question,
			IsActive:         item.IsActive,
			CreatedAt:        item.CreatedAt,
			UpdatedAt:        item.UpdatedAt,
			DeletedAt:        item.DeletedAt,
		})
	}

	return models.Infrastructure{
		Id:                   c.Id,
		TypeId:               c.TypeId,
		Name:                 c.Name,
		Description:          c.Description,
		Address:              c.Address,
		Latitude:             c.Latitude,
		Longitude:            c.Longitude,
		Status:               c.Status,
		OverallRating:        c.OverallRating,
		ContractorName:       c.ContractorName,
		VerifiedReportsCount: c.VerifiedReportsCount,
		CheckItems:           checkItems,
		CreatedAt:            c.CreatedAt,
		UpdatedAt:            c.UpdatedAt,
		DeletedAt:            c.DeletedAt,
	}
}

func ToInfrastructureApiToRepo(req *models.InfrastructureCreateReq) *repo.Infrastructure {
	checkItems := make([]*repo.InfrastructureCheckItem, 0, len(req.CheckItems))

	id := uuid.New()
	for _, item := range req.CheckItems {
		checkItems = append(checkItems, &repo.InfrastructureCheckItem{
			Id:               uuid.New(),
			InfrastructureId: id,
			Category:         item.Category,
			Question:         item.Question,
			IsActive:         item.IsActive,
		})
	}

	return &repo.Infrastructure{
		Id:             id,
		TypeId:         req.TypeId,
		Name:           req.Name,
		Description:    req.Description,
		Address:        req.Address,
		Latitude:       req.Latitude,
		Longitude:      req.Longitude,
		Status:         req.Status,
		OverallRating:  req.OverallRating,
		ContractorName: req.ContractorName,
		CheckItems:     checkItems,
	}
}

func ToInfrastructureListRepoToApi(data *repo.GetAllInfrastructuresResp) models.InfrastructureListResponse {
	resp := models.InfrastructureListResponse{
		Infrastructures: make([]*models.Infrastructure, 0),
		Count:           data.Count,
	}
	for _, elem := range data.Infrastructures {
		infra := ParseInfrastructureRepoToApi(elem)
		resp.Infrastructures = append(resp.Infrastructures, &infra)
	}
	return resp
}
