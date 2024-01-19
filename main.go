package main

import (
	"dating-app/app"
	"dating-app/app/utils"
	"dating-app/model/constant"
	authHandler "dating-app/src/auth/delivery"
	authRepository "dating-app/src/auth/repository"
	authUsecase "dating-app/src/auth/usecase"
	userHandler "dating-app/src/user/delivery"
	userRepository "dating-app/src/user/repository"
	userUsecase "dating-app/src/user/usecase"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("read the config .env file")
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	db, err := app.InitDatabase()
	if err != nil {
		errMessage := fmt.Sprintf("error running database: %s ", err)
		panic(errMessage)
	}
	r := mux.NewRouter()

	// repository
	authRepository := authRepository.NewAuthRepository(db)
	userRepository := userRepository.NewUserRepository(db)

	// usecase
	authUsecase := authUsecase.NewAuthUsecase(authRepository)
	userUsecase := userUsecase.NewUserUsecase(userRepository)

	// handler
	authHandler.NewAuthHandler(r, authUsecase)

	// Example protected route
	r.Handle("/health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "hello world"})
	})).Methods("GET")

	userHandler.NewUserHandler(r, userUsecase)
	
	fmt.Printf("starting the server %s on port %s \n", time.Now().Format(constant.YYYY_MM_DD_HH_MM_SS), os.Getenv("port"))
	
	http.Handle("/", r)
	addr := fmt.Sprintf(":%s", os.Getenv("port"))
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		errMessage := fmt.Sprintf("server error %s ", err)
		panic(errMessage)
	}
}
