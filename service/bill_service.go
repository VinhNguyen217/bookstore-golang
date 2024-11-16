package service

import (
	"book-store/dto/request"
	"book-store/enum"
	"book-store/model"
	"book-store/repository"
	"github.com/samber/do"
)

type BillService interface {
	Create(bill *request.BillRequest) (*model.Bill, error)
}

type billServiceImpl struct {
	billRepo       repository.BillRepository
	cartRepo       repository.CartRepository
	billDetailRepo repository.BillDetailRepository
}

func newBillService(di *do.Injector) (BillService, error) {
	return &billServiceImpl{
		billRepo:       do.MustInvoke[repository.BillRepository](di),
		cartRepo:       do.MustInvoke[repository.CartRepository](di),
		billDetailRepo: do.MustInvoke[repository.BillDetailRepository](di),
	}, nil
}

func (r billServiceImpl) Create(req *request.BillRequest) (*model.Bill, error) {
	var cartIds = req.CartIds
	var carts []model.Cart
	var total = 0
	for _, cartId := range cartIds {
		cart, err := r.cartRepo.FindById(cartId)
		if err != nil {
			return nil, err
		} else {
			total += cart.Price
			carts = append(carts, *cart)
		}
	}
	// Create bill
	bill := &model.Bill{
		Receiver: req.Receiver,
		Phone:    req.Phone,
		Address:  req.Address,
		Email:    req.Email,
		Note:     req.Note,
		Total:    total,
		Status:   enum.WAIT_CONFIRM,
		Payment:  enum.CASH,
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
			if err != nil {
				return nil, err
			}
		}
	}
	return bill, nil
}
