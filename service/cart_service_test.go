package service

import (
	"book-store/dto/request"
	"book-store/dto/response"
	mockRepo "book-store/mock/book-store/repository"
	"book-store/model"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
func TestCreateCart(t *testing.T) {
	mockCartRepo := mockRepo.MockCartRepository{}
	mockBookRepo := mockRepo.MockBookRepository{}
	mockUserRepo := mockRepo.MockUserRepository{}

	mockBookRepo.
		On("FindById", mock.Anything, 1).
		Return(&model.Book{ID: 1, Quantity: 30}, nil)
	mockBookRepo.
		On("FindById", mock.Anything, 2).
		Return(&model.Book{ID: 2, Quantity: 30}, nil)
	mockBookRepo.
		On("FindById", mock.Anything, 3).
		Return(nil, gorm.ErrRecordNotFound)
	mockUserRepo.
		On("FindById", mock.Anything, 1).
		Return(&model.User{ID: 1, Name: "test"}, nil)
	mockUserRepo.
		On("FindById", mock.Anything, 2).
		Return(nil, gorm.ErrRecordNotFound)
	mockCartRepo.
		On("FindByUserIdAndBookId", mock.Anything, 1, 1).
		Return(&model.Cart{ID: 1, UserID: 1, BookID: 1})
	mockCartRepo.
		On("FindByUserIdAndBookId", mock.Anything, 1, 2).
		Return(nil, gorm.ErrRecordNotFound)

	cartService := &cartServiceImpl{
		cartRepo: &mockCartRepo,
		bookRepo: &mockBookRepo,
		userRepo: &mockUserRepo,
	}
	r := gin.Default()
	r.POST("/cart", func(c *gin.Context) {
		// Khởi tạo request giả lập
		var req request.CartItemRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gọi hàm Create trong cartService
		resp, err := cartService.Create(c, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Trả về kết quả dưới dạng JSON
		c.JSON(http.StatusOK, resp)
	})
}*/

func TestCreateCartItem(t *testing.T) {
	// Tạo gin engine giả lập
	r := gin.Default()

	mockCartRepo := mockRepo.MockCartRepository{}
	mockBookRepo := mockRepo.MockBookRepository{}
	mockUserRepo := mockRepo.MockUserRepository{}
	// Mock các repository

	// Khởi tạo cartService
	cartService := &cartServiceImpl{
		bookRepo: &mockBookRepo,
		cartRepo: &mockCartRepo,
		userRepo: &mockUserRepo,
	}

	// Tạo gin context giả lập
	r.POST("/carts", func(c *gin.Context) {
		// Lấy user_id từ context (mock trong test)
		c.Set("user_id", 1)

		// Tạo request CartItemRequest giả lập
		var req request.CartItemRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gọi hàm Create trong cartService
		resp, err := cartService.Create(c, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Trả về kết quả thành công
		c.JSON(http.StatusOK, resp)
	})

	// bookId = 0
	t.Run("BookId is 0", func(t *testing.T) {
		reqBody := `{"bookId": 0, "quantity": 2}`
		req, err := http.NewRequest("POST", "/carts", ginTestBody(reqBody))
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	// quantity = 0
	t.Run("Quantity is 0", func(t *testing.T) {
		reqBody := `{"bookId": 1, "quantity": 0}`
		req, err := http.NewRequest("POST", "/carts", ginTestBody(reqBody))
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	// Book not exist
	t.Run("Book not found", func(t *testing.T) {
		mockBookRepo.On("FindById", 1).
			Return(nil, errors.New("book not found"))

		reqBody := `{"bookId": 1, "quantity": 2}`
		req, err := http.NewRequest("POST", "/carts", ginTestBody(reqBody))
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	// Cart existed
	t.Run("Product already in cart", func(t *testing.T) {
		mockBookRepo.
			On("FindById", 1).
			Return(&model.Book{ID: 1, Quantity: 10, Price: 100}, nil)
		mockCartRepo.
			On("FindByUserIdAndBookId", 1, 1).
			Return(&model.Cart{UserID: 1, BookID: 1, Quantity: 2})

		reqBody := `{"bookId": 1, "quantity": 2}`
		req, err := http.NewRequest("POST", "/carts", ginTestBody(reqBody))
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	// Check success
	t.Run("Successfully added to cart", func(t *testing.T) {
		mockBookRepo.On("FindById", 1).Return(&model.Book{ID: 1, Quantity: 10, Price: 100}, nil)
		mockCartRepo.On("FindByUserIdAndBookId", 1, 1).Return(nil)
		mockCartRepo.On("Create", mock.Anything).Return(&model.Cart{
			UserID:   1,
			BookID:   1,
			Quantity: 2,
			Price:    200,
		}, nil)

		reqBody := `{"bookId": 1, "quantity": 2}`
		req, err := http.NewRequest("POST", "/carts", ginTestBody(reqBody))
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		// Kiểm tra response có đúng không
		var resp response.CartRes
		err = json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, resp.BookID, 1)
		assert.Equal(t, resp.Quantity, 2)
	})
}

func ginTestBody(data string) *bytes.Buffer {
	return bytes.NewBufferString(data)
}
