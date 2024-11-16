package controller

import (
	"book-store/dto/request"
	"book-store/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"net/http"
	"strconv"
)

type UserController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindById(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type userCtl struct {
	userService service.UserService
}

func NewUserController(di *do.Injector) UserController {
	return &userCtl{
		userService: do.MustInvoke[service.UserService](di),
	}
}

func (c *userCtl) Create(ctx *gin.Context) {
	req := &request.UserRequest{}
	_ = ctx.ShouldBindJSON(req)
	user, err := c.userService.Create(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": user})
}

func (c *userCtl) Update(ctx *gin.Context) {
	id, errId := strconv.Atoi(ctx.Param("id"))
	if errId != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": errId.Error()})
		return
	}
	req := &request.UserRequest{}
	_ = ctx.ShouldBind(req)
	user, err := c.userService.Update(req, id)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func (c *userCtl) Delete(ctx *gin.Context) {
	id, errId := strconv.Atoi(ctx.Param("id"))
	if errId != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": errId.Error()})
		return
	}
	err := c.userService.Delete(id)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Xóa thành công"})
}

func (c *userCtl) FindById(ctx *gin.Context) {
	id, errId := strconv.Atoi(ctx.Param("id"))
	if errId != nil {
		ctx.JSON(http.StatusBadRequest, errId.Error())
		return
	}
	user, err := c.userService.FindByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func (c *userCtl) FindAll(ctx *gin.Context) {
	users, err := c.userService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, users)
	}
}