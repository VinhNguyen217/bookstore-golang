package service

import (
	"book-store/dto/request"
	"book-store/model"
	"book-store/repository"
	"errors"
	"github.com/samber/do"
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
	date, err := time.Parse("2006-01-02", req.PublishDate)
	if err != nil {
		return nil, err
	}
	book := &model.Book{
		Name:        req.Name,
		Photo:       req.Photo,
		Quantity:    req.Quantity,
		Price:       req.Price,
		PublishDate: date,
		Description: req.Description,
	}
	return bookService.bookRepo.Create(book)
}

func (bookService *bookServiceImpl) Update(req *request.BookRequest, id int) (*model.Book, error) {
	_, errId := bookService.bookRepo.FindById(id)
	if errId != nil {
		return nil, errors.New("Sách này không tồn tại")
	}

	var date time.Time
	if req.PublishDate != "" {
		date, _ = time.Parse("2006-01-02", req.PublishDate)
	}
	book := &model.Book{
		ID:          id,
		Name:        req.Name,
		Photo:       req.Photo,
		Quantity:    req.Quantity,
		Price:       req.Price,
		PublishDate: date,
		Description: req.Description,
	}
	err := bookService.bookRepo.Update(book)
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
