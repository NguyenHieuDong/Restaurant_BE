package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/foods", controller.CreateFood())
	incomingRoutes.GET("/foods", controller.GetFoods())
	incomingRoutes.GET("/foods/:foodId", controller.GetFood())
	incomingRoutes.PATCH("/foods", controller.UpdateFood())

}
