package swipe

import "dating-app/model/database"

type SwipeRepositoryInterface interface {
	UpsertSwipeMatches(swipeMatches database.SwipeMatches) error
	GetSwipeMatches(firstUserId int, secondUserId int) (database.SwipeMatches, error)
}
