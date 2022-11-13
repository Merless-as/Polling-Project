package handler

import (
	"RestApiProj/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/createPoll/", h.createPoll)
		api.POST("/poll/", h.Poll)
		api.POST("/getResult/", h.getResult)
	}

	return router
}
