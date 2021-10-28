package handler

import (
	"github.com/fuadsuleyman/go-auth1/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)
type Handler struct{
	services *service.Service 
}

// this is dependency injection
func NewHandler(services *service.Service) *Handler{
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine{
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(cors.Default())

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createItem)
		}		
	}

	return router
}