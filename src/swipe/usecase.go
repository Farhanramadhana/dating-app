package swipe

import "dating-app/model/database"

type SwipeUsecaseInterface interface {
	Swipe(userID int) error
	CountUserSwipe(userID int) int
	AddUserSwipe(userID int) error
	AddProfileAppeared(userID int, otherUserID int) error
	GetProfileAppeared(userID int) []int

	UpsertSwipeMatches(firstUserId int, secondUserId int, firstUserLike *bool, secondUserLike *bool, swipeID int) error
	GetSwipeMatches(firstUserId int, secondUserId int) (database.SwipeMatches, error)
}
