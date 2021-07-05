package usersSearch

type UserRepository interface {
	GetUsersByName(limit int, name string) ([]User, error)
}

type UserService interface {
	SearchByName(limit int, name string) ([]User, error)
}

type service struct {
	repo UserRepository
}

func NewService(repo UserRepository) UserService {
	return &service{repo}
}

func (s *service) SearchByName(limit int, name string) ([]User, error) {
	users, err := s.repo.GetUsersByName(limit, name)

	return users, err
}
