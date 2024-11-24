package service

import (
	"book-store/dto/request"
	"book-store/dto/response"
	"book-store/enum"
	"book-store/model"
	"book-store/repository"
	"book-store/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"time"
)

type BillService interface {
	Create(ctx *gin.Context, bill *request.BillRequest) (*response.BillRes, error)
	CancelBill(ctx *gin.Context, id int) (*response.BillRes, error)
	UpdateStatus(ctx *gin.Context, id int, billStatus *request.BillStatusRequest) (*response.BillRes, error)
}

type billServiceImpl struct {
	billRepo       repository.BillRepository
	cartRepo       repository.CartRepository
	billDetailRepo repository.BillDetailRepository
	bookRepo       repository.BookRepository
}

func newBillService(di *do.Injector) (BillService, error) {
	return &billServiceImpl{
		billRepo:       do.MustInvoke[repository.BillRepository](di),
		cartRepo:       do.MustInvoke[repository.CartRepository](di),
		billDetailRepo: do.MustInvoke[repository.BillDetailRepository](di),
		bookRepo:       do.MustInvoke[repository.BookRepository](di),
	}, nil
}

func (r billServiceImpl) Create(ctx *gin.Context, req *request.BillRequest) (*response.BillRes, error) {
	cartIds := req.CartIds
	if len(cartIds) == 0 {
		return nil, errors.New("Yêu cầu nhập danh sách giỏ hàng")
	}

	userId := ctx.GetInt("user_id")
	var carts []model.Cart
	var total = 0
	for _, cartId := range cartIds {
		cart := r.cartRepo.FindByUserIdAndCartId(userId, cartId)
		if cart == nil {
			return nil, errors.New("Giỏ hàng này không tồn tại")
		} else {
			total += cart.Price
			carts = append(carts, *cart)
		}
	}
	// Create bill
	bill := &model.Bill{
		Receiver: req.Receiver,
		UserId:   userId,
		Phone:    req.Phone,
		Address:  req.Address,
		Email:    req.Email,
		Note:     req.Note,
		Total:    total,
		ConfirmDate: time.Date(0001, 2, 1,
			00, 00, 00, 00, time.UTC),
		Status:  enum.WAIT_CONFIRM,
		Payment: enum.CASH,
	}
	_, err := r.billRepo.Create(bill)
	if err != nil {
		return nil, err
	} else {
		// Create bill_detail
		for _, cart := range carts {
			billDetail := &model.BillDetail{
				BillID:   bill.ID,
				BookID:   cart.BookID,
				Quantity: cart.Quantity,
				Price:    cart.Price,
			}
			_, err := r.billDetailRepo.Create(billDetail)
			_ = r.cartRepo.DeleteById(cart.ID)
			if err != nil {
				return nil, err
			}
		}
	}

	return convertBill(bill), nil
}

func (r billServiceImpl) CancelBill(ctx *gin.Context, id int) (*response.BillRes, error) {
	userId := ctx.GetInt("user_id")
	billExisted, err := r.billRepo.FindByIdAndUserId(id, userId)
	if err != nil {
		return nil, err
	}
	status := billExisted.Status
	if status == enum.WAIT_CONFIRM || status == enum.DELIVERY {
		billExisted.Status = enum.CANCELLED
		_ = r.billRepo.Update(billExisted)
		return convertBill(billExisted), nil
	}
	if status == enum.CANCELLED {
		return nil, errors.New("Đơn hàng này đã bị hủy")
	}
	return nil, errors.New("Không thể hủy đơn hàng này")
}

func (r billServiceImpl) UpdateStatus(ctx *gin.Context, id int, billStatus *request.BillStatusRequest) (*response.BillRes, error) {
	billExisted, err := r.billRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	status := billStatus.Status
	switch status {
	case enum.DELIVERY:
		if billExisted.Status == enum.CANCELLED {
			return nil, errors.New("Đơn hàng này đã bị hủy, không thể xác nhận được")
		} else {
			billExisted.Status = enum.DELIVERY
			billExisted.ConfirmDate = time.Now()
			_ = r.billRepo.Update(billExisted)
			return convertBill(billExisted), nil
		}
	case enum.DELIVERED:
		billExisted.Status = enum.DELIVERED
		billExisted.UpdatedAt = time.Now()
		billDetails := r.billDetailRepo.FindByBillId(billExisted.ID)
		for _, billDetail := range billDetails { // Cập nhật số lượng sách sau khi bán
			quantity := billDetail.Quantity
			bookId := billDetail.BookID
			book, _ := r.bookRepo.FindById(bookId)
			book.Sold += quantity
			book.Quantity -= quantity
			_ = r.bookRepo.Update(book)
		}
		_ = r.billRepo.Update(billExisted)
		return convertBill(billExisted), nil
	default:
		return nil, errors.New("Yêu cầu không hợp lệ")
	}
}

func convertBill(bill *model.Bill) *response.BillRes {
	billRes := &response.BillRes{
		ID:          bill.ID,
		Receiver:    bill.Receiver,
		Phone:       bill.Phone,
		Address:     bill.Address,
		Email:       bill.Email,
		Note:        bill.Note,
		Total:       utils.ConvertToVND(bill.Total),
		Status:      bill.Status,
		Payment:     bill.Payment,
		ConfirmDate: bill.ConfirmDate,
	}
	switch bill.Status {
	case enum.WAIT_CONFIRM:
		billRes.Status = "Chờ xác nhận"
	case enum.DELIVERY:
		billRes.Status = "Đang giao hàng"
	case enum.DELIVERED:
		billRes.Status = "Đã giao hàng"
	case enum.CANCELLED:
		billRes.Status = "Đã hủy"
	}
	return billRes
}
