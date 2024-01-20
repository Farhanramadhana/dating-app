package auth

import (
	"dating-app/model/dto"
)

type UsecaseInterface interface {
	SignUp(dto.Signup) error
	Signin(dto.Signin) (string, error)
}
