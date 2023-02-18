package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kotaroyamazaki/go-clean-arch-sample-with-standard/internal/usecases"
)

type UserHandler struct {
	userUsecase usecases.UserUsecase
}

func NewUserHandler(userUC usecases.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUC,
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var input struct {
		BitrhDate       time.Time `json:"birth_date"`
		NickeName       string    `json:"nickname"`
		FavoriteBookIDs []int     `json:"favorite_book_ids,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, fmt.Sprintf(`{"status": "%s"}`, err), http.StatusInternalServerError)
		return
	}

	err := h.userUsecase.Register(ctx, input.BitrhDate, input.NickeName, input.FavoriteBookIDs)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf(`{"status": "%s"}`, err), GetStatusCode(err))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"status": "created"}`))
}
