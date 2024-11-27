package service

import (
	"book-store/dto/request"
	"book-store/enum"
	"book-store/model"
	"book-store/repository"
	"book-store/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"strings"
)

const (
	saltLength = 20
)

type UserService interface {
	CreateUser(user *request.UserRequest) (*model.User, error)
	Update(ctx *gin.Context, user *request.UserRequest) (*model.User, error)
	FindByID(id int) (*model.User, error)
	FindAll() ([]model.User, error)
	GetMyInfo(ctx *gin.Context) (*model.User, error)
}

type userServiceImpl struct {
	userRepo repository.UserRepository
}

func newUserService(di *do.Injector) (UserService, error) {
	return &userServiceImpl{
		userRepo: do.MustInvoke[repository.UserRepository](di),
	}, nil
}

func (s *userServiceImpl) CreateUser(req *request.UserRequest) (*model.User, error) {
	if strings.TrimSpace(req.Username) == "" {
		return nil, errors.New("username is required")
	}
	passwordReq := req.Password
	if strings.TrimSpace(passwordReq) == "" {
		return nil, errors.New("password is required")
	}
	if len(passwordReq) < 8 {
		return nil, errors.New("password contains at least 8 characters")
	}
	existedUser, _ := s.userRepo.FindByUserName(req.Username)
	if existedUser != nil && existedUser.Username != "" {
		return nil, errors.New("account is exist")
	}

	salt := utils.RandomStringWithLength(saltLength)         // Generate salt
	hashedPassword := utils.HashPassword(req.Password, salt) // Hash password

	user := &model.User{
		Name:     req.Name,
		Username: req.Username,
		Password: hashedPassword,
		Salt:     salt,
		Role:     enum.USER,
	}
	return s.userRepo.Create(user)
}

func (s *userServiceImpl) Update(ctx *gin.Context, req *request.UserRequest) (*model.User, error) {
	userId := ctx.GetInt("user_id")
	existedUser, _ := s.userRepo.FindById(userId)
	// validate
	nameReq := req.Name
	if strings.TrimSpace(nameReq) != "" {
		existedUser.Name = nameReq
	}
	usernameReq := req.Username
	if strings.TrimSpace(usernameReq) != "" {
		existedUser.Username = usernameReq
	}
	passwordReq := req.Password
	if strings.TrimSpace(passwordReq) != "" {
		if len(passwordReq) < 8 {
			return nil, errors.New("password contains at least 8 characters")
		}
		salt := utils.RandomStringWithLength(saltLength)
		hashedPassword := utils.HashPassword(req.Password, salt)
		existedUser.Salt = salt
		existedUser.Password = hashedPassword
	}
	err := s.userRepo.Update(existedUser)
	if err != nil {
		return nil, err
	} else {
		return s.userRepo.FindById(userId)
	}
}

func (s *userServiceImpl) FindByID(id int) (*model.User, error) {
	return s.userRepo.FindById(id)
}

func (s *userServiceImpl) FindAll() ([]model.User, error) {
	return s.userRepo.FindAll()
}

func (s *userServiceImpl) GetMyInfo(ctx *gin.Context) (*model.User, error) {
	userId := ctx.GetInt("user_id")
	return s.userRepo.FindById(userId)
}
