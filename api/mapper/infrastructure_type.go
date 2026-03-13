package api

import (
	"real-holat/api/models"
	"real-holat/storage/repo"
)

func ParseInfrastructureTypeRepoToApi2(c *repo.InfrastructureType) models.InfrastructureTypeModelResp {
	return models.InfrastructureTypeModelResp{
		Id:      c.Id,
		Name:    c.Name,
		IconUrl: c.IconUrl,
	}
}
