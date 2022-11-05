package handler

import (
	"github.com/Zhiyenbek/users-auth-service/config"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
	cfg     *config.Configs
}

func NewHandler(services *service.Service, cfg *config.Configs) *Handler {
	return &Handler{
		service: services,
		cfg:     cfg,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	patient := router.Group("/patient")

	patient.POST("/sign_up")

	admin := router.Group("/admin")

	admin.POST("/sign-in")

	doctor := router.Group("/doctor")
	doctor.POST("/sign-up")
	doctor.PUT("/edit")

	return router
}
