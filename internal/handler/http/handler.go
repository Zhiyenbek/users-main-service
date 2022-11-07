package handler

import (
	"github.com/Zhiyenbek/users-main-service/config"
	"github.com/Zhiyenbek/users-main-service/internal/service"
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
	router.POST("/refresh-token", h.RefreshToken)

	patient := router.Group("/patients", h.VerifyToken)

	patient.PUT("/:patient_id", h.UpdatePatient)
	patient.GET("/:patient_id", h.VerifyToken, h.GetPatient)
	patient.POST("/sign_up", h.VerifyToken, h.RegisterPatient)

	doctor := router.Group("/doctors", h.VerifyToken)
	doctor.POST("/sign-up", h.RegisterDoctor)
	doctor.PUT("/:doctor_id", h.UpdateDoctor)
	doctor.GET("/:doctor_id", h.GetDoctor)
	doctor.DELETE("/:doctor_id", h.DeleteDoctor)

	return router
}

func sendResponse(status int, data interface{}, err error) gin.H {
	var errResponse gin.H
	if err != nil {
		errResponse = gin.H{
			"message": err.Error(),
		}
	} else {
		errResponse = nil
	}

	return gin.H{
		"data":   data,
		"status": status,
		"error":  errResponse,
	}
}
