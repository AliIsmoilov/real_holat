package v1

import (
	"context"
	"fmt"
	"net/http"
	"real-holat/api/models"
	"real-holat/pkg/jwt"
	"real-holat/pkg/libs"
	"real-holat/pkg/mapper/api"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *handlerV1) Login(ctx *gin.Context) {
	var req models.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.User().GetByPhone(ctx.Request.Context(), req.Phone)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "invalid credentials"})
		return
	}

	// if err := bcrypt.CompareHashAndPassword(
	// 	[]byte(user.PasswordHash),
	// 	[]byte(req.Password),
	// ); err != nil {
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{"message": "invalid credentials"})
	// 	return
	// }
	if user.FullName != req.Password { // Temporary: using full_name as password for testing
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "password does not match"})
		return
	}

	token, err := jwt.GenerateJWT(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"access_token": token})
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
