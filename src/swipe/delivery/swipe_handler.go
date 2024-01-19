package delivery

import (
	"dating-app/app/middleware"
	"dating-app/src/swipe"
	"net/http"

	"github.com/gorilla/mux"
)

type SwipeHandler struct {
	swipeUsecase swipe.SwipeUsecaseInterface
}

func NewSwipeHandler(r *mux.Router, swipeUsecase swipe.SwipeUsecaseInterface) {
	handler := SwipeHandler{swipeUsecase}

	r.Handle("/swipe/", middleware.AuthenticateMiddleware(http.HandlerFunc(handler.Swipe))).Methods("GET")
}

func (h *SwipeHandler) Swipe(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)
	h.swipeUsecase.Swipe(userID)
}
