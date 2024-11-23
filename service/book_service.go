package service

import (
	"book-store/dto/request"
	"book-store/model"
	"book-store/repository"
	"errors"
	"github.com/samber/do"
	"strings"
	"time"
)

type BookService interface {
	Create(book *request.BookRequest) (*model.Book, error)
	Update(book *request.BookRequest, id int) (*model.Book, error)
	Delete(id int) error
	FindById(id int) (*model.Book, error)
	FindAll() ([]model.Book, error)
}

type bookServiceImpl struct {
	bookRepo repository.BookRepository
}

func newBookService(di *do.Injector) (BookService, error) {
	return &bookServiceImpl{
		bookRepo: do.MustInvoke[repository.BookRepository](di),
	}, nil
}

func (bookService *bookServiceImpl) Create(req *request.BookRequest) (*model.Book, error) {
	if strings.TrimSpace(req.Name) == "" {
		return nil, errors.New("Yêu cầu nhập tên sách")
	}
	if req.Price == 0 {
		return nil, errors.New("Yêu cầu nhập giá tiền")
	}
	bookExist, _ := bookService.bookRepo.FindByName(req.Name)
	if bookExist != nil {
		return nil, errors.New("Tên sách đã tồn tại")
	}
	date, err := time.Parse("2006-01-02", req.PublishDate)
	if err != nil {
		return nil, err
	}
	book := &model.Book{
		Name:        req.Name,
		Quantity:    req.Quantity,
		Price:       req.Price,
		PublishDate: date,
		Description: req.Description,
	}
	return bookService.bookRepo.Create(book)
}

func (bookService *bookServiceImpl) Update(req *request.BookRequest, id int) (*model.Book, error) {
	bookExisted, errId := bookService.bookRepo.FindById(id)
	if errId != nil {
		return nil, errors.New("Sách này không tồn tại")
	}
	// validate
	nameReq := req.Name
	if strings.TrimSpace(nameReq) != "" {
		bookExisted.Name = nameReq
	}
	quantityReq := req.Quantity
	if quantityReq != 0 {
		bookExisted.Quantity = quantityReq
	}
	priceReq := req.Price
	if priceReq != 0 {
		bookExisted.Price = priceReq
	}
	var date time.Time
	if req.PublishDate != "" {
		date, _ = time.Parse("2006-01-02", req.PublishDate)
		bookExisted.PublishDate = date
	}
	descriptionReq := req.Description
	if descriptionReq != "" {
		bookExisted.Description = descriptionReq
	}
	err := bookService.bookRepo.Update(bookExisted)
	if err != nil {
		return nil, err
	} else {
		return bookService.bookRepo.FindById(id)
	}
}

func (bookService *bookServiceImpl) Delete(id int) error {
	_, err := bookService.bookRepo.FindById(id)
	if err != nil {
		return err
	}
	return bookService.bookRepo.Delete(id)
}

func (bookService *bookServiceImpl) FindAll() ([]model.Book, error) {
	return bookService.bookRepo.FindAll()
}

func (bookService *bookServiceImpl) FindById(id int) (*model.Book, error) {
	return bookService.bookRepo.FindById(id)
}
