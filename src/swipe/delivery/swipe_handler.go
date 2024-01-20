package delivery

import (
	"dating-app/app/middleware"
	"dating-app/app/utils"
	"dating-app/model/constant"
	"dating-app/model/dto"

	"dating-app/src/swipe"
	"dating-app/src/user"
	"net/http"

	"github.com/gorilla/mux"
)

type SwipeHandler struct {
	swipeUsecase swipe.SwipeUsecaseInterface
	userUsecase  user.UserUsecaseInterface
}

func NewSwipeHandler(r *mux.Router, swipeUsecase swipe.SwipeUsecaseInterface, userUsecase user.UserUsecaseInterface) {
	handler := SwipeHandler{swipeUsecase, userUsecase}

	r.Handle("/show", middleware.AuthenticateMiddleware(http.HandlerFunc(handler.Show))).Methods("GET")
	r.Handle("/swipe", middleware.AuthenticateMiddleware(http.HandlerFunc(handler.Swipe))).Methods("GET")
}

func (h *SwipeHandler) Show(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)

	userProfile, err := h.userUsecase.GetUserProfileByUserID(userID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	limit := 10
	// check user is premium or not, need to call user service
	if !userProfile.IsPremiumUser {
		// if !premium, user swipe limited to 10
		count := h.swipeUsecase.CountUserSwipe(userID)
		if count == constant.MAX_SWIPE {
			utils.RespondWithError(w, http.StatusInternalServerError, "maximum swipe")
			return
		}

		limit = constant.MAX_SWIPE - count
	}

	// load other user profiles up to 10, and set to redis
	userProfilesID := h.swipeUsecase.GetProfileAppeared(userID)

	userProfiles, _ := h.userUsecase.GetUserProfilesNotIn(userProfilesID, limit)
	var response []dto.UserProfile
	for _, v := range userProfiles {
		profile := dto.UserProfile{
			UserID:           v.UserID,
			Gender:           v.Gender,
			Birthdate:        v.Birthdate.String(),
			GenderPreference: v.GenderPreference,
			IsPremiumUser:    v.IsPremiumUser,
		}

		response = append(response, profile)
	}

	utils.RespondWithJSON(w, http.StatusOK, userProfiles)
}

func (h *SwipeHandler) Swipe(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)

	// get profile appeared from redis
	// profileIDs := h.swipeUsecase.GetProfileAppeared(userID)

	// load userProfiles without previous appeared profiles

	// add appeared profiles into redis
	h.swipeUsecase.AddProfileAppeared(userID, 0)
}
