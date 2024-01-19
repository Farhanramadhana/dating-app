package user

import "dating-app/model/database"

type UserRepositoryInterface interface {
	UpsertUserProfile(database.UserProfile) error
	AddUserImage(database.UserImage) error
}
