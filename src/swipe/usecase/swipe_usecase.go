package usecase

import (
	"context"
	"dating-app/app"
	"dating-app/model/constant"
	"dating-app/model/database"
	"dating-app/src/swipe"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type SwipeUsecase struct {
	swipeRepository swipe.SwipeRepositoryInterface
	redis           app.Redis
}

func NewSwipeUsecase(
	swipeRepository swipe.SwipeRepositoryInterface,
	redis app.Redis,
) swipe.SwipeUsecaseInterface {
	return &SwipeUsecase{swipeRepository, redis}
}

func (u *SwipeUsecase) Swipe(userID int) error {
	// add showed profile into redis

	// query to database user_profile, except user_id not in redis, randomize the user

	return nil
}

func (u *SwipeUsecase) CountUserSwipe(userID int) int {
	countStr := u.redis.RedisGet(context.Background(), fmt.Sprintf(constant.USER_SWIPE_KEY, userID))
	count, _ := strconv.Atoi(countStr)

	return count
}

func (u *SwipeUsecase) AddUserSwipe(userID int) error {
	countStr := u.redis.RedisGet(context.Background(), fmt.Sprintf(constant.USER_SWIPE_KEY, userID))
	count, _ := strconv.Atoi(countStr)
	count++

	err := u.redis.RedisSet(context.Background(), fmt.Sprintf(constant.USER_SWIPE_KEY, userID), count)
	return err
}

func (u *SwipeUsecase) AddProfileAppeared(userID int, otherUserID int) error {
	profilesStr := u.redis.RedisGet(context.Background(), fmt.Sprintf(constant.USER_PROFILE_APPEARED, userID))

	var profileIDs []string
	if profilesStr != "" {
		profileIDs = strings.Split(profilesStr, ",")
	}

	profileIDs = append(profileIDs, strconv.Itoa(otherUserID))

	value := strings.Join(profileIDs, ",")

	err := u.redis.RedisSet(context.Background(), fmt.Sprintf(constant.USER_PROFILE_APPEARED, userID), value)

	return err
}

func (u *SwipeUsecase) GetProfileAppeared(userID int) []int {
	var profileIDs []int
	profilesStr := u.redis.RedisGet(context.Background(), fmt.Sprintf(constant.USER_PROFILE_APPEARED, userID))

	if profilesStr != "" {
		profiles := strings.Split(profilesStr, ",")
		for _, v := range profiles {
			profileID, _ := strconv.Atoi(v)
			profileIDs = append(profileIDs, profileID)
		}
	}

	return profileIDs
}

func (u *SwipeUsecase) UpsertSwipeMatches(firstUserId int, secondUserId int, firstUserLike *bool, secondUserLike *bool, swipeID int) error {
	data := database.SwipeMatches{
		FirstUserID:      firstUserId,
		SecondUserID:     secondUserId,
		IsFirstUserLike:  firstUserLike,
		IsSecondUserLike: secondUserLike,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	if swipeID != 0 {
		data.ID = swipeID
	}
	return u.swipeRepository.UpsertSwipeMatches(data)
}

func (u *SwipeUsecase) GetSwipeMatches(firstUserId int, secondUserId int) (database.SwipeMatches, error) {
	return u.swipeRepository.GetSwipeMatches(firstUserId, secondUserId)
}

func (u *SwipeUsecase) GetAsFirstUserLikeProfiles(userID int) ([]database.SwipeMatches, error) {
	return u.swipeRepository.GetAsFirstUserLikeProfiles(userID)
}

func (u *SwipeUsecase) GetAsSecondUserLikeProfiles(userID int) ([]database.SwipeMatches, error) {
	return u.swipeRepository.GetAsSecondUserLikeProfiles(userID)
}
