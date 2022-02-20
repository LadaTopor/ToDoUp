package handler

import (
	"github.com/LadaTopor/ToDoUp/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sing-in", h.signUp)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.signUp)
			lists.GET("/", h.signUp)
			lists.GET("/:id", h.signUp)
			lists.PUT("/:id", h.signUp)
			lists.DELETE("/:id", h.signUp)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.signUp)
				items.GET("/", h.signUp)
				items.GET("/:item_id", h.signUp)
				items.PUT("/:item_id", h.signUp)
				items.DELETE("/:item_id", h.signUp)
			}
		}
	}

	return router
}
