package v1

import (
	"fmt"
	"strconv"

	"real-holat/api/models"
	"real-holat/pkg/constants"
	"real-holat/pkg/libs"
	m "real-holat/pkg/mapper/api"
	"real-holat/storage/repo"

	"github.com/gin-gonic/gin"
)

func (h *handlerV1) CreateReport(ctx *gin.Context) {
	var req models.ReportCreateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		libs.HandleBadRequestErrWithMessage(ctx.Writer, err, "error while binding report data:")
		return
	}

	userID, err := libs.GetUserIDFromToken(ctx.Request.Header.Get(constants.AuthorizationHeader))
	if err != nil {
		libs.HandleBadRequestErrWithMessage(ctx.Writer, err, "error in getting userId from token")
		return
	}
	req.UserId = &userID

	fmt.Printf("Creating report for user ID: %s\n", userID)

	data, err := h.service.Report().Create(ctx.Request.Context(), m.ToReportApiToRepo(&req))
	if err != nil {
		libs.HandleInternalServerError(ctx.Writer, err)
		return
	}
	fmt.Printf("Report created with ID: %s\n", data.Id)

	// Award coins for creating a report
	if err := h.service.User().AddCoins(ctx.Request.Context(), *req.UserId, constants.CoinsForReport); err != nil {
		// Log error but don't fail the request
		// You might want to add proper logging here
	}

	// Get updated user data to include coins in response
	user, err := h.service.User().GetByID(ctx.Request.Context(), *req.UserId)
	if err != nil {
		libs.HandleInternalServerError(ctx.Writer, err)
		return
	}

	response := models.CreateReportResponse{
		Report:     m.ParseReportRepoToResponse(data),
		GivenCoins: constants.CoinsForReport,
		TotalCoins: user.Coins + constants.CoinsForReport,
	}

	fmt.Print("report created with ID: ", data.Id)

	libs.WriteJSONWithSuccess(ctx.Writer, response)
}

func (h *handlerV1) UpdateReport(ctx *gin.Context) {
	id := ctx.Param("id")
	var req models.ReportUpdateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		libs.HandleBadRequestErrWithMessage(ctx.Writer, err, "error while binding report data:")
		return
	}

	report, err := h.service.Report().GetByID(ctx.Request.Context(), id)
	if err != nil {
		libs.HandleNotFoundError(ctx.Writer, err, "report not found")
		return
	}

	if len(req.PhotoUrl) > 0 {
		report.PhotoUrl = req.PhotoUrl
	}
	if req.Comment != "" {
		report.Comment = req.Comment
	}
	if req.LatAtSubmission != 0 {
		report.LatAtSubmission = req.LatAtSubmission
	}
	if req.LongAtSubmission != 0 {
		report.LongAtSubmission = req.LongAtSubmission
	}

	data, err := h.service.Report().Update(ctx.Request.Context(), report)
	if err != nil {
		libs.HandleInternalServerError(ctx.Writer, err)
		return
	}

	libs.WriteJSONWithSuccess(ctx.Writer, m.ParseReportRepoToApi(data))
}

func (h *handlerV1) DeleteReport(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.service.Report().Delete(ctx.Request.Context(), id); err != nil {
		libs.HandleInternalServerError(ctx.Writer, err)
		return
	}

	libs.WriteJSONWithSuccess(ctx.Writer, gin.H{"message": "report deleted successfully"})
}

func (h *handlerV1) GetReportById(ctx *gin.Context) {
	id := ctx.Param("id")

	data, err := h.service.Report().GetByID(ctx.Request.Context(), id)
	if err != nil {
		libs.HandleNotFoundError(ctx.Writer, err, "report not found")
		return
	}

	libs.WriteJSONWithSuccess(ctx.Writer, m.ParseReportRepoToApi(data))
}

func (h *handlerV1) GetReportsByInfrastructureId(ctx *gin.Context) {
	infrastructureId := ctx.Param("id")

	limit := ctx.DefaultQuery("limit", "10")
	page := ctx.DefaultQuery("page", "1")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		libs.HandleBadRequestErrWithMessage(ctx.Writer, err, "invalid limit parameter")
		return
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		libs.HandleBadRequestErrWithMessage(ctx.Writer, err, "invalid page parameter")
		return
	}

	req := repo.GetReportsByInfrastructureReq{
		Limit: int32(limitInt),
		Page:  int32(pageInt),
	}

	data, err := h.service.Report().GetByInfrastructureID(ctx.Request.Context(), infrastructureId, req)
	if err != nil {
		libs.HandleInternalServerError(ctx.Writer, err)
		return
	}

	libs.WriteJSONWithSuccess(ctx.Writer, m.ToReportListRepoToApi(data))
}
