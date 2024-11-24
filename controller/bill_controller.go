package controller

import (
	"book-store/dto/request"
	"book-store/service"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"net/http"
	"strconv"
)

type BillController interface {
	Create(ctx *gin.Context)
	CancelBill(ctx *gin.Context)
	UpdateStatusBill(ctx *gin.Context)
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
	bill, err := c.billService.Create(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": bill})
}

func (c *billCtl) CancelBill(ctx *gin.Context) {
	id, errId := strconv.Atoi(ctx.Param("id"))
	if errId != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Id không hợp lệ"})
		return
	}
	bill, err := c.billService.CancelBill(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Hủy đơn hàng thành công", "data": bill})
}

func (c *billCtl) UpdateStatusBill(ctx *gin.Context) {
	id, errId := strconv.Atoi(ctx.Param("id"))
	if errId != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Id không hợp lệ"})
		return
	}
	billStatusReq := &request.BillStatusRequest{}
	_ = ctx.ShouldBindJSON(billStatusReq)
	bill, err := c.billService.UpdateStatus(ctx, id, billStatusReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Cập nhật đơn hàng thành công", "data": bill})
}
