package dto

type UserProfile struct {
	UserID           int    `json:"user_id,omitempty"`
	Gender           string `json:"gender"`
	Birthdate        string `json:"birthdate"`
	GenderPreference string `json:"gender_preference"`
	IsPremiumUser    bool   `json:"is_premium_user"`
}

type UserImage struct {
	UserID   int    `json:"user_id,omitempty"`
	ImageURL string `json:"image_url"`
}
