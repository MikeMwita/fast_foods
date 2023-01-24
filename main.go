package main

import "os"

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

//install go gin,mongodb+ getenv
