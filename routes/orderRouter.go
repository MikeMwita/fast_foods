package routes

func orderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders", controller.GetOrders())
	incomingRoutes.GET("/orders/:order", controller.GetOrder())
	incomingRoutes.POST("/orders", controller.Createorder())
	incomingRoutes.PATCH("/orders/:orderID", controller.UpdateOrder())
}
