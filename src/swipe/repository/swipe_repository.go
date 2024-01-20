package repository

import (
	"dating-app/model/database"
	"dating-app/src/swipe"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SwipeRepository struct {
	db *gorm.DB
}

func NewSwipeRepository(db *gorm.DB) swipe.SwipeRepositoryInterface {
	return &SwipeRepository{db}
}

func (r *SwipeRepository) UpsertSwipeMatches(swipeMatches database.SwipeMatches) error {
	tx := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"is_first_user_like", "is_second_user_like", "updated_at"}),
	}).Create(&swipeMatches)

	return tx.Error
}

func (r *SwipeRepository) GetSwipeMatches(firstUserId int, secondUserId int) (database.SwipeMatches, error) {
	var swipeMatches database.SwipeMatches
	tx := r.db.Where("first_user_id = ? and second_user_id = ?", firstUserId, secondUserId).First(&swipeMatches)

	return swipeMatches, tx.Error
}
