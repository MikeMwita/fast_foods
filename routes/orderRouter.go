package routes

import (
	controller "github.com/MikeMwita/fast_foods/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders", controller.GetOrders())
	incomingRoutes.GET("/orders/:order", controller.GetOrder())
	incomingRoutes.POST("/orders", controller.Createorder())
	incomingRoutes.PATCH("/orders/:orderID", controller.UpdateOrder())
}
