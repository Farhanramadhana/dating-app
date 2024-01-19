package auth

import (
	"dating-app/model/dto"
)

type AuthUsecaseInterface interface {
	SignUp(dto.Signup) error
	Signin(dto.Signin) (string, error)
}
