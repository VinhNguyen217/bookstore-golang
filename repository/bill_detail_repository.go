package repository

import (
	"book-store/model"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type BillDetailRepository interface {
	Create(billDetail *model.BillDetail) (*model.BillDetail, error)
	FindByBillId(billId int) []model.BillDetail
}

type billDetailRepo struct {
	db *gorm.DB
}

func newBillDetailRepository(di *do.Injector) (BillDetailRepository, error) {
	db := do.MustInvoke[*gorm.DB](di)
	return &billDetailRepo{db: db}, nil
}

func (r billDetailRepo) Create(billDetail *model.BillDetail) (*model.BillDetail, error) {
	err := r.db.Create(billDetail).Error
	return billDetail, err
}

func (r billDetailRepo) FindByBillId(billId int) []model.BillDetail {
	var billDetails []model.BillDetail
	r.db.Where("bill_id = ?", billId).Find(&billDetails)
	return billDetails
}
