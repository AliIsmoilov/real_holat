package api

import (
	"real-holat/api/models"
	"real-holat/storage/repo"

	"github.com/google/uuid"
)

func ParseInfrastructureTypeRepoToApi(c *repo.InfrastructureType) models.InfrastructureTypeModelResp {
	return models.InfrastructureTypeModelResp{
		Id:      c.Id,
		Name:    c.Name,
		IconUrl: c.IconUrl,
	}
}

func ToInfrastructureTypeApiToRepo(req *models.InfrastructureTypeCreateReq) *repo.InfrastructureType {
	return &repo.InfrastructureType{
		Id:      uuid.New(),
		Name:    req.Name,
		IconUrl: req.IconUrl,
	}
}

func ToInfrastructureTypeListRepoToApi(data repo.GetAllInfrastructureTypesResp) models.GetInfrastructureTypesListResp {
	resp := models.GetInfrastructureTypesListResp{
		InfrastructureTypes: make([]*models.InfrastructureTypeModelResp, 0),
		Count:               data.Count,
	}
	for _, elem := range data.InfrastructureTypes {
		infraType := ParseInfrastructureTypeRepoToApi(&elem)
		resp.InfrastructureTypes = append(resp.InfrastructureTypes, &infraType)
	}
	return resp
}
