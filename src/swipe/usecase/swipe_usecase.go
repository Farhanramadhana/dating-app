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
	swipeRepository swipe.RepositoryInterface
	redis           app.Redis
}

func NewSwipeUsecase(
	swipeRepository swipe.RepositoryInterface,
	redis app.Redis,
) swipe.UsecaseInterface {
	return &SwipeUsecase{swipeRepository, redis}
}

func (u *SwipeUsecase) CountUserSwipe(userID int) int {
	today := time.Now().Format(constant.YYYYMMDD)
	countStr := u.redis.RedisGet(context.Background(), fmt.Sprintf(constant.UserSwipeKey, userID, today))
	count, _ := strconv.Atoi(countStr)

	return count
}

func (u *SwipeUsecase) AddUserSwipe(userID int) error {
	today := time.Now().Format(constant.YYYYMMDD)
	countStr := u.redis.RedisGet(context.Background(), fmt.Sprintf(constant.UserSwipeKey, userID, today))
	count, _ := strconv.Atoi(countStr)
	count++

	err := u.redis.RedisSet(context.Background(), fmt.Sprintf(constant.UserSwipeKey, userID, today), count)
	return err
}

func (u *SwipeUsecase) AddProfileAppeared(userID int, otherUserID int) error {
	today := time.Now().Format(constant.YYYYMMDD)
	profilesStr := u.redis.RedisGet(context.Background(), fmt.Sprintf(constant.UserProfileAppeared, userID, today))

	var profileIDs []string
	if profilesStr != "" {
		profileIDs = strings.Split(profilesStr, ",")
	}

	profileIDs = append(profileIDs, strconv.Itoa(otherUserID))

	value := strings.Join(profileIDs, ",")

	err := u.redis.RedisSet(context.Background(), fmt.Sprintf(constant.UserProfileAppeared, userID, today), value)

	return err
}

func (u *SwipeUsecase) GetProfileAppeared(userID int) []int {
	var profileIDs []int
	today := time.Now().Format(constant.YYYYMMDD)
	profilesStr := u.redis.RedisGet(context.Background(), fmt.Sprintf(constant.UserProfileAppeared, userID, today))

	if profilesStr != "" {
		profiles := strings.Split(profilesStr, ",")
		for _, v := range profiles {
			profileID, _ := strconv.Atoi(v)
			profileIDs = append(profileIDs, profileID)
		}
	}

	return profileIDs
}

func (u *SwipeUsecase) UpsertSwipeMatches(firstUserID int, secondUserID int, firstUserLike *bool, secondUserLike *bool, swipeID int) error {
	data := database.SwipeMatches{
		FirstUserID:      firstUserID,
		SecondUserID:     secondUserID,
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

func (u *SwipeUsecase) GetSwipeMatches(firstUserID int, secondUserID int) (database.SwipeMatches, error) {
	return u.swipeRepository.GetSwipeMatches(firstUserID, secondUserID)
}

func (u *SwipeUsecase) GetAsFirstUserLikeProfiles(userID int) ([]database.SwipeMatches, error) {
	return u.swipeRepository.GetAsFirstUserLikeProfiles(userID)
}

func (u *SwipeUsecase) GetAsSecondUserLikeProfiles(userID int) ([]database.SwipeMatches, error) {
	return u.swipeRepository.GetAsSecondUserLikeProfiles(userID)
}
