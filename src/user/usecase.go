package user

import (
	"dating-app/model/database"
	"dating-app/model/dto"
)

type UserUsecaseInterface interface {
	UpsertUserProfile(dto.UserProfile) error
	AddUserImage(dto.UserImage) error
	GetUserProfileByUserID(userID int) (database.UserProfile, error)
	GetUserProfilesNotIn(userIDs []int, limit int) ([]database.UserProfile, error)
}
