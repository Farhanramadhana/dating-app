package delivery

import (
	"dating-app/app/middleware"
	"dating-app/app/utils"
	"dating-app/model/constant"
	"dating-app/model/dto"
	"errors"
	"strconv"

	"dating-app/src/swipe"
	"dating-app/src/user"
	"net/http"

	"log"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type SwipeHandler struct {
	swipeUsecase swipe.SwipeUsecaseInterface
	userUsecase  user.UserUsecaseInterface
}

func NewSwipeHandler(r *mux.Router, swipeUsecase swipe.SwipeUsecaseInterface, userUsecase user.UserUsecaseInterface) {
	handler := SwipeHandler{swipeUsecase, userUsecase}

	r.Handle("/show", middleware.AuthenticateMiddleware(http.HandlerFunc(handler.Show))).Methods("GET")
	r.Handle("/swipe", middleware.AuthenticateMiddleware(http.HandlerFunc(handler.Swipe))).Methods("POST")

}

func (h *SwipeHandler) Show(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.Context().Value("user_id").(string)
	userID, _ := strconv.Atoi(userIdStr)

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
		if count >= constant.MAX_SWIPE {
			utils.RespondWithError(w, http.StatusInternalServerError, "maximum swipe")
			return
		}

		limit = constant.MAX_SWIPE - count
	}

	// load other user profiles up to 10 except visited and like profile
	userProfilesID := h.swipeUsecase.GetProfileAppeared(userID)
	userProfilesID = append(userProfilesID, userID)

	firstLikeProfiles, _ := h.swipeUsecase.GetAsFirstUserLikeProfiles(userID)
	if len(firstLikeProfiles) > 0 {
		for _, v := range firstLikeProfiles {
			userProfilesID = append(userProfilesID, v.SecondUserID)
		}
	}

	secondLikeProfiles, _ := h.swipeUsecase.GetAsSecondUserLikeProfiles(userID)
	if len(secondLikeProfiles) > 0 {
		for _, v := range secondLikeProfiles {
			userProfilesID = append(userProfilesID, v.FirstUserID)
		}
	}

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

	utils.RespondWithJSON(w, http.StatusOK, response)
}

func (h *SwipeHandler) Swipe(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.Context().Value("user_id").(string)
	userID, _ := strconv.Atoi(userIdStr)

	userProfile, err := h.userUsecase.GetUserProfileByUserID(userID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// get swipe action
	swipeAction := r.URL.Query().Get("action")
	otherUserID, _ := strconv.Atoi(r.URL.Query().Get("other_user_id"))

	// add counter if user !premium
	if !userProfile.IsPremiumUser {
		count := h.swipeUsecase.CountUserSwipe(userID)
		if count >= constant.MAX_SWIPE {
			utils.RespondWithError(w, http.StatusInternalServerError, "maximum swipe")
			return
		}

		err := h.swipeUsecase.AddUserSwipe(userID)
		log.Printf("user %v error add counter swipe for user %v: %s \n", userID, otherUserID, err)
	}

	// add appeared profiles into redis
	err = h.swipeUsecase.AddProfileAppeared(userID, otherUserID)
	if err != nil {
		log.Printf("user %v error add profile appeared for user %v: %s \n", userID, otherUserID, err)
	}

	// get from database matcher as second person
	swipe, err := h.swipeUsecase.GetSwipeMatches(otherUserID, userID)

	var isLike *bool
	if swipeAction == constant.SWIPE_LIKE {
		like := true
		isLike = &like
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	} else if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		// insert new data, swipe == skip not save to database, profile will appear on the next time
		if swipeAction == constant.SWIPE_LIKE {
			err := h.swipeUsecase.UpsertSwipeMatches(userID, otherUserID, isLike, nil, 0)
			log.Printf("error save swipe matches %v %s %v: %s \n", userID, swipeAction, otherUserID, err)
		}
	} else {
		// data swipe exist, update
		err := h.swipeUsecase.UpsertSwipeMatches(swipe.FirstUserID, swipe.SecondUserID, swipe.IsFirstUserLike, isLike, swipe.ID)
		log.Printf("error update swipe matches %v %s %v: %s", userID, swipeAction, otherUserID, err)
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "success"})
}
