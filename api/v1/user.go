package v1

import (
	"context"
	"fmt"
	"real-holat/api/models"
	"real-holat/pkg/jwt"
	"real-holat/pkg/libs"
	"real-holat/pkg/mapper/api"
	"real-holat/storage/repo"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (h *handlerV1) Login(ctx *gin.Context) {
	var req models.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		libs.HandleBadRequestErr(ctx.Writer, err)
		return
	}

	user, err := h.service.User().GetByPhone(ctx.Request.Context(), req.Phone)
	if err != nil {
		libs.HandleUnauthorizedErr(ctx.Writer, fmt.Errorf("user not found or invalid credentials"))
		return
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	); err != nil {
		libs.HandleUnauthorizedErr(ctx.Writer, fmt.Errorf("invalid credentials"))
		return
	}

	token, err := jwt.GenerateJWT(user)
	if err != nil {
		libs.HandleInternalServerError(ctx.Writer, err)
		return
	}

	// ctx.JSON(http.StatusOK, gin.H{"access_token": token})
	libs.WriteJSONWithSuccess(ctx.Writer, api.ParseLoginWithPhoneToResponse(token, user))
}

func (h *handlerV1) LoginWithTgOtp(c *gin.Context) {
	var req models.VerifyOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		libs.HandleBadRequestErr(c.Writer, err)
		return
	}

	verification, err := h.service.Verification().GetByCode(context.Background(), req.Code, time.Now())
	if err != nil {
		libs.HandleInternalServerError(c.Writer, err)
		return
	}
	if verification == nil {
		libs.HandleBadRequestErr(c.Writer, fmt.Errorf("invalid or expired code"))
		return
	}

	ok, err := h.service.Verification().VerifyByCode(context.Background(), req.Code)
	if err != nil {
		libs.HandleInternalServerError(c.Writer, err)
		return
	}
	if !ok {
		libs.HandleBadRequestErr(c.Writer, fmt.Errorf("invalid or expired code"))
		return
	}

	// Create or update user from verification data
	user, err := h.service.User().CreateOrUpdateFromVerification(context.Background(), verification)
	if err != nil {
		libs.HandleInternalServerError(c.Writer, err)
		return
	}

	// Generate JWT token
	token, err := jwt.GenerateJWT(user)
	if err != nil {
		libs.HandleInternalServerError(c.Writer, err)
		return
	}
	response := api.ParseLoginWithTgOtpToResponse(token, user, verification)

	libs.WriteJSONWithSuccess(c.Writer, response)
}

func (h *handlerV1) CreateUser(c *gin.Context) {
	var req models.UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		libs.HandleBadRequestErr(c.Writer, err)
		return
	}

	user := api.ParseUserCreateRequestToUser(req)
	createdUser, err := h.service.User().Create(c.Request.Context(), &user)
	if err != nil {
		libs.HandleInternalServerError(c.Writer, err)
		return
	}

	response := api.ParseUserToResponse(createdUser)
	libs.WriteJSONWithSuccess(c.Writer, response)
}

func (h *handlerV1) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		libs.HandleBadRequestErr(c.Writer, fmt.Errorf("invalid user ID"))
		return
	}

	user, err := h.service.User().GetByID(c.Request.Context(), id)
	if err != nil {
		libs.HandleNotFoundErr(c.Writer, fmt.Errorf("user not found"))
		return
	}

	libs.WriteJSONWithSuccess(c.Writer, api.ParseUserToResponse(user))
}

func (h *handlerV1) GetUsers(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	pageStr := c.DefaultQuery("page", "1")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	users, err := h.service.User().GetAll(c.Request.Context(), repo.GetAllUsersReq{
		Limit: int32(limit),
		Page:  int32(page),
	})
	if err != nil {
		libs.HandleInternalServerError(c.Writer, err)
		return
	}

	libs.WriteJSONWithSuccess(c.Writer, api.ParseUsersToListResponse(users))
}

func (h *handlerV1) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		libs.HandleBadRequestErr(c.Writer, fmt.Errorf("invalid user ID"))
		return
	}

	var req models.UserUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		libs.HandleBadRequestErr(c.Writer, err)
		return
	}

	existingUser, err := h.service.User().GetByID(c.Request.Context(), id)
	if err != nil {
		libs.HandleNotFoundErr(c.Writer, fmt.Errorf("user not found"))
		return
	}

	updatedUser := api.ParseUserUpdateRequestToUser(req, existingUser)
	result, err := h.service.User().Update(c.Request.Context(), &updatedUser)
	if err != nil {
		libs.HandleInternalServerError(c.Writer, err)
		return
	}

	response := api.ParseUserToResponse(result)
	libs.WriteJSONWithSuccess(c.Writer, response)
}

func (h *handlerV1) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		libs.HandleBadRequestErr(c.Writer, fmt.Errorf("invalid user ID"))
		return
	}

	err = h.service.User().Delete(c.Request.Context(), id)
	if err != nil {
		libs.HandleInternalServerError(c.Writer, err)
		return
	}

	libs.WriteJSONWithSuccess(c.Writer, "user deleted successfully")
}
