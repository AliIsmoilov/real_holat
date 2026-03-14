package api

import (
	"real-holat/api/models"
	"real-holat/storage/repo"

	"github.com/google/uuid"
)

func ParseReportRepoToApi(c *repo.Report) models.Report {
	return models.Report{
		Id:               c.Id,
		UserId:           c.UserId,
		InfrastructureId: c.InfrastructureId,
		PhotoUrl:         c.PhotoUrl,
		Comment:          c.Comment,
		LatAtSubmission:  c.LatAtSubmission,
		LongAtSubmission: c.LongAtSubmission,
		CreatedAt:        c.CreatedAt,
		UpdatedAt:        c.UpdatedAt,
		DeletedAt:        c.DeletedAt,
	}
}

func ParseReportRepoToResponse(c *repo.Report) *models.ReportResponse {
	return &models.ReportResponse{
		Id:               c.Id,
		UserId:           c.UserId,
		InfrastructureId: c.InfrastructureId,
		PhotoUrl:         c.PhotoUrl,
		Comment:          c.Comment,
		LatAtSubmission:  c.LatAtSubmission,
		LongAtSubmission: c.LongAtSubmission,
		CreatedAt:        c.CreatedAt,
		UpdatedAt:        c.UpdatedAt,
	}
}

func ToReportApiToRepo(req *models.ReportCreateReq) *repo.Report {
	return &repo.Report{
		Id:               uuid.New(),
		UserId:           req.UserId,
		InfrastructureId: req.InfrastructureId,
		PhotoUrl:         req.PhotoUrl,
		Comment:          req.Comment,
		LatAtSubmission:  req.LatAtSubmission,
		LongAtSubmission: req.LongAtSubmission,
	}
}

func ToReportListRepoToApi(data *repo.GetReportsByInfrastructureResp) models.ReportListResponse {
	resp := models.ReportListResponse{
		Reports: make([]*models.Report, 0),
		Count:   data.Count,
	}
	for _, elem := range data.Reports {
		report := ParseReportRepoToApi(elem)
		resp.Reports = append(resp.Reports, &report)
	}
	return resp
}
