package auth

import "dating-app/model/database"

type AuthRepositoryInterface interface {
	CreateUser(user database.User) (database.User, error)
	CreateUserPassword(user database.UserPassword) error
	FindByEmail(email string) (database.User, error)
	FindUserPasswordByUserID(userID int) (database.UserPassword, error)
}
