package repository

import (
	mactcher "matcher-service/internal"

	"gorm.io/gorm"
)

type MatcherRepository struct {
	db *gorm.DB
}

func NewMatcherRepository(db *gorm.DB) mactcher.MatcherRepositoryInterface {
	return &MatcherRepository{db}
}
