package main

import (
	"os"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/shashank/ecommerce-yt/controllers"
	"github.com/shashank/ecommerce-yt/database"
	"github.com/shashank/ecommerce-yt/middleware"
	"github.com/shashank/ecommerce-yt/routes"
	
)


func main(){
	port := os.Getenv("PORT")
	IF port == ""{
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "product"), database.UserData(databse.client, "users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFRomCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}	
