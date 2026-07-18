package services

import "github.com/Yer01/workout_tracker/internal/models"

type AuthService interface {
	GetSingle(string) (int, error)
	CreateUser(string, string, string) (int, error)
}

type authService struct {
	user_repo *models.UserRepo
}

func NewAuthService(ur *models.UserRepo) AuthService {
	return &authService{
		user_repo: ur,
	}
}

func (as *authService) GetSingle(email string) (int, error) {
	id, err := as.user_repo.Select(email)

	if err == nil && id != 0 {
		return id, nil
	}

	return id, err
}

func (as *authService) CreateUser(email string, password string, name string) (int, error) {
	id, err := as.user_repo.Insert(email, password, name)

	if id == 0 || err != nil {
		return id, err
	}

	return id, nil
}
