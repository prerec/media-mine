package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/prerec/media-mine/docs"
	"github.com/prerec/media-mine/internal/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct{}

// NewHandler возвращает экземпляр структуры Handler
func NewHandler() *Handler {
	return &Handler{}
}

// InitRoutes конфигурирует роутер и в зависимости от настроек cfg может включить middleware logger библиотеки logrus
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	middleware.CheckCfg(router)

	api := router.Group("/api")
	{
		exchange := api.Group("/exchange")
		{
			exchange.POST("/", h.exchange)
		}
	}

	return router
}
