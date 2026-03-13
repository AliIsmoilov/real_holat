package api

import (
	v1 "real-holat/api/v1"
	"real-holat/config"
	"real-holat/internal/service"
	"real-holat/storage"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Cfg     *config.Config
	Service service.ServiceI
	Strg    storage.StorageI
	Enf     *casbin.Enforcer
}

func New(h *Handler) *gin.Engine {
	engine := gin.Default()

	// --- Add CORS Middleware Here ---
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // For production, replace "*" with your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	handlerV1 := v1.New(&v1.HandleV1{
		Cfg:     h.Cfg,
		Service: h.Service,
	})

	api := engine.Group("/api")
	apiV1 := api.Group("/v1")

	apiV1.Use(
	// middleware.AuthMiddleware(),
	// middleware.CasbinMiddleware(h.Enf),
	)

	// infrastructure-types routes
	apiV1.POST("/infrastructure-types", handlerV1.CreateInfrastructureType)
	apiV1.PUT("/infrastructure-types/:id", handlerV1.UpdateInfrastructureType)
	apiV1.GET("/infrastructure-types/:id", handlerV1.GetInfrastructureTypeById)
	apiV1.GET("/infrastructure-types", handlerV1.GetListInfrastructureTypes)
	apiV1.DELETE("/infrastructure-types/:id", handlerV1.DeleteInfrastructureType)

	// auth routes
	apiV1.POST("/users/login", handlerV1.Login)
	apiV1.POST("/users/login-with-tg-otp", handlerV1.LoginWithTgOtp)

	// user routes
	apiV1.POST("/users", handlerV1.CreateUser)
	apiV1.GET("/users/:id", handlerV1.GetUserByID)
	apiV1.GET("/users", handlerV1.GetUsers)
	apiV1.PUT("/users/:id", handlerV1.UpdateUser)
	apiV1.DELETE("/users/:id", handlerV1.DeleteUser)

	// image route
	apiV1.POST("/image/upload", handlerV1.UploadImage)

	return engine
}
