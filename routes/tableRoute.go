package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/tables", controller.CreateTables())
	incomingRoutes.GET("/tables", controller.GetTables())
	incomingRoutes.GET("/tables/:tableId", controller.GetTable())
	incomingRoutes.PATCH("/tables", controller.UpdateTables())

}
