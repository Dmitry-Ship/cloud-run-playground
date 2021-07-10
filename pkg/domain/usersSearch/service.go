package usersSearch

type UserRepository interface {
	GetUsersByName(limit int, name string) ([]User, error)
}

type UserService interface {
	SearchByName(limit int, name string) ([]User, error)
}

type userService struct {
	usersRepo UserRepository
}

func NewService(usersRepo UserRepository) UserService {
	return &userService{usersRepo}
}

func (s *userService) SearchByName(limit int, name string) ([]User, error) {
	users, err := s.usersRepo.GetUsersByName(limit, name)

	return users, err
}
