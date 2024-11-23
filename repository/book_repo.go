package repository

import (
	"book-store/model"
	"errors"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book *model.Book) (*model.Book, error)
	Update(book *model.Book) error
	Delete(id int) error
	FindAll() ([]model.Book, error)
	FindById(id int) (*model.Book, error)
	FindByName(name string) (*model.Book, error)
}

type bookRepo struct {
	db *gorm.DB
}

func newBookRepository(di *do.Injector) (BookRepository, error) {
	db := do.MustInvoke[*gorm.DB](di)
	return &bookRepo{db: db}, nil
}

func (r *bookRepo) Create(book *model.Book) (*model.Book, error) {
	err := r.db.Create(book).Error
	return book, err
}

func (r *bookRepo) FindAll() ([]model.Book, error) {
	var books []model.Book
	err := r.db.Find(&books).Error
	if err != nil {
		return books, err
	} else {
		return books, nil
	}
}

func (r *bookRepo) FindById(id int) (*model.Book, error) {
	var book model.Book
	err := r.db.Model(&model.Book{}).
		Where("id = ?", id).
		First(&book).Error
	if err != nil {
		return nil, errors.New("Sách này không tồn tại")
	} else {
		return &book, nil
	}
}

func (r *bookRepo) Update(book *model.Book) error {
	return r.db.Updates(book).Error
}

func (r *bookRepo) Delete(id int) error {
	var book model.Book
	return r.db.Where("id = ?", id).Delete(&book).Error
}

func (r *bookRepo) FindByName(name string) (*model.Book, error) {
	var book model.Book
	err := r.db.Model(&model.Book{}).
		Where("name = ?", name).
		First(&book).Error
	if err != nil {
		return nil, errors.New("Sách này không tồn tại")
	} else {
		return &book, nil
	}
}
