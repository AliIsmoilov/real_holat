package v1

import (
	"net/http"
	"strconv"

	"real-holat/api/models"
	"real-holat/pkg/libs"
	m "real-holat/pkg/mapper/api"
	"real-holat/storage/repo"

	"github.com/gin-gonic/gin"
)

func (h *handlerV1) CreateInfrastructureType(ctx *gin.Context) {
	var req models.InfrastructureTypeCreateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		libs.HandleBadRequestErrWithMessage(ctx.Writer, err, "error while binding ride data:")
		return
	}

	data, err := h.service.InfrastructureType().Create(ctx.Request.Context(), m.ToInfrastructureTypeApiToRepo(&req))
	if err != nil {
		libs.HandleInternalServerError(ctx.Writer, err)
		return
	}

	libs.RespondSuccess(ctx, http.StatusCreated, m.ParseInfrastructureTypeRepoToApi(data))
}

func (h *handlerV1) UpdateInfrastructureType(ctx *gin.Context) {
	id := ctx.Param("id")
	var req models.InfrastructureTypeUpdateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		libs.HandleBadRequestErrWithMessage(ctx.Writer, err, "error while binding ride data:")
		return
	}

	infraType, err := h.service.InfrastructureType().GetById(ctx.Request.Context(), id)
	if err != nil {
		libs.HandleNotFoundError(ctx.Writer, err, "infrastructure type not found")
		return
	}

	if req.Name != "" {
		infraType.Name = req.Name
	}
	if req.IconUrl != "" {
		infraType.IconUrl = req.IconUrl
	}

	data, err := h.service.InfrastructureType().Update(ctx.Request.Context(), infraType)
	if err != nil {
		libs.HandleInternalServerError(ctx.Writer, err)
		return
	}

	libs.RespondSuccess(ctx, http.StatusOK, m.ParseInfrastructureTypeRepoToApi(data))
}

func (h *handlerV1) GetInfrastructureTypeById(ctx *gin.Context) {
	id := ctx.Param("id")

	data, err := h.service.InfrastructureType().GetById(ctx.Request.Context(), id)
	if err != nil {
		libs.HandleNotFoundError(ctx.Writer, err, "infrastructure type not found")
		return
	}

	libs.RespondSuccess(ctx, http.StatusOK, m.ParseInfrastructureTypeRepoToApi(data))
}

func (h *handlerV1) GetListInfrastructureTypes(ctx *gin.Context) {
	limit := ctx.DefaultQuery("limit", "")
	offset := ctx.DefaultQuery("page", "")
	query := ctx.DefaultQuery("query", "")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10 // default limit
	}
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		offsetInt = 1 // default page
	}

	data, err := h.service.InfrastructureType().GetListInfrastructureTypes(ctx.Request.Context(), repo.GetAllInfrastructureTypesReq{
		Limit: int32(limitInt),
		Page:  int32(offsetInt),
		Query: query,
	})
	if err != nil {
		libs.HandleInternalServerError(ctx.Writer, err)
		return
	}

	libs.RespondSuccess(ctx, http.StatusOK, m.ToInfrastructureTypeListRepoToApi(data))
}

func (h *handlerV1) DeleteInfrastructureType(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.service.InfrastructureType().Delete(ctx.Request.Context(), id)
	if err != nil {
		libs.HandleInternalServerError(ctx.Writer, err)
		return
	}

	libs.RespondSuccess(ctx, http.StatusOK,
		"infrastructure type deleted successfully")
}
