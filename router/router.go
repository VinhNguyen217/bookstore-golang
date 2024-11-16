package router

import (
	"book-store/controller"
	"book-store/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(di *do.Injector) (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")

	authController := controller.NewAuthController(di)
	userController := controller.NewUserController(di)
	bookController := controller.NewBookController(di)
	cartController := controller.NewCartController(di)
	billController := controller.NewBillController(di)

	authGroup := v1.Group("/auth")
	authGroup.POST("/login", authController.PasswordLogin)
	authGroup.POST("/sign-in", userController.Create)

	v1.GET("/books/:id", bookController.FindById)
	v1.Use(middlewares.Auth(di)).POST("/books", bookController.Create)
	v1.GET("/books", bookController.FindAll)
	v1.PUT("/books/:id", bookController.Update)
	v1.DELETE("/books/:id", bookController.Delete)

	v1.GET("/users/:id", userController.FindById)
	v1.GET("/users", userController.FindAll)
	v1.PUT("/users/:id", userController.Update)
	v1.DELETE("/users/:id", userController.Delete)

	v1.POST("/carts", cartController.Create) // create cart

	v1.POST("/bills", billController.Create)

	return r, nil
}
