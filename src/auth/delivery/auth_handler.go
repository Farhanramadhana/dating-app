package handler

import (
	"dating-app/app/utils"
	"dating-app/model/dto"
	"dating-app/src/auth"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type AuthHandler struct {
	authUsecase auth.UsecaseInterface
}

func NewAuthHandler(r *mux.Router, authUsecase auth.UsecaseInterface) {
	handler := AuthHandler{authUsecase}
	r.HandleFunc("/signup", handler.SignUpHandler).Methods("POST")
	r.HandleFunc("/signin", handler.SignInHandler).Methods("POST")
}

func (h *AuthHandler) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var authDto dto.Signup
	if err := json.NewDecoder(r.Body).Decode(&authDto); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err := h.authUsecase.SignUp(authDto)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "success", "message": "created user successful"})
}

func (h *AuthHandler) SignInHandler(w http.ResponseWriter, r *http.Request) {
	var authDto dto.Signin
	if err := json.NewDecoder(r.Body).Decode(&authDto); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	jwt, err := h.authUsecase.Signin(authDto)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "success", "jwt": jwt})
}
