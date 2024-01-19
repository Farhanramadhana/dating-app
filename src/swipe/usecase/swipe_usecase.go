package usecase

import (
	"dating-app/src/swipe"
)

type SwipeUsecase struct {
	swipeRepository swipe.SwipeRepositoryInterface
}

func NewSwipeUsecase(
	swipeRepository swipe.SwipeRepositoryInterface,
) swipe.SwipeUsecaseInterface {
	return &SwipeUsecase{swipeRepository}
}

func (u *SwipeUsecase) Swipe(userID int) error {
	// check user is premium or not, need to call user service

	// if !premium, count user swipe, integrate counter to redis

	// add showed profile into redis

	// query to database user_profile, except user_id not in redis, randomize the user

	return nil
}
