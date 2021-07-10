package usersSearch

type UserRepository interface {
	GetUsersByName(limit int, name string) ([]User, error)
	GetUserById(id int) (User, error)
}

type UserService interface {
	SearchByName(limit int, name string) ([]User, error)
	GetById(id int) (User, error)
}

type userService struct {
	usersRepo UserRepository
}

func NewService(usersRepo UserRepository) UserService {
	return &userService{usersRepo}
}

func (s *userService) SearchByName(limit int, name string) ([]User, error) {
	return s.usersRepo.GetUsersByName(limit, name)
}

func (s *userService) GetById(id int) (User, error) {
	return s.usersRepo.GetUserById(id)
}
