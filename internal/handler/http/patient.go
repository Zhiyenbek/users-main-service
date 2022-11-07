package handler

import (
	"errors"
	"log"
	"strconv"

	"github.com/Zhiyenbek/users-main-service/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h *handler) RegisterPatient(c *gin.Context) {
	req := &models.CreatePatientRequest{}
	if err := c.ShouldBindWith(req, binding.JSON); err != nil {
		log.Printf("ERROR: invalid input, some fields are incorrect: %s\n", err.Error())
		c.AbortWithStatusJSON(400, sendResponse(-1, nil, models.ErrInvalidInput))
		return
	}
	resp, err := h.service.PatientService.CreatePatient(req)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrPatientNotFound):
			c.AbortWithStatusJSON(404, sendResponse(-1, nil, models.ErrInvalidInput))
			return
		default:
			c.AbortWithStatusJSON(500, sendResponse(-1, nil, models.ErrInternalServer))
			return
		}
	}
	c.JSON(201, sendResponse(0, resp, nil))
}

func (h *handler) UpdatePatient(c *gin.Context) {
	idParam := c.Param("patient_id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		log.Printf("ERROR: invalid input, missing user id: \n")
		c.AbortWithStatusJSON(400, sendResponse(-1, nil, models.ErrInvalidInput))
	}
	req := &models.UpdatePatientRequest{}
	if err := c.ShouldBindWith(req, binding.JSON); err != nil {
		log.Printf("ERROR: invalid input, some fields are incorrect: %s\n", err.Error())
		c.AbortWithStatusJSON(400, sendResponse(-1, nil, models.ErrInvalidInput))
		return
	}
	req.ID = int64(id)
	err = h.service.PatientService.UpdatePatient(req)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrPatientNotFound):
			c.AbortWithStatusJSON(404, sendResponse(-1, nil, models.ErrInvalidInput))
			return
		default:
			c.AbortWithStatusJSON(500, sendResponse(-1, nil, models.ErrInternalServer))
			return
		}
	}
	c.JSON(200, sendResponse(0, req, nil))
}

func (h *handler) DeletePatient(c *gin.Context) {
	idParam := c.Param("patient_id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		log.Printf("ERROR: invalid input, missing user id: \n")
		c.AbortWithStatusJSON(400, sendResponse(-1, nil, models.ErrInvalidInput))
	}

	err = h.service.PatientService.DeletePatient(int64(id))
	if err != nil {
		switch {
		case errors.Is(err, models.ErrPatientNotFound):
			c.AbortWithStatusJSON(404, sendResponse(-1, nil, models.ErrInvalidInput))
			return
		default:
			c.AbortWithStatusJSON(500, sendResponse(-1, nil, models.ErrInternalServer))
			return
		}
	}
}
func (h *handler) GetPatient(c *gin.Context) {
	idParam := c.Param("patient_id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		log.Printf("ERROR: invalid input, missing user id: \n")
		c.AbortWithStatusJSON(400, sendResponse(-1, nil, models.ErrInvalidInput))
	}

	res, err := h.service.PatientService.GetPatient(int64(id))
	if err != nil {
		switch {
		case errors.Is(err, models.ErrPatientNotFound):
			c.AbortWithStatusJSON(404, sendResponse(-1, nil, models.ErrInvalidInput))
			return
		default:
			c.AbortWithStatusJSON(500, sendResponse(-1, nil, models.ErrInternalServer))
			return
		}
	}
	c.JSON(200, sendResponse(0, res, nil))
}
