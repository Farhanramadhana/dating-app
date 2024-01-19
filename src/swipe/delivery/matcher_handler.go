package delivery

import (
	matcher "matcher-service/internal"
	"net/http"

	"github.com/gorilla/mux"
)

type MatcherHandler struct {
	matcherUsecase matcher.MatcherUsecaseInterface
}

func NewMatcherHandler(r *mux.Router, matcherUsecase matcher.MatcherUsecaseInterface) {
	handler := MatcherHandler{matcherUsecase}

	// Attach authentication middleware to all routes
	r.Use(AuthenticateMiddleware)

	r.HandleFunc("/swipe/", handler.Swipe).Methods("GET")
}

func (h *MatcherHandler) Swipe(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)
	h.matcherUsecase.Swipe(userID)
}
