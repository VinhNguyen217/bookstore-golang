package dto

import (
	"github.com/golang-jwt/jwt/v5"
)

type LoginResponse struct {
	Meta *Meta  `json:"meta"`
	Data *Token `json:"data"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type JwtClaims struct {
	jwt.RegisteredClaims
	UserId   int    `json:"userId"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	Role     string `json:"role"`
}

type PasswordLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TwoFaResponseData struct {
	Secret string `json:"secret"`
}

type SetupTwoFaRequest struct {
	OTP string `json:"otp"`
}
