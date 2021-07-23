package usersSearch

import "cloud-run-playground/pkg/domain"

type UserService interface {
	SearchByName(limit int, name string) ([]domain.User, error)
	GetById(id int) (domain.User, error)
}

type userService struct {
	usersRepo domain.UserRepository
}

func NewService(usersRepo domain.UserRepository) UserService {
	return &userService{usersRepo}
}

func (s *userService) SearchByName(limit int, name string) ([]domain.User, error) {
	return s.usersRepo.GetUsersByName(limit, name)
}

func (s *userService) GetById(id int) (domain.User, error) {
	return s.usersRepo.Find(id)
}
