package user

import "dating-app/model/database"

type UserRepositoryInterface interface {
	UpsertUserProfile(database.UserProfile) error
	AddUserImage(database.UserImage) error
	GetUserProfileByUserID(userID int) (database.UserProfile, error)
	GetUserProfilesNotIn(userIDs []int, limit int) ([]database.UserProfile, error)
}
