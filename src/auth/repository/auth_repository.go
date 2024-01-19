package repository

import (
	"dating-app/model/database"
	"dating-app/src/auth"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) auth.AuthRepositoryInterface {
	return &AuthRepository{db}
}

func (r *AuthRepository) CreateUser(user database.User) (database.User, error) {
	tx := r.db.Create(&user)
	return user, tx.Error
}

func (r *AuthRepository) CreateUserPassword(user database.UserPassword) error {
	tx := r.db.Create(&user)
	return tx.Error
}

func (r *AuthRepository) FindByEmail(email string) (database.User, error) {
	var user database.User
	tx := r.db.Where("email = ?", email).First(&user)

	return user, tx.Error
}

func (r *AuthRepository) FindUserPasswordByUserID(userID int) (database.UserPassword, error) {
	var userPassword database.UserPassword
	tx := r.db.Where("user_id = ?", userID).First(&userPassword)

	return userPassword, tx.Error
}
