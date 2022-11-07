package handler

import (
	"errors"
	"log"
	"strconv"

	"github.com/Zhiyenbek/users-main-service/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h *handler) RegisterDoctor(c *gin.Context) {
	req := &models.CreateDoctorRequest{}
	if err := c.ShouldBindWith(req, binding.JSON); err != nil {
		log.Printf("ERROR: invalid input, some fields are incorrect: %s\n", err.Error())
		c.AbortWithStatusJSON(400, sendResponse(-1, nil, models.ErrInvalidInput))
		return
	}
	resp, err := h.service.DoctorService.CreateDoctor(req)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrDoctorNotFound):
			c.AbortWithStatusJSON(404, sendResponse(-1, nil, models.ErrInvalidInput))
			return
		default:
			c.AbortWithStatusJSON(500, sendResponse(-1, nil, models.ErrInternalServer))
			return
		}
	}
	c.JSON(201, sendResponse(0, resp, nil))
}

func (h *handler) UpdateDoctor(c *gin.Context) {
	req := &models.UpdateDoctorRequest{}
	if err := c.ShouldBindWith(req, binding.JSON); err != nil {
		log.Printf("ERROR: invalid input, some fields are incorrect: %s\n", err.Error())
		c.AbortWithStatusJSON(400, sendResponse(-1, nil, models.ErrInvalidInput))
		return
	}
	idParam := c.Param("doctor_id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		log.Printf("ERROR: invalid input, missing user id: \n")
		c.AbortWithStatusJSON(400, sendResponse(-1, nil, models.ErrInvalidInput))
	}
	req.ID = int64(id)
	err = h.service.DoctorService.UpdateDoctor(req)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrDoctorNotFound):
			c.AbortWithStatusJSON(404, sendResponse(-1, nil, models.ErrInvalidInput))
			return
		default:
			c.AbortWithStatusJSON(500, sendResponse(-1, nil, models.ErrInternalServer))
			return
		}
	}
	c.JSON(200, sendResponse(0, req, nil))
}

func (h *handler) DeleteDoctor(c *gin.Context) {
	idParam := c.Param("doctor_id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		log.Printf("ERROR: invalid input, missing user id: \n")
		c.AbortWithStatusJSON(400, sendResponse(-1, nil, models.ErrInvalidInput))
	}

	err = h.service.DoctorService.DeleteDoctor(int64(id))
	if err != nil {
		switch {
		case errors.Is(err, models.ErrDoctorNotFound):
			c.AbortWithStatusJSON(404, sendResponse(-1, nil, models.ErrInvalidInput))
			return
		default:
			c.AbortWithStatusJSON(500, sendResponse(-1, nil, models.ErrInternalServer))
			return
		}
	}
}

func (h *handler) GetDoctor(c *gin.Context) {
	idParam := c.Param("doctor_id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		log.Printf("ERROR: invalid input, missing user id: \n")
		c.AbortWithStatusJSON(400, sendResponse(-1, nil, models.ErrInvalidInput))
	}

	res, err := h.service.DoctorService.GetDoctor(int64(id))
	if err != nil {
		switch {
		case errors.Is(err, models.ErrDoctorNotFound):
			c.AbortWithStatusJSON(404, sendResponse(-1, nil, models.ErrInvalidInput))
			return
		default:
			c.AbortWithStatusJSON(500, sendResponse(-1, nil, models.ErrInternalServer))
			return
		}
	}
	c.JSON(200, sendResponse(0, res, nil))
}

func (h *handler) GetAllDoctors(c *gin.Context) {
	res, err := h.service.DoctorService.GetAllDoctors()
	if err != nil {
		switch {
		case errors.Is(err, models.ErrDoctorNotFound):
			c.AbortWithStatusJSON(404, sendResponse(-1, nil, models.ErrInvalidInput))
			return
		default:
			c.AbortWithStatusJSON(500, sendResponse(-1, nil, models.ErrInternalServer))
			return
		}
	}
	c.JSON(200, sendResponse(0, res, nil))
}
