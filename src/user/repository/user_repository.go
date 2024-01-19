package repository

import (
	"dating-app/model/database"
	"dating-app/src/user"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.UserRepositoryInterface {
	return &UserRepository{db}
}

func (r *UserRepository) UpsertUserProfile(userProfile database.UserProfile) error {
	tx := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"gender", "birthdate", "gender_preference", "is_premium_user", "updated_at"}),
	}).Create(&userProfile)

	return tx.Error
}

func (r *UserRepository) AddUserImage(userImage database.UserImage) error {
	tx := r.db.Create(&userImage)
	return tx.Error
}

func (r *UserRepository) GetUserProfileByUserID(userID int) (database.UserProfile, error) {
	var userProfile database.UserProfile
	tx := r.db.Where("user_id = ?", userID).Find(&userProfile)
	return userProfile, tx.Error
}