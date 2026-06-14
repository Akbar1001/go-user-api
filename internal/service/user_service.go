package service

import (
	"time"

	"github.com/go-playground/validator/v10"
	"go-user-api/internal/logger"
	"go-user-api/internal/models"
	"go-user-api/internal/repository"
	"go.uber.org/zap"
)

type UserService struct {
	repo      *repository.UserRepository
	validator *validator.Validate
}

func NewUserService(
	repo *repository.UserRepository,
) *UserService {

	return &UserService{
		repo:      repo,
		validator: validator.New(),
	}
}

func CalculateAge(dob time.Time) int {

	now := time.Now()

	age := now.Year() - dob.Year()

	if now.YearDay() < dob.YearDay() {
		age--
	}

	return age
}

func (s *UserService) GetUser(
	id int32,
) (*models.UserResponse, error) {

	user, err := s.repo.GetUser(id)

	if err != nil {
		return nil, err
	}

	response := &models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
		Age:  CalculateAge(user.Dob),
	}

	return response, nil
}

func (s *UserService) ListUsers() ([]models.UserResponse, error) {

	users, err := s.repo.ListUsers()

	if err != nil {
		return nil, err
	}

	var response []models.UserResponse

	for _, user := range users {

		response = append(response, models.UserResponse{
			ID:   user.ID,
			Name: user.Name,
			DOB:  user.Dob.Format("2006-01-02"),
			Age:  CalculateAge(user.Dob),
		})
	}

	return response, nil
}

func (s *UserService) CreateUser(
	req models.CreateUserRequest,
) (*models.UserResponse, error) {

	err := s.validator.Struct(req)

	if err != nil {
		return nil, err
	}

	dob, err := time.Parse(
		"2006-01-02",
		req.DOB,
	)

	if err != nil {
		return nil, err
	}

	user, err := s.repo.CreateUser(
		req.Name,
		dob,
	)

	if err != nil {
		return nil, err
	}

	logger.Log.Info(
	"user created",
	zap.Int32("id", user.ID),
	zap.String("name", user.Name),
	)

	return &models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
		Age:  CalculateAge(user.Dob),
	}, nil


}


func (s *UserService) UpdateUser(
	id int32,
	req models.UpdateUserRequest,
) (*models.UserResponse, error) {

	err := s.validator.Struct(req)

	if err != nil {
		return nil, err
	}

	dob, err := time.Parse(
		"2006-01-02",
		req.DOB,
	)

	if err != nil {
		return nil, err
	}

	user, err := s.repo.UpdateUser(
		id,
		req.Name,
		dob,
	)

	if err != nil {
		return nil, err
	}

	return &models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
		Age:  CalculateAge(user.Dob),
	}, nil
}


func (s *UserService) DeleteUser(
	id int32,
) error {

	logger.Log.Info(
	"user deleted",
	zap.Int32("id", id),
)

	return s.repo.DeleteUser(id)
}