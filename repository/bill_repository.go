package repository

import (
	"book-store/model"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type BillRepository interface {
	Create(bill *model.Bill) (*model.Bill, error)
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
