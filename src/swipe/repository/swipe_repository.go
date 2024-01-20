package repository

import (
	"dating-app/src/swipe"

	"gorm.io/gorm"
)

type SwipeRepository struct {
	db *gorm.DB
}

func NewSwipeRepository(db *gorm.DB) swipe.SwipeRepositoryInterface {
	return &SwipeRepository{db}
}
