package main

import (
	"github.com/MikeMwita/fast_foods/database"
	"github.com/MikeMwita/fast_foods/middleware"
	"github.com/MikeMwita/fast_foods/routes"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	var foodCollection *mongo.Collection = database.OpenCollection("database.Client", food)

	router := gin.New()
	router.Use(gin.Logger())
	router.UserRoutes(router)
	router.Use(middleware.Authetication())

	//initializing all our routes
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemsRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}
