package repository

import (
	"book-store/model"
	"errors"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) error
	Delete(id int) error
	FindById(id int) (*model.User, error)
	FindAll() ([]model.User, error)
	FindByUserName(username string) (*model.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func newUserRepository(di *do.Injector) (UserRepository, error) {
	db := do.MustInvoke[*gorm.DB](di)
	return &userRepo{db: db}, nil
}

func (r userRepo) Create(user *model.User) (*model.User, error) {
	err := r.db.Create(user).Error
	return user, err
}

func (r userRepo) Update(user *model.User) error {
	return r.db.Updates(&user).Error
}

func (r userRepo) Delete(id int) error {
	var user model.User
	return r.db.Where("id = ?", id).
		Delete(&user).Error
}

func (r userRepo) FindById(id int) (*model.User, error) {
	var user model.User
	err := r.db.Model(&model.User{}).
		Where("id = ?", id).
		First(&user).Error
	if err != nil {
		return nil, errors.New("Người dùng không tồn tại")
	} else {
		return &user, nil
	}
}

func (r userRepo) FindAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	} else {
		return users, nil
	}
}

func (r userRepo) FindByUserName(username string) (*model.User, error) {
	var user *model.User
	err := r.db.Model(&model.User{}).
		Where("username = ?", username).
		First(&user).Error
	return user, err
}
