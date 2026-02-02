package routes

import (
	v1 "rmn-backend/services/user-service/internal/app/routes/v1"
	"rmn-backend/services/user-service/internal/common"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	*gin.Engine
}

func (r *Router) NewRouter(ctx *common.App) error {
	// Set mode
	gin.SetMode(ctx.Cfg.HttpServer.Mode)

	r.Engine = gin.New()

	return nil
}

func (r *Router) Run(addr string) {
	r.Engine.Run(addr)
}

func (r *Router) SetupRouter(ctx *common.App) {
	// Swagger route should be outside of any group to be accessible
	r.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health Check Route
	r.Engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	v1Group := r.Engine.Group("/api/v1")
	{
		v1.SetupV1APIRoutes(v1Group, ctx)
	}
}
