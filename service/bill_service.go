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
	"strings"
	"time"
)

type BillService interface {
	Create(ctx *gin.Context, bill *request.BillRequest) (*response.BillRes, error)
	CancelBill(ctx *gin.Context, id int) (*response.BillRes, error)
	UpdateStatus(id int, billStatus *request.BillStatusRequest) (*response.BillRes, error)
	FindAllByUserId(ctx *gin.Context) ([]response.BillRes, error)
	FindAll() ([]response.BillRes, error)
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
	if strings.TrimSpace(req.Receiver) == "" {
		return nil, errors.New("Receiver is required")
	}
	if strings.TrimSpace(req.Address) == "" {
		return nil, errors.New("Address is required")
	}
	if strings.TrimSpace(req.Phone) == "" {
		return nil, errors.New("Phone is required")
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
	// Create
	bill := &model.Bill{
		Receiver: req.Receiver,
		UserId:   userId,
		Phone:    req.Phone,
		Address:  req.Address,
		Email:    req.Email,
		Note:     req.Note,
		Total:    total,
		//ConfirmDate: time.Date(0001, 2, 1,
		//	00, 00, 00, 00, time.UTC),
		//CreatedDate: time.Now(),
		CreatedDate: time.Now().UTC().Format("2006-01-02 15:04:05"),
		//UpdatedDate: time.Date(0001, 2, 1,
		//	00, 00, 00, 00, time.UTC),
		Status:  enum.WAIT_CONFIRM,
		Payment: enum.CASH,
	}
	_, err := r.billRepo.Create(bill)
	var billDetails []response.BillDetailRes
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
			book, _ := r.bookRepo.FindById(cart.BookID)
			billDetailRes := response.BillDetailRes{
				ID:        billDetail.ID,
				BookID:    billDetail.BookID,
				BookName:  book.Name,
				Quantity:  billDetail.Quantity,
				Price:     utils.ConvertToVND(billDetail.Price),
				UnitPrice: utils.ConvertToVND(billDetail.Price / billDetail.Quantity),
			}
			billDetails = append(billDetails, billDetailRes)
			_ = r.cartRepo.DeleteById(cart.ID)
			if err != nil {
				return nil, err
			}
		}
	}
	billRes := convertBill(bill)
	billRes.BillDetails = billDetails
	return billRes, nil
}

func (r billServiceImpl) CancelBill(ctx *gin.Context, id int) (*response.BillRes, error) {
	userId := ctx.GetInt("user_id")
	billExisted, err := r.billRepo.FindByIdAndUserId(id, userId)
	if err != nil {
		return nil, err
	}
	status := billExisted.Status
	if status == enum.WAIT_CONFIRM || status == enum.DELIVERY {
		//billExisted.UpdatedDate = time.Now()
		billExisted.UpdatedDate = time.Now().UTC().Format("2006-01-02 15:04:05")
		billExisted.Status = enum.CANCELLED
		_ = r.billRepo.Update(billExisted)
		return convertBill(billExisted), nil
	}
	if status == enum.CANCELLED {
		return nil, errors.New("Đơn hàng này đã bị hủy")
	}
	if status == enum.DELIVERED {
		return nil, errors.New("Đơn hàng này đã hoàn thành")
	}
	return nil, errors.New("Không thể hủy đơn hàng này")
}

func (r billServiceImpl) UpdateStatus(id int, billStatus *request.BillStatusRequest) (*response.BillRes, error) {
	billExisted, err := r.billRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	if billExisted.Status == enum.CANCELLED {
		return nil, errors.New("Đơn hàng này đã bị hủy, không thể cập nhật được")
	}
	if billExisted.Status == enum.DELIVERED {
		return nil, errors.New("Đơn hàng này đã hoàn thành, không thể cập nhật được")
	}
	status := billStatus.Status
	switch status {
	case enum.DELIVERY:
		if billExisted.Status == enum.DELIVERY {
			return nil, errors.New("Đơn hàng này đang được giao, không thể cập nhật được")
		}
		billExisted.Status = enum.DELIVERY
		//billExisted.ConfirmDate = time.Now()
		billExisted.ConfirmDate = time.Now().UTC().Format("2006-01-02 15:04:05")
		_ = r.billRepo.Update(billExisted)
		return convertBill(billExisted), nil
	case enum.DELIVERED:
		if billExisted.Status == enum.WAIT_CONFIRM {
			return nil, errors.New("Đơn hàng này đang chờ xác nhận, không thể cập nhật được")
		}
		billExisted.Status = enum.DELIVERED
		//billExisted.CreatedDate = time.Now()
		billExisted.UpdatedDate = time.Now().UTC().Format("2006-01-02 15:04:05")
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

func (r billServiceImpl) FindAllByUserId(ctx *gin.Context) ([]response.BillRes, error) {
	userId := ctx.GetInt("user_id")
	bills := r.billRepo.FindByUserId(userId)
	var billResList []response.BillRes
	for _, bill := range bills {
		billRes := convertBill(&bill)
		billId := bill.ID
		billDetails := r.billDetailRepo.FindByBillId(billId)
		var billDetailResList []response.BillDetailRes
		for _, billDetail := range billDetails {
			billDetailRes := convertBillDetail(&billDetail)
			book, _ := r.bookRepo.FindById(billDetail.BookID)
			billDetailRes.BookName = book.Name
			billDetailResList = append(billDetailResList, *billDetailRes)
		}
		billRes.BillDetails = billDetailResList
		billResList = append(billResList, *billRes)
	}
	return billResList, nil
}

func (r billServiceImpl) FindAll() ([]response.BillRes, error) {
	bills := r.billRepo.FindAll()
	var billResList []response.BillRes
	for _, bill := range bills {
		billRes := convertBill(&bill)
		billId := bill.ID
		billDetails := r.billDetailRepo.FindByBillId(billId)
		var billDetailResList []response.BillDetailRes
		for _, billDetail := range billDetails {
			billDetailRes := convertBillDetail(&billDetail)
			book, _ := r.bookRepo.FindById(billDetail.BookID)
			billDetailRes.BookName = book.Name
			billDetailResList = append(billDetailResList, *billDetailRes)
		}
		billRes.BillDetails = billDetailResList
		billResList = append(billResList, *billRes)
	}
	return billResList, nil
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
		CreatedDate: bill.CreatedDate,
	}
	switch bill.Status {
	case enum.WAIT_CONFIRM:
		billRes.Status = "Chờ xác nhận"
	case enum.DELIVERY:
		billRes.Status = "Đang giao hàng"
		billRes.ConfirmDate = bill.ConfirmDate
	case enum.DELIVERED:
		billRes.Status = "Đã giao hàng"
		billRes.ConfirmDate = bill.ConfirmDate
		billRes.UpdatedDate = bill.UpdatedDate
	case enum.CANCELLED:
		billRes.Status = "Đã hủy"
		billRes.UpdatedDate = bill.UpdatedDate
	}
	return billRes
}

func convertBillDetail(billDetail *model.BillDetail) *response.BillDetailRes {
	billDetailRes := &response.BillDetailRes{
		ID:        billDetail.ID,
		BookID:    billDetail.BookID,
		Quantity:  billDetail.Quantity,
		Price:     utils.ConvertToVND(billDetail.Price),
		UnitPrice: utils.ConvertToVND(billDetail.Price / billDetail.Quantity),
	}
	return billDetailRes
}
