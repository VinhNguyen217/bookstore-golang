package controller

import (
	"book-store/dto"
	"book-store/service"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"net/http"
)

type AuthController interface {
	PasswordLogin(*gin.Context)
}

type authCtl struct {
	authService service.AuthService
}

func NewAuthController(di *do.Injector) AuthController {
	return &authCtl{
		authService: do.MustInvoke[service.AuthService](di),
	}
}

func (c *authCtl) PasswordLogin(ctx *gin.Context) {
	req := &dto.PasswordLoginRequest{}
	_ = ctx.ShouldBind(req)
	resp, err := c.authService.PasswordLogin(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
