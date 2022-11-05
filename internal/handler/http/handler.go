package handler

import (
	"github.com/Zhiyenbek/users-auth-service/config"
	"github.com/Zhiyenbek/users-auth-service/internal/service"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service *service.Service
	cfg     *config.Configs
}

type Handler interface {
	InitRoutes() *gin.Engine
}

func New(services *service.Service, cfg *config.Configs) Handler {
	return &handler{
		service: services,
		cfg:     cfg,
	}
}

func (h *handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.POST("/sign-in", h.SignIn)
	patient := router.Group("/patient")

	patient.POST("", h.UpdatePatient)
	patient.GET("", h.GetPatient)
	patient.POST("/sign_up", h.RegisterPatient)

	doctor := router.Group("/doctor")
	doctor.POST("/sign-up", h.RegisterDoctor)
	doctor.PUT("", h.UpdateDoctor)
	doctor.GET("", h.GetDoctor)

	return router
}
