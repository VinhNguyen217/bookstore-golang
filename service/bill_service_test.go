package service

import (
	"book-store/dto/request"
	mockRepo "book-store/mock/book-store/repository"
	"book-store/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateBill(t *testing.T) {
	mockCartRepo := mockRepo.MockCartRepository{}
	mockBillRepo := mockRepo.MockBillRepository{}
	mockBillDetailRepo := mockRepo.MockBillDetailRepository{}
	mockBookRepo := mockRepo.MockBookRepository{}

	billService := &billServiceImpl{
		cartRepo:       &mockCartRepo,
		billRepo:       &mockBillRepo,
		billDetailRepo: &mockBillDetailRepo,
		bookRepo:       &mockBookRepo,
	}

	mockCart := &model.Cart{
		ID:       1,
		UserID:   1,
		BookID:   1,
		Quantity: 2,
		Price:    100,
	}
	mockCartRepo.On("FindByUserIdAndCartId", 1, 1).Return(mockCart)

	mockBook := &model.Book{
		ID:   1,
		Name: "Book 1",
	}
	mockBookRepo.On("FindById", 1).Return(mockBook, nil)

	mockBill := &model.Bill{
		ID:       1,
		Receiver: "Receiver",
		UserId:   1,
		Phone:    "123456789",
		Address:  "Address",
		Email:    "email@example.com",
		Note:     "Note",
		Total:    200,
		Status:   "WAIT_CONFIRM",
		Payment:  "CASH",
	}
	mockBillRepo.
		On("Create", mock.Anything).
		Return(mockBill, nil)

	mockBillDetail := &model.BillDetail{
		ID:       1,
		BillID:   1,
		BookID:   1,
		Quantity: 2,
		Price:    200,
	}
	mockBillDetailRepo.
		On("Create", mock.Anything).
		Return(mockBillDetail, nil)

	mockCartRepo.
		On("DeleteById", mockCart.ID).
		Return(nil)

	req := &request.BillRequest{
		CartIds:  []int{1},
		Receiver: "Receiver",
		Phone:    "123456789",
		Address:  "Address",
		Email:    "email@example.com",
		Note:     "Note",
	}

	resp, err := billService.Create(nil, req)
	assert.NoError(t, err)

	assert.Equal(t, resp.Receiver, "Receiver")
	assert.Equal(t, resp.Total, 200)
	assert.Len(t, resp.BillDetails, 1)
	assert.Equal(t, resp.BillDetails[0].BookName, "Book 1")
	assert.Equal(t, resp.BillDetails[0].Price, "200 VND")
}

func TestCreateBill_EmptyCartIds(t *testing.T) {
	mockCartRepo := mockRepo.MockCartRepository{}
	mockBillRepo := mockRepo.MockBillRepository{}
	mockBillDetailRepo := mockRepo.MockBillDetailRepository{}
	mockBookRepo := mockRepo.MockBookRepository{}

	billService := &billServiceImpl{
		cartRepo:       &mockCartRepo,
		billRepo:       &mockBillRepo,
		billDetailRepo: &mockBillDetailRepo,
		bookRepo:       &mockBookRepo,
	}

	req := &request.BillRequest{
		CartIds: []int{},
	}

	// Gọi hàm Create với dữ liệu CartIds trống
	_, err := billService.Create(nil, req)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "Yêu cầu nhập danh sách giỏ hàng")
}

func TestCreateBill_CartNotFound(t *testing.T) {
	mockCartRepo := mockRepo.MockCartRepository{}
	mockBillRepo := mockRepo.MockBillRepository{}
	mockBillDetailRepo := mockRepo.MockBillDetailRepository{}
	mockBookRepo := mockRepo.MockBookRepository{}

	billService := &billServiceImpl{
		cartRepo:       &mockCartRepo,
		billRepo:       &mockBillRepo,
		billDetailRepo: &mockBillDetailRepo,
		bookRepo:       &mockBookRepo,
	}

	// Giả lập giỏ hàng không tồn tại
	mockCartRepo.On("FindByUserIdAndCartId", 1, 1).Return(nil)

	req := &request.BillRequest{
		CartIds:  []int{1},
		Receiver: "Receiver",
		Phone:    "123456789",
		Address:  "Address",
		Email:    "email@example.com",
		Note:     "Note",
	}

	// Gọi hàm Create với giỏ hàng không tồn tại
	_, err := billService.Create(nil, req)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "Giỏ hàng này không tồn tại")
}
