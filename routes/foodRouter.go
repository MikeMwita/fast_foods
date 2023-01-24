package routes

import (
	controller "......"
	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/food", controller.GetFoods())
	incomingRoutes.GET("/food/:foodId", controller.GetFood())
	incomingRoutes.POST("/food", controller.CreateFood())
	incomingRoutes.PATCH("/foods/:foodId", controller.UpdateFood())
}
