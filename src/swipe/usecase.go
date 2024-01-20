package swipe

type SwipeUsecaseInterface interface {
	Swipe(userID int) error
	CountUserSwipe(userID int) int
	AddUserSwipe(userID int) error
	AddProfileAppeared(userID int, otherUserID int) error
	GetProfileAppeared(userID int) []int
}
