package service

import (
	"book-store/dto/request"
	"book-store/model"
	"book-store/repository"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type CartService interface {
	Create(ctx *gin.Context, request *request.CartItemRequest) (*model.Cart, error)
	GetCartsByUserId(ctx *gin.Context) []model.Cart
	DeleteById(ctx *gin.Context, id int) error
	Update(ctx *gin.Context, request *request.CartItemUpdateRequest) (*model.Cart, error)
}

type cartServiceImpl struct {
	cartRepo repository.CartRepository
	userRepo repository.UserRepository
	bookRepo repository.BookRepository
}

func newCartService(di *do.Injector) (CartService, error) {
	return &cartServiceImpl{
		cartRepo: do.MustInvoke[repository.CartRepository](di),
		userRepo: do.MustInvoke[repository.UserRepository](di),
		bookRepo: do.MustInvoke[repository.BookRepository](di),
	}, nil
}

func (c cartServiceImpl) Create(ctx *gin.Context, req *request.CartItemRequest) (*model.Cart, error) {
	bookId := req.BookId
	quantity := req.Quantity
	if bookId == 0 {
		return nil, errors.New("Yêu cầu nhập id sách")
	}
	if quantity == 0 {
		return nil, errors.New("Yêu cầu nhập số lượng mua")
	}
	userId := ctx.GetInt("user_id")
	bookExisted, bookErr := c.bookRepo.FindById(bookId)
	if bookErr != nil {
		return nil, bookErr
	}
	if bookExisted.Quantity == 0 {
		return nil, errors.New("Sản phẩm này đã hết hàng")
	}
	if bookExisted.Quantity < quantity {
		return nil, errors.New("Số lượng mua không được phép vượt quá số lượng sản phẩm hiện có")
	}
	cartExisted := c.cartRepo.FindByUserIdAndBookId(userId, bookId)
	if cartExisted != nil {
		return cartExisted, errors.New("Sản phẩm này đã có trong giỏ hàng")
	}
	cartNew := &model.Cart{
		UserID:   userId,
		BookID:   req.BookId,
		Quantity: req.Quantity,
		Price:    bookExisted.Price,
	}
	return c.cartRepo.Create(cartNew)
}

func (c cartServiceImpl) GetCartsByUserId(ctx *gin.Context) []model.Cart {
	userId := ctx.GetInt("user_id")
	return c.cartRepo.FindByUserId(userId)
}

func (c cartServiceImpl) DeleteById(ctx *gin.Context, id int) error {
	userId := ctx.GetInt("user_id")
	cartExisted := c.cartRepo.FindByUserIdAndCartId(userId, id)
	if cartExisted == nil {
		return errors.New("Giỏ hàng này không tồn tại")
	}
	return c.cartRepo.DeleteById(id)
}

func (c cartServiceImpl) Update(ctx *gin.Context, req *request.CartItemUpdateRequest) (*model.Cart, error) {
	cartId := req.CartId
	if cartId == 0 {
		return nil, errors.New("Yêu cầu nhập id giỏ hàng")
	}
	userId := ctx.GetInt("user_id")
	cartExisted := c.cartRepo.FindByUserIdAndCartId(userId, cartId)
	if cartExisted == nil {
		return nil, errors.New("Giỏ hàng này không tồn tại")
	}
	quantityReq := req.Quantity
	book, _ := c.bookRepo.FindById(cartExisted.BookID)
	if quantityReq >= book.Quantity {
		return nil, errors.New("Số lượng mua không được phép vượt quá số lượng sản phẩm hiện có")
	}
	if quantityReq == 0 {
		quantityReq = cartExisted.Quantity
	}
	cartExisted.Quantity = quantityReq
	err := c.cartRepo.Update(cartExisted)
	if err != nil {
		return nil, err
	}
	return cartExisted, nil
}
