package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/invoices", controller.CreateInvoices())
	incomingRoutes.GET("/invoices", controller.GetInvoices())
	incomingRoutes.GET("/invoices/:invoiceId", controller.GetInvoice())
	incomingRoutes.PATCH("/invoices", controller.UpdateInvoice())

}
