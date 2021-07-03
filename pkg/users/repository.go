package users

import "gorm.io/gorm"

type UserStorage struct {
	db *gorm.DB
}

type UserRepository interface {
	GetAllUsers(limit int) ([]User, error)
	CreateUser(user User) (User, error)
	GetUserById(userId int) (User, error)
}

func NewUsersRepository(db *gorm.DB) UserRepository {
	return &UserStorage{db}
}

func (bs *UserStorage) GetAllUsers(limit int) ([]User, error) {
	if limit == 0 {
		limit = 50
	}
	users := []User{}

	err := bs.db.Limit(limit).Find(&users).Error

	return users, err
}

func (bs *UserStorage) CreateUser(user User) (User, error) {
	err := bs.db.Create(&user).Error

	return user, err
}

func (bs *UserStorage) GetUserById(userId int) (User, error) {
	user := User{}
	err := bs.db.Find(&user, userId).Error

	return user, err
}
