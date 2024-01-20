package usecase

import (
	"dating-app/model/database"
	"dating-app/model/dto"
	"dating-app/src/auth"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	authRepository auth.AuthRepositoryInterface
}

func NewAuthUsecase(
	authRepository auth.AuthRepositoryInterface,
) auth.AuthUsecaseInterface {
	return &AuthUsecase{authRepository}
}

func (u *AuthUsecase) SignUp(request dto.Signup) (err error) {
	user := database.User{
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	userDB, err := u.authRepository.CreateUser(user)
	if err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	// create user password
	userPassword := database.UserPassword{
		UserID:       userDB.ID,
		PasswordHash: string(hash),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err = u.authRepository.CreateUserPassword(userPassword)
	return err
}

func (u *AuthUsecase) Signin(request dto.Signin) (string, error) {
	userDB, err := u.authRepository.FindByEmail(request.Email)
	if err != nil {
		return "", err
	}

	userPassword, err := u.authRepository.FindUserPasswordByUserID(userDB.ID)
	if err != nil {
		return "", err
	}

	// verify password
	byteHash := []byte(userPassword.PasswordHash)
	err = bcrypt.CompareHashAndPassword(byteHash, []byte(request.Password))
	if err != nil {
		return "", errors.New("wrong password")
	}

	// generate JWT
	jwtToken, err := generateToken(userDB.ID)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func generateToken(userID int) (string, error) {
	var jwtSecret = []byte(os.Getenv("jwt_secret_key"))
	claims := jwt.MapClaims{
		"sub": fmt.Sprintf("%v", userID),
		"exp": time.Now().Add(time.Hour * 1).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
