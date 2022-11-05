package route

import (
	controller "github.com/arvinpaundra/go-rent-bike/internal/controller/rest-http"
	"github.com/arvinpaundra/go-rent-bike/internal/repository/gormdb"
	"github.com/arvinpaundra/go-rent-bike/internal/usecase"
	"github.com/labstack/echo/v4"
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

	u := v1.Group("/customers")
	u.GET("", userController.HandlerFindAllUsers)
	u.GET("/:id", userController.HandlerFindUserById)
	u.GET("/:id/histories", userController.HandlerFindAllUserHistories)
	u.PUT("/:id", userController.HandlerUpdateUser)
	u.DELETE("/:id", userController.HandlerDeleteUser)

	// renter
	renterRepository := gormdb.NewRenterRepositoryGorm(db)
	renterUsecase := usecase.NewRenterUsecase(renterRepository)
	renterController := controller.NewRenterController(renterUsecase)

	r := v1.Group("/renters")
	r.POST("", renterController.HandlerCreateRenter)
	r.GET("", renterController.HandlerFindAllRenters)
	r.GET("/:id", renterController.HandlerFindRenterById)
	r.PUT("/:id", renterController.HandlerUpdateRenter)
	r.DELETE("/:id", renterController.HandlerDeleteRenter)

	// category
	categoryRepository := gormdb.NewCategoryRepositoryGorm(db)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepository)
	categoryController := controller.NewCategoryController(categoryUsecase)

	c := v1.Group("/categories")
	c.POST("", categoryController.HandlerCreateCategory)
	c.GET("", categoryController.HandlerFindAllCategories)
	c.GET("/:id", categoryController.HandlerFindCategoryById)
	c.PUT("/:id", categoryController.HandlerUpdateCategory)
	c.DELETE("/:id", categoryController.HandlerDeleteCategory)

	// bike
	bikeRepository := gormdb.NewBikeRepositoryGorm(db)
	bikeUsecase := usecase.NewBikeUsecase(bikeRepository)
	bikeController := controller.NewBikeController(bikeUsecase)

	b := v1.Group("/bikes")
	b.POST("", bikeController.HandlerAddNewBike)
	b.GET("", bikeController.HandlerFindAllBikes)
	b.GET("/renters/:renterId", bikeController.HandlerFindBikesByRenter)
	b.GET("/categories/:categoryId", bikeController.HandlerFindBikesByCategory)
	b.GET("/:id", bikeController.HandlerFindByIdBike)
	b.PUT("/:id", bikeController.HandlerUpdateBike)
	b.DELETE("/:id", bikeController.HandlerDeleteBike)
	b.POST("/:id/reviews", bikeController.HandlerCreateNewBikeReview)

	// order
	orderRepository := gormdb.NewOrderRepository(db)
	orderUsecase := usecase.NeworderUsecase(orderRepository)
	orderController := controller.NewOrderController(orderUsecase)

	o := v1.Group("/orders")
	o.POST("", orderController.HandlerCreateNewOrder)
	o.GET("/customers/:userId", orderController.HandlerFindAllOrdersUser)
	o.GET("/:orderId/customers", orderController.HandlerFindByIdOrderUser)
	o.GET("/:orderId/return", orderController.HandlerReturnBike)
}
