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
	authGroup.POST("/sign-in", authController.PasswordLogin)
	authGroup.POST("/sign-up", userController.CreateUser)

	v1.Use(middlewares.Auth(di)).
		GET("/users/my-info", middlewares.Authorization(di), userController.GetMyInfo)
	v1.Use(middlewares.Auth(di)).
		GET("/users", middlewares.Authorization(di), userController.FindAll)
	v1.Use(middlewares.Auth(di)).
		PUT("/users", middlewares.Authorization(di), userController.UpdateUser)

	bookGroup := v1.Group("/books")
	bookGroup.Use(middlewares.Auth(di))
	bookGroup.POST("", middlewares.Authorization(di), bookController.Create)
	bookGroup.PUT("/:id", middlewares.Authorization(di), bookController.Update)
	bookGroup.DELETE("/:id", middlewares.Authorization(di), bookController.Delete)
	bookGroup.GET("/:id", bookController.FindById)
	bookGroup.GET("", bookController.FindAll)

	cartGroup := v1.Group("/carts")
	cartGroup.Use(middlewares.Auth(di))
	cartGroup.POST("", middlewares.Authorization(di), cartController.Create)
	cartGroup.PUT("", middlewares.Authorization(di), cartController.Update)
	cartGroup.GET("", middlewares.Authorization(di), cartController.GetCartsByUserId)
	cartGroup.DELETE("/:id", middlewares.Authorization(di), cartController.DeleteCartById)

	billGroup := v1.Group("/bills")
	billGroup.Use(middlewares.Auth(di))
	billGroup.POST("", middlewares.Authorization(di), billController.Create)
	billGroup.PUT("/cancel/:id", middlewares.Authorization(di), billController.CancelBill)
	billGroup.PUT("/update-status/:id", middlewares.Authorization(di), billController.UpdateStatusBill)
	billGroup.GET("/user", middlewares.Authorization(di), billController.FindByUser)
	billGroup.GET("", middlewares.Authorization(di), billController.FindAll)

	return r, nil
}
