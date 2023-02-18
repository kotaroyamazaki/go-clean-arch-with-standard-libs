package main

import (
	"log"
	"net/http"

	"github.com/kotaroyamazaki/go-clean-arch-sample-with-standard/internal/handlers"
	"github.com/kotaroyamazaki/go-clean-arch-sample-with-standard/internal/infra/repository"
	"github.com/kotaroyamazaki/go-clean-arch-sample-with-standard/internal/middleware"
	"github.com/kotaroyamazaki/go-clean-arch-sample-with-standard/internal/usecases"
)

func main() {
	userRepo := repository.NewUserRepository()
	userFavoriteBookRepo := repository.NewUserFavoriteBooksRepository()
	userUC := usecases.NewUserUsecase(userRepo, userFavoriteBookRepo)
	userH := handlers.NewUserHandler(userUC)

	http.HandleFunc("/users", middleware.Cors(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("todo: implement get method"))
		case http.MethodPost:
			userH.Register(w, r)
		default:
			http.Error(w, `{"status":"permits only GET or POST"}`, http.StatusMethodNotAllowed)
		}
	}))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
