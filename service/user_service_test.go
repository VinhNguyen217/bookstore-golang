package service

import (
	"book-store/dto/request"
	mockRepo "book-store/mock/book-store/repository"
	"book-store/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"testing"
)

func TestCreateUser(t *testing.T) {
	mockUserRepo := &mockRepo.MockUserRepository{}

	mockUserRepo.
		On("FindByUserName", mock.Anything, "user-existed").
		Return(&model.User{Username: "user-existed", Password: "123456789"}, nil)
	mockUserRepo.
		On("FindByUserName", mock.Anything, "user-not-existed").
		Return(nil, gorm.ErrRecordNotFound)

	userService := &userServiceImpl{
		userRepo: mockUserRepo,
	}

	// user existed
	_, err := userService.CreateUser(&request.UserRequest{Username: "user-existed"})
	if err == nil {
		t.Fail()
	}

	// password < 8
	_, err = userService.CreateUser(&request.UserRequest{Username: "user-not-existed", Password: "1234567"})
	if err == nil {
		t.Fail()
	}

	// valid
	resp, err := userService.CreateUser(&request.UserRequest{Username: "user-not-existed", Password: "123456789"})
	if err != nil {
		t.Fail()
	}
	if resp == nil {
		t.Fail()
	}
	assert.Equal(t, "user-not-existed", resp.Username)
}
