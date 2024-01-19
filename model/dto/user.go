package dto

type UserProfile struct {
	UserID           string `json:"-`
	Gender           string `json:"gender"`
	Birthdate        string `json:"birthdate"`
	GenderPreference string `json:"gender_preference"`
	IsPremiumUser    bool   `json:"is_premium_user"`
}

type UserImage struct {
	UserID   string `json:"-`
	ImageURL string `json:"image_url"`
}
