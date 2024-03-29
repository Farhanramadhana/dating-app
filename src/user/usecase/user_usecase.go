package usecase

import (
	"dating-app/model/constant"
	"dating-app/model/database"
	"dating-app/model/dto"
	"dating-app/src/user"
	"fmt"
	"time"
)

type UserUsecase struct {
	userRepository user.RepositoryInterface
}

func NewUserUsecase(
	userRepository user.RepositoryInterface,
) user.UsecaseInterface {
	return &UserUsecase{userRepository}
}

func (u *UserUsecase) UpsertUserProfile(request dto.UserProfile) error {
	fmt.Println(request.Birthdate)
	birthdate, _ := time.Parse(constant.YYYYMMDD, request.Birthdate)

	userProfile := database.UserProfile{
		UserID:           request.UserID,
		Gender:           request.Gender,
		Birthdate:        birthdate,
		GenderPreference: request.GenderPreference,
		IsPremiumUser:    request.IsPremiumUser,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	err := u.userRepository.UpsertUserProfile(userProfile)
	return err
}

func (u *UserUsecase) AddUserImage(request dto.UserImage) error {
	userImage := database.UserImage{
		UserID:    request.UserID,
		ImageURL:  request.ImageURL,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := u.userRepository.AddUserImage(userImage)
	return err
}

func (u *UserUsecase) GetUserProfileByUserID(userID int) (database.UserProfile, error) {
	return u.userRepository.GetUserProfileByUserID(userID)
}

func (u *UserUsecase) GetUserProfilesNotIn(userIDs []int, limit int) ([]database.UserProfile, error) {
	return u.userRepository.GetUserProfilesNotIn(userIDs, limit)
}
