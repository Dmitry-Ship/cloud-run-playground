package storage

import (
	"cloud-run-playground/pkg/domain/usersSearch"

	"gorm.io/gorm"
)

type UserStorage struct {
	db *gorm.DB
}

func NewUsersStorage(db *gorm.DB) usersSearch.UserRepository {
	return &UserStorage{db}
}

func (bs *UserStorage) GetUsersByName(limit int, name string) ([]usersSearch.User, error) {
	if limit == 0 {
		limit = 50
	}
	users := []usersSearch.User{}

	err := bs.db.Limit(limit).Where("first_name ILIKE ?", name+"%").Or("last_name ILIKE ?", name+"%").Or("username ILIKE ?", name+"%").Find(&users).Error

	return users, err
}

func (bs *UserStorage) CreateUser(user usersSearch.User) (usersSearch.User, error) {
	err := bs.db.Create(&user).Error

	return user, err
}

func (bs *UserStorage) GetUserById(userId int) (usersSearch.User, error) {
	user := usersSearch.User{}
	err := bs.db.Find(&user, userId).Error

	return user, err
}
