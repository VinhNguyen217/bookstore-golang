package controller

import (
	"book-store/dto/request"
	"book-store/service"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"net/http"
	"strconv"
)

type CartController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	GetCartsByUserId(ctx *gin.Context)
	DeleteCartById(ctx *gin.Context)
}

type cartCtl struct {
	cartService service.CartService
}

func NewCartController(di *do.Injector) CartController {
	return &cartCtl{
		cartService: do.MustInvoke[service.CartService](di),
	}
}

func (c *cartCtl) Create(ctx *gin.Context) {
	req := &request.CartItemRequest{}
	_ = ctx.ShouldBindJSON(req)
	cart, err := c.cartService.Create(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, cart)
}

func (c *cartCtl) Update(ctx *gin.Context) {
	req := &request.CartItemUpdateRequest{}
	_ = ctx.ShouldBindJSON(req)
	cart, err := c.cartService.Update(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": cart})
}

func (c *cartCtl) GetCartsByUserId(ctx *gin.Context) {
	carts := c.cartService.GetCartsByUserId(ctx)
	ctx.JSON(http.StatusOK, carts)
}

func (c *cartCtl) DeleteCartById(ctx *gin.Context) {
	id, errId := strconv.Atoi(ctx.Param("id"))
	if errId != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": errId.Error()})
		return
	}
	err := c.cartService.DeleteById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Xóa thành công giỏ hàng"})
}
