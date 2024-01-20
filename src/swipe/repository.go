package swipe

import "dating-app/model/database"

type RepositoryInterface interface {
	UpsertSwipeMatches(swipeMatches database.SwipeMatches) error
	GetSwipeMatches(firstUserID int, secondUserID int) (database.SwipeMatches, error)
	GetAsFirstUserLikeProfiles(userID int) ([]database.SwipeMatches, error)
	GetAsSecondUserLikeProfiles(userID int) ([]database.SwipeMatches, error)
}
