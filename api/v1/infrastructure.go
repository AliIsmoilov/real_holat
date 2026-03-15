package v1

import (
	"net/http"
	"strconv"

	"real-holat/api/models"
	"real-holat/pkg/libs"
	m "real-holat/pkg/mapper/api"
	"real-holat/storage/repo"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *handlerV1) CreateInfrastructure(ctx *gin.Context) {
	var req models.InfrastructureCreateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		libs.HandleBadRequestErrWithMessage(ctx.Writer, err, "error while binding infrastructure data:")
		return
	}

	data, err := h.service.Infrastructure().Create(ctx.Request.Context(), m.ToInfrastructureApiToRepo(&req))
	if err != nil {
		libs.HandleInternalServerError(ctx.Writer, err)
		return
	}

	libs.WriteJSONWithSuccess(ctx.Writer, m.ParseInfrastructureRepoToApi(data))
}

func (h *handlerV1) UpdateInfrastructure(ctx *gin.Context) {
	id := ctx.Param("id")
	var req models.InfrastructureUpdateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		libs.HandleBadRequestErrWithMessage(ctx.Writer, err, "error while binding infrastructure data:")
		return
	}

	infrastructure, err := h.service.Infrastructure().GetByID(ctx.Request.Context(), id)
	if err != nil {
		libs.HandleNotFoundError(ctx.Writer, err, "infrastructure not found")
		return
	}

	if req.TypeId != uuid.Nil {
		infrastructure.TypeId = req.TypeId
	}
	if req.Name != "" {
		infrastructure.Name = req.Name
	}
	if req.Description != "" {
		infrastructure.Description = req.Description
	}
	if req.Address != "" {
		infrastructure.Address = req.Address
	}
	if req.Latitude != 0 {
		infrastructure.Latitude = req.Latitude
	}
	if req.Longitude != 0 {
		infrastructure.Longitude = req.Longitude
	}
	if req.Status != "" {
		infrastructure.Status = req.Status
	}
	if req.OverallRating != 0 {
		infrastructure.OverallRating = req.OverallRating
	}
	if req.ContractorName != "" {
		infrastructure.ContractorName = req.ContractorName
	}

	data, err := h.service.Infrastructure().Update(ctx.Request.Context(), infrastructure)
	if err != nil {
		libs.HandleInternalServerError(ctx.Writer, err)
		return
	}

	libs.RespondSuccess(ctx, http.StatusOK, m.ParseInfrastructureRepoToApi(data))
}

func (h *handlerV1) GetInfrastructureById(ctx *gin.Context) {
	id := ctx.Param("id")

	data, err := h.service.Infrastructure().GetByID(ctx.Request.Context(), id)
	if err != nil {
		libs.HandleNotFoundError(ctx.Writer, err, "infrastructure not found")
		return
	}

	libs.WriteJSONWithSuccess(ctx.Writer, m.ParseInfrastructureRepoToApi(data))
}

func (h *handlerV1) GetListInfrastructures(ctx *gin.Context) {
	limit := ctx.DefaultQuery("limit", "")
	page := ctx.DefaultQuery("page", "")
	query := ctx.DefaultQuery("query", "")
	tops := ctx.DefaultQuery("tops", "")
	condition := ctx.DefaultQuery("condition", "")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10 // default limit
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1 // default page
	}
	topsInt, err := strconv.Atoi(tops)
	if err != nil {
		topsInt = 0 // default tops
	}

	data, err := h.service.Infrastructure().GetAll(ctx.Request.Context(), repo.GetAllInfrastructuresReq{
		Limit:     int32(limitInt),
		Page:      int32(pageInt),
		Query:     query,
		Tops:      int32(topsInt),
		Condition: condition,
	})
	if err != nil {
		libs.HandleInternalServerError(ctx.Writer, err)
		return
	}

	libs.WriteJSONWithSuccess(ctx.Writer, m.ToInfrastructureListRepoToApi(data))
}

func (h *handlerV1) DeleteInfrastructure(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.service.Infrastructure().Delete(ctx.Request.Context(), id)
	if err != nil {
		libs.HandleInternalServerError(ctx.Writer, err)
		return
	}

	libs.WriteJSONWithSuccess(ctx.Writer, gin.H{"message": "infrastructure deleted successfully"})
}
