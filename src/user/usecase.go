package user

import "dating-app/model/dto"

type UserUsecaseInterface interface {
	UpsertUserProfile(dto.UserProfile) error
	AddUserImage(dto.UserImage) error
}
