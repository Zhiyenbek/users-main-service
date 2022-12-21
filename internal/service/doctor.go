package service

import (
	"time"

	"github.com/Zhiyenbek/users-main-service/config"
	"github.com/Zhiyenbek/users-main-service/internal/models"
	"github.com/Zhiyenbek/users-main-service/internal/repository"
	"go.uber.org/zap"
)

type doctorService struct {
	doctorRepo repository.DoctorRepository
	logger     *zap.SugaredLogger
	cfg        *config.Configs
}

func NewDoctorService(repo *repository.Repository, logger *zap.SugaredLogger, cfg *config.Configs) DoctorService {
	return &doctorService{
		doctorRepo: repo.DoctorRepository,
		logger:     logger,
		cfg:        cfg,
	}
}

func (s *doctorService) UpdateDoctor(doctorReq *models.UpdateDoctorRequest) error {
	err := s.doctorRepo.UpdateDoctor(doctorReq)
	if err != nil {
		s.logger.Error(err)
		return err
	}
	return nil
}
func (s *doctorService) CreateDoctor(doctorReq *models.CreateDoctorRequest) (*models.CreateDoctorResponse, error) {

	res, err := s.doctorRepo.CreateDoctor(doctorReq)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return res, nil
}
func (s *doctorService) DeleteDoctor(ID int64) error {
	err := s.doctorRepo.DeleteDoctor(ID)
	if err != nil {
		s.logger.Error(err)
		return err
	}
	return nil
}
func (s *doctorService) GetDoctor(ID int64) (*models.GetDoctorResponse, error) {
	res, err := s.doctorRepo.GetDoctor(ID)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return res, nil
}

func (s *doctorService) SearchDoctors(searchArgs *models.Search) (*models.SearchDoctorsResponse, error) {
	res, err := s.doctorRepo.SearchDoctors(searchArgs)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return res, nil
}

func (s *doctorService) GetDoctorByDepartment(ID int64, search *models.Search) (*models.SearchDoctorsResponse, error) {
	res, err := s.doctorRepo.SearchDoctorsByDepartment(search, ID)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return res, nil
}

func (s *doctorService) GetDepartments() (*models.GetDepartments, error) {
	res, err := s.doctorRepo.GetDepartments()
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return res, nil
}

func (s *doctorService) CreateAppointment(req *models.CreateAppointmentRequest) (*models.CreateAppointmentResponse, error) {
	res, err := s.doctorRepo.CreateAppointment(req)
	if err != nil {
		s.logger.Error(err)
		return &models.CreateAppointmentResponse{
			Error: err.Error(),
		}, err
	}
	return res, nil
}

func (s *doctorService) GetAppointmentsByDate(bookArgs *models.Appointment) (*models.GetAppointmentsResponse, error) {
	res, err := s.doctorRepo.GetBookedAppointmentsByDate(bookArgs)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	// busines logic
	timeIt, _ := time.Parse("15:04", "08:00")
	endTime, _ := time.Parse("15:04", "18:00")
	var emptySlots []string
	for timeIt != endTime {
		timeIt = timeIt.Add(time.Hour)
		timeItFormatted := timeIt.Format("15:04")
		if res[timeItFormatted] {
			continue
		}
		emptySlots = append(emptySlots, timeItFormatted)
		// log.Println(timeItFormatted)
	}
	return &models.GetAppointmentsResponse{
		EmptySlots: emptySlots,
	}, nil
}
