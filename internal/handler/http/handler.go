package handler

import (
	"github.com/Zhiyenbek/users-main-service/config"
	"github.com/Zhiyenbek/users-main-service/internal/service"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"go.uber.org/zap"
)

type handler struct {
	service *service.Service
	cfg     *config.Configs
	logger  *zap.SugaredLogger
}

type Handler interface {
	InitRoutes() *gin.Engine
}

func New(services *service.Service, logger *zap.SugaredLogger, cfg *config.Configs) Handler {
	return &handler{
		service: services,
		cfg:     cfg,
		logger:  logger,
	}
}

func (h *handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(cors.AllowAll())
	router.POST("/sign-in", h.SignIn)
	router.POST("/refresh-token", h.RefreshToken)

	patient := router.Group("/patients")

	patient.PUT("/:patient_id", h.UpdatePatient)
	patient.GET("/:patient_id", h.VerifyToken, h.GetPatient)
	patient.POST("/sign-up", h.VerifyToken, h.RegisterPatient)
	patient.GET("/all", h.GetAllPatients)

	doctor := router.Group("/doctors")
	doctor.POST("/sign-up", h.RegisterDoctor)
	doctor.PUT("/:doctor_id", h.UpdateDoctor)
	doctor.GET("/:doctor_id", h.GetDoctor)
	doctor.DELETE("/:doctor_id", h.DeleteDoctor)
	doctor.GET("", h.SearchDoctor)
	doctor.GET("/departments/:department_id", h.GetDoctorByDepartment)

	doctor.GET("/departments", h.GetDepartments)

	doctor.POST("/appointments", h.CreateAppointment)
	doctor.GET("/appointments", h.GetAppointmentsByDate)
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
