package service

import (
	"time"

	"go-user-api/internal/models"
	"go-user-api/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(
	repo *repository.UserRepository,
) *UserService {
	return &UserService{
		repo: repo,
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