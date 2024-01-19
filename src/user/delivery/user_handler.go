package delivery

import (
	"dating-app/app/utils"
	"dating-app/model/dto"
	user "dating-app/src/user"
	"encoding/json"
	"net/http"

	"dating-app/app/middleware"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	userUsecase user.UserUsecaseInterface
}

func NewUserHandler(r *mux.Router, userUsecase user.UserUsecaseInterface) {
	handler := UserHandler{userUsecase}

	r.Handle("/user/profile", middleware.AuthenticateMiddleware(http.HandlerFunc(handler.UpsertUserProfile))).Methods("POST")
	r.Handle("/user/image", middleware.AuthenticateMiddleware(http.HandlerFunc(handler.AddUserImage))).Methods("POST")
}

func (h *UserHandler) UpsertUserProfile(w http.ResponseWriter, r *http.Request) {
	var userDto dto.UserProfile
	if err := json.NewDecoder(r.Body).Decode(&userDto); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	userID := r.Context().Value("user_id").(string)
	userDto.UserID = userID
	err := h.userUsecase.UpsertUserProfile(userDto)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "success", "message": "success update profile"})
}

func (h *UserHandler) AddUserImage(w http.ResponseWriter, r *http.Request) {
	var userDto dto.UserImage
	if err := json.NewDecoder(r.Body).Decode(&userDto); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	userID := r.Context().Value("user_id").(string)
	userDto.UserID = userID
	err := h.userUsecase.AddUserImage(userDto)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "success", "message": "success add user image"})
}
