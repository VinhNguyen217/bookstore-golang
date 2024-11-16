package controller

import (
	"book-store/dto/request"
	"book-store/service"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"net/http"
)

type BillController interface {
	Create(ctx *gin.Context)
}

type billCtl struct {
	billService service.BillService
}

func NewBillController(di *do.Injector) BillController {
	return &billCtl{
		billService: do.MustInvoke[service.BillService](di),
	}
}

func (c *billCtl) Create(ctx *gin.Context) {
	req := &request.BillRequest{}
	_ = ctx.ShouldBindJSON(req)
	bill, err := c.billService.Create(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": bill})
}
