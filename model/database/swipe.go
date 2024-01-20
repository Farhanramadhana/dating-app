package database

import "time"

type SwipeMatches struct {
	ID               int
	FirstUserID      int
	SecondUserID     int
	IsFirstUserLike  *bool
	IsSecondUserLike *bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
