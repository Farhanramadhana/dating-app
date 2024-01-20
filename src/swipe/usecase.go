package swipe

import "dating-app/model/database"

type UsecaseInterface interface {
	Swipe(userID int) error
	CountUserSwipe(userID int) int
	AddUserSwipe(userID int) error
	AddProfileAppeared(userID int, otherUserID int) error
	GetProfileAppeared(userID int) []int

	UpsertSwipeMatches(firstUserID int, secondUserID int, firstUserLike *bool, secondUserLike *bool, swipeID int) error
	GetSwipeMatches(firstUserID int, secondUserID int) (database.SwipeMatches, error)
	GetAsFirstUserLikeProfiles(userID int) ([]database.SwipeMatches, error)
	GetAsSecondUserLikeProfiles(userID int) ([]database.SwipeMatches, error)
}
