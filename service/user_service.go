package service

import (
	"book-store/dto/request"
	"book-store/enum"
	"book-store/model"
	"book-store/repository"
	"book-store/utils"
	"errors"
	"github.com/samber/do"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

const (
	saltLength = 20
)

type UserService interface {
	Create(user *request.UserRequest) (*model.User, error)
	Update(user *request.UserRequest, id int) (*model.User, error)
	Delete(id int) error
	FindByID(id int) (*model.User, error)
	FindAll() ([]model.User, error)
}

type userServiceImpl struct {
	userRepo repository.UserRepository
}

func newUserService(di *do.Injector) (UserService, error) {
	return &userServiceImpl{
		userRepo: do.MustInvoke[repository.UserRepository](di),
	}, nil
}

func (s *userServiceImpl) Create(req *request.UserRequest) (*model.User, error) {
	existedUser, _ := s.userRepo.FindByUserName(req.Username)
	if existedUser != nil && existedUser.Name != "" {
		return nil, errors.New("Tài khoản đã tồn tại")
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

func (s *userServiceImpl) Update(req *request.UserRequest, id int) (*model.User, error) {
	_, errId := s.userRepo.FindById(id)
	if errId != nil {
		return nil, errors.New("Người dùng không tồn tại")
	}
	hashedPassword, _ := encodePassword(req.Password)
	user := &model.User{
		ID:       id,
		Name:     req.Name,
		Username: req.Username,
		Password: hashedPassword,
	}
	err := s.userRepo.Update(user)
	if err != nil {
		return nil, err
	} else {
		return s.userRepo.FindById(id)
	}
}

func (s *userServiceImpl) Delete(id int) error {
	_, err := s.userRepo.FindById(id)
	if err != nil {
		return errors.New("Người dùng không tồn tại")
	}
	return s.userRepo.Delete(id)
}

func (s *userServiceImpl) FindByID(id int) (*model.User, error) {
	return s.userRepo.FindById(id)
}

func (s *userServiceImpl) FindAll() ([]model.User, error) {
	return s.userRepo.FindAll()
}

func encodePassword(pw string) (string, error) {
	if len(strings.TrimSpace(pw)) == 0 {
		return "", nil
	}
	pwByte := []byte(pw)
	hashedPassword, err := bcrypt.GenerateFromPassword(pwByte, bcrypt.DefaultCost)
	return string(hashedPassword), err
}
