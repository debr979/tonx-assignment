package routes

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	docs "tonx-assignment/docs"
	"tonx-assignment/internal/app/controllers"
	"tonx-assignment/internal/middlewares"
)

func Router() *gin.Engine {
	runMode := os.Getenv(`RUN_MODE`)
	gin.SetMode(runMode)
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	middlewares.UseMiddlewares(router)

	api := router.Group("/api")

	users := api.Group("/users")
	{
		// create new account
		users.POST("/user", controllers.UserController.Register)
		// modify password
		users.PATCH("/user", middlewares.UserAuthorization, controllers.UserController.ChangePassword)
		// delete account
		users.DELETE("/user", middlewares.UserAuthorization, controllers.UserController.DeleteAccount)
		// login
		users.POST("/login", controllers.UserController.Login)

	}

	coupons := api.Group("/coupons")
	{
		// for user
		// get available coupons
		coupons.GET("/coupons", controllers.CouponController.GetCoupons)
		// reserve coupons
		coupons.POST("/reserve", middlewares.UserAuthorization, controllers.CouponController.Reserve)
		// grab current available coupon
		coupons.POST("/grab", middlewares.UserAuthorization, controllers.CouponController.Grab)
		// use anyone coupon
		coupons.POST("/useCoupon", middlewares.UserAuthorization, controllers.CouponController.UseCoupon)
	}

	authorization := api.Group("/auth")
	{
		// refresh access token
		authorization.POST("/refreshToken", controllers.AuthController.RefreshToken)
	}

	mgr := api.Group("/mgr")
	{
		// for manager
		mgr.POST("/login", controllers.MgrController.Login)

		mgr.POST("/coupons", middlewares.MgrAuthorization, controllers.CouponController.AddCoupon)
		mgr.DELETE("/coupons", middlewares.MgrAuthorization, controllers.CouponController.DeleteCoupon)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
