package handler

import (
	"geant4help/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	rtr := gin.New()

	index := rtr.Group("/")
	{

	}
}
