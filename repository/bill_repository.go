package repository

import (
	"book-store/model"
	"errors"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type BillRepository interface {
	Create(bill *model.Bill) (*model.Bill, error)
	FindByIdAndUserId(id, userId int) (*model.Bill, error)
	Update(bill *model.Bill) error
	FindById(id int) (*model.Bill, error)
	FindByUserId(userId int) []model.Bill
	FindAll() []model.Bill
}

type billRepo struct {
	db *gorm.DB
}

func newBillRepository(di *do.Injector) (BillRepository, error) {
	db := do.MustInvoke[*gorm.DB](di)
	return &billRepo{db: db}, nil
}

func (r billRepo) Create(bill *model.Bill) (*model.Bill, error) {
	err := r.db.Create(bill).Error
	return bill, err
}

func (r billRepo) FindByIdAndUserId(id, userId int) (*model.Bill, error) {
	var bill model.Bill
	r.db.Model(&model.Bill{}).
		Where("user_id = ? AND id = ?", userId, id).
		Find(&bill)
	if bill.ID == 0 {
		return nil, errors.New("Đơn hàng không tồn tại")
	} else {
		return &bill, nil
	}
}

func (r billRepo) Update(bill *model.Bill) error {
	return r.db.Save(bill).Error
}

func (r billRepo) FindById(id int) (*model.Bill, error) {
	var bill model.Bill
	err := r.db.Model(&model.Bill{}).
		Where("id = ?", id).
		First(&bill).Error
	if err != nil {
		return nil, errors.New("Đơn hàng này không tồn tại")
	} else {
		return &bill, nil
	}
}

func (r billRepo) FindByUserId(userId int) []model.Bill {
	var bills []model.Bill
	r.db.Where("user_id = ?", userId).
		Find(&bills)
	return bills
}

func (r billRepo) FindAll() []model.Bill {
	var bills []model.Bill
	_ = r.db.Find(&bills)
	return bills
}
