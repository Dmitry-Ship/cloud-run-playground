package domain

type UserRepository interface {
	GetUsersByName(limit int, name string) ([]User, error)
	Find(id int) (User, error)
}
