package routes

import (
	controller "github.com/MikeMwita/fast_foods/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemsRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controller.GetOrderItems())
	incomingRoutes.GET("/tables/:tableId", controller.GetTables())
	incomingRoutes.POST("/tables", controller.CreateTable())
	incomingRoutes.PATCH("/tables/:tableID", controller.UpdateTable())
}
