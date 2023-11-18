package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/orderItems", controller.CreateOrderItems())
	incomingRoutes.GET("/orderItems", controller.GetOrderItems())
	incomingRoutes.GET("/orderItems-order/:orderItemId", controller.GetOrderItemsByOrder())
	incomingRoutes.GET("/orderItems/:orderItemId", controller.GetOrderItem())
	incomingRoutes.PATCH("/orderItems", controller.UpdateOrderItems())

}
