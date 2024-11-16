package service

import (
	"book-store/dto/request"
	"book-store/model"
	"book-store/repository"
	"errors"
	"github.com/samber/do"
)

type CartService interface {
	Create(request *request.CartItemRequest) (*model.Cart, error)
	GetCartsByUserId(userId int) []model.Cart
	DeleteById(id int) error
}

type cartServiceImpl struct {
	cartRepo repository.CartRepository
}

func newCartService(di *do.Injector) (CartService, error) {
	return &cartServiceImpl{
		cartRepo: do.MustInvoke[repository.CartRepository](di),
	}, nil
}

func (c cartServiceImpl) Create(req *request.CartItemRequest) (*model.Cart, error) {
	var userId = req.UserId
	var bookId = req.BookId
	cartFind := c.cartRepo.FindByUserIdAndBookId(userId, bookId)
	if cartFind == nil {
		cartNew := &model.Cart{
			UserID:   req.UserId,
			BookID:   req.BookId,
			Quantity: req.Quantity,
		}
		return c.cartRepo.Create(cartNew)
	} else {
		return cartFind, errors.New("Sản phẩm này đã có trong giỏ hàng")
	}
}

func (c cartServiceImpl) GetCartsByUserId(userId int) []model.Cart {
	return c.cartRepo.FindByUserId(userId)
}

func (c cartServiceImpl) DeleteById(id int) error {
	return c.cartRepo.DeleteById(id)
}
