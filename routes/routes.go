package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shashank/ecommerce-yt/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp()),
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.productviewerAdmin())
	incomingRoutes.GET("/users/productview", controllers.SearchProduct())
	incomingRoutes.GET("/users/search")
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())

}
