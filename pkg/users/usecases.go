package users

type Service interface {
	GetAllUsers(limit int) ([]User, error)
}

type service struct {
	repo UserRepository
}

func (s *service) GetAllUsers(limit int) ([]User, error) {
	users, err := s.repo.GetAllUsers(limit)

	return users, err
}

func NewService(repo UserRepository) Service {
	return &service{repo}
}
