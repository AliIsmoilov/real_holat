package v1

import (
	"net/http"
	"real-holat/api/models"
	"real-holat/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func (h *handlerV1) Login(ctx *gin.Context) {
	var req models.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.User().GetByEmail(ctx.Request.Context(), req.Email)
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
	if user.Password != req.Password {
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
