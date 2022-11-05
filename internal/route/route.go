package route

import (
	"github.com/arvinpaundra/go-rent-bike/configs"
	controller "github.com/arvinpaundra/go-rent-bike/internal/controller/rest-http"
	mddlwrs "github.com/arvinpaundra/go-rent-bike/internal/middlewares"
	"github.com/arvinpaundra/go-rent-bike/internal/repository/gormdb"
	"github.com/arvinpaundra/go-rent-bike/internal/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func New(db *gorm.DB, e *echo.Echo) {
	v1 := e.Group("/api/v1")

	// midtrans notif
	paymentGatewayUsecase := usecase.NewPaymentGatewayUsecase()
	paymentGatewayController := controller.NewMidtransNotificationController(paymentGatewayUsecase)

	v1.POST("/webhook/midtrans", paymentGatewayController.HandlerNotification)

	//	user auth
	userRepository := gormdb.NewUserRepositoryGorm(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	// customer
	auth := v1.Group("/auth")
	auth.POST("/register", userController.HandlerRegister)
	auth.POST("/login", userController.HandlerLogin)

	u := v1.Group("/customers", middleware.JWT([]byte(configs.Cfg.JWTSecret)))
	u.GET("", userController.HandlerFindAllUsers)
	u.GET("/:id/histories", userController.HandlerFindAllUserHistories)
	u.GET("/:id/orders", userController.HandlerFindAllOrdersUser)
	u.GET("/:id/orders/:orderId", userController.HandlerFindByIdOrderUser)
	u.GET("/:id", userController.HandlerFindUserById)
	u.PUT("/:id", userController.HandlerUpdateUser)
	u.DELETE("/:id", userController.HandlerDeleteUser)

	// renter
	renterRepository := gormdb.NewRenterRepositoryGorm(db)
	renterUsecase := usecase.NewRenterUsecase(renterRepository)
	renterController := controller.NewRenterController(renterUsecase)

	r := v1.Group("/renters")
	r.POST("", renterController.HandlerCreateRenter, middleware.JWT([]byte(configs.Cfg.JWTSecret)))
	r.GET("", renterController.HandlerFindAllRenters)
	r.GET("/:id", renterController.HandlerFindRenterById)
	r.PUT("/:id", renterController.HandlerUpdateRenter, middleware.JWT([]byte(configs.Cfg.JWTSecret)), mddlwrs.CheckIsRenter)
	r.DELETE("/:id", renterController.HandlerDeleteRenter, middleware.JWT([]byte(configs.Cfg.JWTSecret)), mddlwrs.CheckIsRenter)

	// category
	categoryRepository := gormdb.NewCategoryRepositoryGorm(db)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepository)
	categoryController := controller.NewCategoryController(categoryUsecase)

	c := v1.Group("/categories")
	c.POST("", categoryController.HandlerCreateCategory, middleware.JWT([]byte(configs.Cfg.JWTSecret)), mddlwrs.CheckIsRenter)
	c.GET("", categoryController.HandlerFindAllCategories)
	c.GET("/:id", categoryController.HandlerFindCategoryById)
	c.PUT("/:id", categoryController.HandlerUpdateCategory, middleware.JWT([]byte(configs.Cfg.JWTSecret)), mddlwrs.CheckIsRenter)
	c.DELETE("/:id", categoryController.HandlerDeleteCategory, middleware.JWT([]byte(configs.Cfg.JWTSecret)), mddlwrs.CheckIsRenter)

	// bike
	bikeRepository := gormdb.NewBikeRepositoryGorm(db)
	bikeUsecase := usecase.NewBikeUsecase(bikeRepository)
	bikeController := controller.NewBikeController(bikeUsecase)

	b := v1.Group("/bikes")
	b.POST("", bikeController.HandlerAddNewBike, middleware.JWT([]byte(configs.Cfg.JWTSecret)), mddlwrs.CheckIsRenter)
	b.GET("", bikeController.HandlerFindAllBikes)
	b.GET("/renters/:renterId", bikeController.HandlerFindBikesByRenter)
	b.GET("/categories/:categoryId", bikeController.HandlerFindBikesByCategory)
	b.GET("/:id", bikeController.HandlerFindByIdBike)
	b.PUT("/:id", bikeController.HandlerUpdateBike, middleware.JWT([]byte(configs.Cfg.JWTSecret)), mddlwrs.CheckIsRenter)
	b.DELETE("/:id", bikeController.HandlerDeleteBike, middleware.JWT([]byte(configs.Cfg.JWTSecret)), mddlwrs.CheckIsRenter)
	b.POST("/:id/reviews", bikeController.HandlerCreateNewBikeReview, middleware.JWT([]byte(configs.Cfg.JWTSecret)))

	// order
	orderRepository := gormdb.NewOrderRepository(db)
	orderUsecase := usecase.NeworderUsecase(orderRepository)
	orderController := controller.NewOrderController(orderUsecase)

	o := v1.Group("/orders", middleware.JWT([]byte(configs.Cfg.JWTSecret)))
	o.POST("", orderController.HandlerCreateNewOrder)
	o.GET("/:id/return", orderController.HandlerReturnBike)
}
