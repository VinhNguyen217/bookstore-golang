package service

import (
	"book-store/dto"
	"book-store/log"
	"book-store/repository"
	"book-store/utils"
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/samber/do"
	"net/http"
	"time"
)

type AuthService interface {
	PasswordLogin(ctx context.Context, req *dto.PasswordLoginRequest) (*dto.LoginResponse, error)
}

type authServiceImpl struct {
	userRepo repository.UserRepository
	jwtUtil  utils.JWTUtil
}

func newAuthService(di *do.Injector) (AuthService, error) {
	return &authServiceImpl{
		userRepo: do.MustInvoke[repository.UserRepository](di),
		jwtUtil:  do.MustInvoke[utils.JWTUtil](di),
	}, nil
}

func (s *authServiceImpl) PasswordLogin(ctx context.Context, req *dto.PasswordLoginRequest) (*dto.LoginResponse, error) {
	user, err := s.userRepo.FindByUserName(req.Username)
	if err != nil {
		return nil, err
	}
	hashedPass := utils.HashPassword(req.Password, user.Salt)
	if hashedPass != user.Password {
		return nil, errors.New("invalid password")
	}

	// generate access token
	currentTime := time.Now()
	claims := dto.JwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(currentTime.Add(30 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(currentTime),
		},
		UserName: user.Username,
		Name:     user.Name,
		UserId:   user.ID,
		Role:     user.Role.String(),
	}

	accessToken, err := s.jwtUtil.GenerateToken(&claims)
	if err != nil {
		log.Errorw(ctx, "Error when generating token for user : ", err)
		return nil, err
	}

	return &dto.LoginResponse{
		Meta: &dto.Meta{
			Code:    http.StatusOK,
			Message: "Login success",
		},
		Data: &dto.Token{
			AccessToken: accessToken,
		},
	}, nil
}
