package service

import (
	"time"

	"github.com/go-playground/validator/v10"

	"go-user-api/internal/models"
	"go-user-api/internal/repository"
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
) error {

	err := s.validator.Struct(req)

	if err != nil {
		return err
	}

	dob, err := time.Parse(
		"2006-01-02",
		req.DOB,
	)

	if err != nil {
		return err
	}

	_, err = s.repo.CreateUser(
		req.Name,
		dob,
	)

	return err
}

func (s *UserService) UpdateUser(
	id int32,
	req models.UpdateUserRequest,
) error {

	err := s.validator.Struct(req)

	if err != nil {
		return err
	}

	dob, err := time.Parse(
		"2006-01-02",
		req.DOB,
	)

	if err != nil {
		return err
	}

	_, err = s.repo.UpdateUser(
		id,
		req.Name,
		dob,
	)

	return err
}

func (s *UserService) DeleteUser(
	id int32,
) error {

	return s.repo.DeleteUser(id)
}