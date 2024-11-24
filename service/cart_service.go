package service

import (
	"book-store/dto/request"
	"book-store/dto/response"
	"book-store/model"
	"book-store/repository"
	"book-store/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type CartService interface {
	Create(ctx *gin.Context, request *request.CartItemRequest) (*response.CartRes, error)
	GetCartsByUserId(ctx *gin.Context) []response.CartRes
	DeleteById(ctx *gin.Context, id int) error
	Update(ctx *gin.Context, request *request.CartItemUpdateRequest) (*response.CartRes, error)
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

func (c cartServiceImpl) Create(ctx *gin.Context, req *request.CartItemRequest) (*response.CartRes, error) {
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
		return nil, errors.New("Sản phẩm này đã có trong giỏ hàng")
	}
	cartNew := &model.Cart{
		UserID:   userId,
		BookID:   req.BookId,
		Quantity: quantity,
		Price:    bookExisted.Price * quantity,
	}
	_, cartErr := c.cartRepo.Create(cartNew)
	if cartErr != nil {
		return nil, cartErr
	} else {
		return convertCart(cartNew), nil
	}
}

func (c cartServiceImpl) GetCartsByUserId(ctx *gin.Context) []response.CartRes {
	userId := ctx.GetInt("user_id")
	carts := c.cartRepo.FindByUserId(userId)
	var cartResList []response.CartRes
	for _, cart := range carts {
		cartRes := convertCart(&cart)
		cartResList = append(cartResList, *cartRes)
	}
	return cartResList
}

func (c cartServiceImpl) DeleteById(ctx *gin.Context, id int) error {
	userId := ctx.GetInt("user_id")
	cartExisted := c.cartRepo.FindByUserIdAndCartId(userId, id)
	if cartExisted == nil {
		return errors.New("Giỏ hàng này không tồn tại")
	}
	return c.cartRepo.DeleteById(id)
}

func (c cartServiceImpl) Update(ctx *gin.Context, req *request.CartItemUpdateRequest) (*response.CartRes, error) {
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
		return convertCart(cartExisted), nil
	}
	priceNew := (cartExisted.Price / cartExisted.Quantity) * quantityReq
	cartExisted.Quantity = quantityReq
	cartExisted.Price = priceNew
	err := c.cartRepo.Update(cartExisted)
	if err != nil {
		return nil, err
	}
	return convertCart(cartExisted), nil
}

func convertCart(cart *model.Cart) *response.CartRes {
	return &response.CartRes{
		ID:       cart.ID,
		BookID:   cart.BookID,
		Quantity: cart.Quantity,
		Price:    utils.ConvertToVND(cart.Price),
	}
}
