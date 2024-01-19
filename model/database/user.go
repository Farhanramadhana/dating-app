package database

import "time"

type UserProfile struct {
	ID               int
	UserID           string
	Gender           string
	Birthdate        time.Time
	GenderPreference string
	IsPremiumUser    bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type UserImage struct {
	ID        int
	UserID    string
	ImageURL  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
