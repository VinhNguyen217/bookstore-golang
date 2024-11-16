package repository

import (
	"book-store/model"
	"errors"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type CartRepository interface {
	Create(cart *model.Cart) (*model.Cart, error)
	FindByUserIdAndBookId(userId int, bookId int) *model.Cart
	FindByUserId(userId int) []model.Cart
	FindById(id int) (*model.Cart, error)
	DeleteById(id int) error
	Update(cart *model.Cart) error
}

type cartRepo struct {
	db *gorm.DB
}

func newCartRepository(di *do.Injector) (CartRepository, error) {
	db := do.MustInvoke[*gorm.DB](di)
	return &cartRepo{db: db}, nil
}

func (c cartRepo) FindById(id int) (*model.Cart, error) {
	var cart model.Cart
	err := c.db.Model(&model.Cart{}).
		Where("id = ?", id).
		First(&cart).Error
	if err != nil {
		return nil, errors.New("Giỏ hàng không tồn tại")
	} else {
		return &cart, nil
	}
}

func (c cartRepo) Create(cart *model.Cart) (*model.Cart, error) {
	err := c.db.Create(cart).Error
	return cart, err
}

func (c cartRepo) FindByUserIdAndBookId(userId int, bookId int) *model.Cart {
	var cart model.Cart
	c.db.Model(&model.Cart{}).
		Where("user_id = ? AND book_id = ?", userId, bookId).
		Find(&cart)
	if cart.ID == 0 {
		return nil
	} else {
		return &cart
	}
}

func (c cartRepo) FindByUserId(userId int) []model.Cart {
	var carts []model.Cart
	c.db.Where("user_id = ?", userId).
		Find(&carts)
	return carts
}

func (c cartRepo) DeleteById(id int) error {
	return c.db.Delete(&model.Cart{}, id).Error
}

func (c cartRepo) Update(cart *model.Cart) error {
	return c.db.Updates(cart).Error
}
