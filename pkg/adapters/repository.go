package adapters

import (
	"cloud-run-playground/pkg/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db}
}

func (bs *userRepository) GetUsersByName(limit int, name string) ([]domain.User, error) {
	if limit == 0 {
		limit = 50
	}
	users := []domain.User{}

	err := bs.db.Limit(limit).Where("first_name ILIKE ?", name+"%").Or("last_name ILIKE ?", name+"%").Or("username ILIKE ?", name+"%").Find(&users).Error

	return users, err
}

func (bs *userRepository) Store(user domain.User) (domain.User, error) {
	err := bs.db.Create(&user).Error

	return user, err
}

func (bs *userRepository) Find(userId int) (domain.User, error) {
	user := domain.User{}
	err := bs.db.Find(&user, userId).Error

	return user, err
}

func (bs *userRepository) DeleteUser(userId int) error {
	return bs.db.Delete(domain.User{}, userId).Error
}

func (bs *userRepository) GetUserByUsername(username string) (domain.User, error) {
	user := domain.User{}
	err := bs.db.Where("username = ?", username).Find(&user).Error

	return user, err
}

func (bs *userRepository) GetUserByEmail(email string) (domain.User, error) {
	user := domain.User{}
	err := bs.db.Where("email = ?", email).Find(&user).Error

	return user, err
}
