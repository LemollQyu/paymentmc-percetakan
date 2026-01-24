package routes

import (
	"paymentmc/cmd/app/handler"
	"paymentmc/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, paymentHandler handler.PaymentHandler) {
	// cek request logging
	router.Use(middleware.RequestLogger())

	groupV1 := router.Group("/api/v1")

	// admin
	groupV1.POST("/method-payment", paymentHandler.CreatePaymentMethod)
	groupV1.GET("/method-payment", paymentHandler.GetAllPaymentMethod)
	groupV1.DELETE("/method-payment/:methodID", paymentHandler.DeletePaymentMethod)

	groupV1.GET("/code-payment/:code", paymentHandler.GetFullPaymentCodeByCode)

	groupV1.POST("/checkout", paymentHandler.Checkout)
	groupV1.POST("/payment/:code", paymentHandler.Payment)

	groupV1.GET("/cancelled/payment/:code", paymentHandler.CancelledPaymentAndOrder)

}
