package routes

import (
	api "gin_mall/api/v1"
	"gin_mall/middleware"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "gin_mall/docs"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	// r.Use(func(c *gin.Context) {
	// 	defer func(c *gin.Context) {
	// 		err := recover()
	// 		if err != nil {
	// 			c.JSON(serializer.FailResponse(100, (err.(error)).Error()))
	// 		}
	// 	}(c)
	// 	c.Next()
	// })

	r.StaticFS("/static", http.Dir("./static"))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(200, "pong")
			// a := 1 - 1
			// _ = 0 / a
		})

		v1.GET("carousels", api.ListCarousels)
		v1.POST("upload", api.Upload)
		// 用户操作
		user := v1.Group("/user")
		{
			user.POST("register", api.UserRegister)
			user.POST("login", api.UserLogin)
		}

		authed := v1.Group("/")
		authed.Use(middleware.Auth())
		{
			authed.PUT("user", api.UserUpdate)
			authed.POST("avatar", api.UserUploadAvatar)
			authed.POST("sending-email", api.SendEmail)
			authed.POST("valid-email", api.ValidateEmail)
			authed.POST("money", api.ShowMoney)
		}

		product := v1.Group("/product")
		{
			product.GET("", api.ListProduct)
			product.GET(":id", api.GetProduct)
			product.GET("search", api.SearchProduct)

			product.Use(middleware.Auth())
			product.POST("", api.ProductCreate)

			// 	product.PUT("", 	api.ProductUpdate)
			// 	product.GET("", 	api.ProductList)
			// 	product.GET("/:id", api.ProductShow)
		}

		categories := v1.Group("/categories")
		{
			categories.GET("", api.ListCategories)
			// categories.GET(":id", api.GetCategory)

			categories.Use(middleware.Auth())
			// categories.POST("", api.CategoryCreate)
			// categories.PUT("", api.CategoryUpdate)
			// categories.DELETE("", api.CategoryDelete)
		}

		favorites := v1.Group("/favorites")
		favorites.Use(middleware.Auth())
		{
			favorites.GET("", api.ListFavorites)
			favorites.POST("", api.FavoriteCreate)
			favorites.DELETE("", api.FavoriteDelete)
		}

		addresses := v1.Group("/addresses")
		addresses.Use(middleware.Auth())
		{
			addresses.GET("", api.ListAddresses)
			addresses.GET(":id", api.GetAddresses)
			addresses.POST("", api.AddressCreate)
			addresses.PUT("", api.AddressUpdate)
			addresses.DELETE(":id", api.AddressDelete)
		}

		// 购物车
		cart := v1.Group("/cart")
		cart.Use(middleware.Auth())
		{
			cart.GET("", api.ListCart)
			// 添加商品到购物车
			cart.POST("", api.AddCart)
			// 删除购物车中的商品
			cart.DELETE("", api.DeleteCart)
			// 更新购物车中的商品
			cart.PUT("", api.UpdateCart)
		}

		// 订单
		orders := v1.Group("/orders")
		orders.Use(middleware.Auth())
		{
			orders.GET("", api.ListOrders)
			orders.GET("/:id", api.GetOrder)
			orders.POST("", api.OrderCreate)
			orders.PUT(":id", api.OrderUpdate)
			orders.DELETE(":id", api.OrderDelete)
		}

		// 支付功能
		pay := v1.Group("/pay")
		pay.Use(middleware.Auth())
		{
			pay.POST("", api.OrderPay)
		}

	}
	return r
}
