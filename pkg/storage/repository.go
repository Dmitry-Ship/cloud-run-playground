package storage

import (
	"cloud-run-playground/pkg/domain/usersSearch"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) usersSearch.UserRepository {
	return &userRepository{db}
}

func (bs *userRepository) GetUsersByName(limit int, name string) ([]usersSearch.User, error) {
	if limit == 0 {
		limit = 50
	}
	users := []usersSearch.User{}

	err := bs.db.Limit(limit).Where("first_name ILIKE ?", name+"%").Or("last_name ILIKE ?", name+"%").Or("username ILIKE ?", name+"%").Find(&users).Error

	return users, err
}

func (bs *userRepository) CreateUser(user usersSearch.User) (usersSearch.User, error) {
	err := bs.db.Create(&user).Error

	return user, err
}

func (bs *userRepository) GetUserById(userId int) (usersSearch.User, error) {
	user := usersSearch.User{}
	err := bs.db.Find(&user, userId).Error

	return user, err
}
