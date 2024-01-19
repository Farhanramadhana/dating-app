package repository

import (
	"dating-app/app"
	"dating-app/src/swipe"

	"gorm.io/gorm"
)

type SwipeRepository struct {
	db    *gorm.DB
	redis app.Redis
}

func NewSwipeRepository(db *gorm.DB, redis app.Redis) swipe.SwipeRepositoryInterface {
	return &SwipeRepository{db, redis}
}
