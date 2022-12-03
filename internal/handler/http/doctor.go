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
		return
	}
	req.ID = int64(id)
	err = h.service.DoctorService.UpdateDoctor(req)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrDoctorNotFound):
			c.AbortWithStatusJSON(404, sendResponse(-1, nil, models.ErrDoctorNotFound))
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
		return
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
	c.JSON(200, sendResponse(1, nil, nil))
}

func (h *handler) GetDoctor(c *gin.Context) {
	idParam := c.Param("doctor_id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		log.Printf("ERROR: invalid input, missing user id: \n")
		c.AbortWithStatusJSON(400, sendResponse(-1, nil, models.ErrInvalidInput))
		return
	}

	res, err := h.service.DoctorService.GetDoctor(int64(id))
	if err != nil {
		switch {
		case errors.Is(err, models.ErrDoctorNotFound):
			c.AbortWithStatusJSON(404, sendResponse(-1, nil, models.ErrDoctorNotFound))
			return
		default:
			c.AbortWithStatusJSON(500, sendResponse(-1, nil, models.ErrInternalServer))
			return
		}
	}
	c.JSON(200, sendResponse(0, res, nil))
}

func (h *handler) SearchDoctor(c *gin.Context) {
	search := c.Query("search")

	pageNum, err := strconv.Atoi(c.Query("page_num"))
	if err != nil || pageNum < 1 {
		pageNum = 1
	}
	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}
	searchArgs := &models.Search{
		Search:   search,
		PageNum:  pageNum,
		PageSize: pageSize,
	}

	res, err := h.service.DoctorService.SearchDoctors(searchArgs)
	if err != nil {
		c.AbortWithStatusJSON(500, sendResponse(-1, nil, models.ErrInternalServer))
		return
	}
	c.JSON(200, sendResponse(0, res, nil))
}

func (h *handler) GetDoctorByDepartment(c *gin.Context) {
	idParam := c.Param("department_id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		log.Printf("ERROR: invalid input, missing user id: \n")
		c.AbortWithStatusJSON(400, sendResponse(-1, nil, models.ErrInvalidInput))
		return
	}
	pageNum, err := strconv.Atoi(c.Query("page_num"))
	if err != nil || pageNum < 1 {
		pageNum = 1
	}
	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}
	searchArgs := &models.Search{
		PageNum:  pageNum,
		PageSize: pageSize,
	}
	res, err := h.service.GetDoctorByDepartment(int64(id), searchArgs)
	if err != nil {
		c.AbortWithStatusJSON(500, sendResponse(-1, nil, models.ErrInternalServer))
		return
	}
	c.JSON(200, sendResponse(0, res, nil))
}

func (h *handler) GetDepartments(c *gin.Context) {

	res, err := h.service.GetDepartments()
	if err != nil {
		c.AbortWithStatusJSON(500, sendResponse(-1, nil, models.ErrInternalServer))
		return
	}
	c.JSON(200, sendResponse(0, res, nil))
}

func (h *handler) CreateAppointment(c *gin.Context) {
	req := &models.CreateAppointmentRequest{}
	if err := c.ShouldBindWith(req, binding.JSON); err != nil {
		log.Printf("ERROR: invalid input, some fields are incorrect: %s\n", err.Error())
		c.AbortWithStatusJSON(400, sendResponse(-1, nil, models.ErrInvalidInput))
		return
	}
	resp, err := h.service.DoctorService.CreateAppointment(req)
	if err != nil {
		switch {
		default:
			c.AbortWithStatusJSON(500, sendResponse(-1, resp.Error, models.ErrInvalidInput))
			return
		}
	}
	c.JSON(201, sendResponse(0, resp, nil))
}
