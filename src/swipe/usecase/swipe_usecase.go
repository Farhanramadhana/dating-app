package usecase

import (
	"context"
	"dating-app/app"
	"dating-app/model/constant"
	"dating-app/src/swipe"
	"fmt"
	"strconv"
	"strings"
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
	err := u.redis.RedisAppend(context.Background(), fmt.Sprintf(constant.USER_PROFILE_APPEARED, userID), strconv.Itoa(otherUserID))

	return err
}

func (u *SwipeUsecase) GetProfileAppeared(userID int) []int {
	var profileIDs []int
	profilesStr := u.redis.RedisGet(context.Background(), fmt.Sprintf(constant.USER_PROFILE_APPEARED, userID))

	profiles := strings.Split(profilesStr, " ")
	for _, v := range profiles {
		profileID, _ := strconv.Atoi(v)
		profileIDs = append(profileIDs, profileID)
	}

	return profileIDs
}
