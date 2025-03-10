package routes

import (
	authcontroller "example.com/go-api/controller/auth"
	productcontroller "example.com/go-api/controller/product"
	productcategorycontroller "example.com/go-api/controller/product_category"
	"example.com/go-api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/register", authcontroller.RegisterUser)
	server.POST("/login", authcontroller.LoginUser)
	server.POST("/forget-password", authcontroller.ForgetPassword)
	server.POST("/set-new-password", authcontroller.SetNewPassword)

	categoryRoutes := server.Group("/categoies")
	categoryRoutes.Use(middleware.VerifyToken)
	// productRoutes := server.Group("/products")
	categoryRoutes.POST("/list-all-product", productcontroller.ListAllProduct)
	categoryRoutes.POST("/list-all-product-category", productcategorycontroller.ListAllProductCategory)
}
