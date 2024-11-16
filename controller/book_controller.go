package controller

import (
	"book-store/dto/request"
	"book-store/service"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"net/http"
	"strconv"
)

type BookController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindById(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type bookCtl struct {
	bookService service.BookService
}

func NewBookController(di *do.Injector) BookController {
	return &bookCtl{
		bookService: do.MustInvoke[service.BookService](di),
	}
}

func (c *bookCtl) Create(ctx *gin.Context) {
	req := &request.BookRequest{}
	_ = ctx.ShouldBind(req)
	book, err := c.bookService.Create(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": book})
}

func (c *bookCtl) Update(ctx *gin.Context) {
	id, errId := strconv.Atoi(ctx.Param("id"))
	if errId != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": errId.Error()})
		return
	}
	req := &request.BookRequest{}
	_ = ctx.ShouldBind(req)
	book, err := c.bookService.Update(req, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

func (c *bookCtl) Delete(ctx *gin.Context) {
	id, errId := strconv.Atoi(ctx.Param("id"))
	if errId != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": errId.Error()})
		return
	}
	err := c.bookService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "Xóa thành công"})
}

func (c *bookCtl) FindById(ctx *gin.Context) {
	id, errId := strconv.Atoi(ctx.Param("id"))
	if errId != nil {
		ctx.JSON(http.StatusBadRequest, errId.Error())
		return
	}
	res, err := c.bookService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": res})
}

func (c *bookCtl) FindAll(ctx *gin.Context) {
	books, err := c.bookService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	} else {
		ctx.JSON(http.StatusOK, gin.H{"data": books})
	}
}
