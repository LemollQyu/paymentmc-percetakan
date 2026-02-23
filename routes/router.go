package routes

import (
	"paymentmc/cmd/app/handler"
	"paymentmc/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, paymentHandler handler.PaymentHandler, jwtSecret string) {
	router.Use(middleware.RequestLogger())

	// middleware hak akses user (harus login)
	authMiddleware := middleware.AuthMiddleware(jwtSecret)
	user := router.Group("/api/v1/i")
	user.Use(authMiddleware)

	// middleware hak akses admin
	admin := router.Group("/api/v1/admin")
	admin.Use(middleware.AdminMiddleware(jwtSecret, "admin"))

	// --- Admin only: /api/v1/admin ---
	admin.POST("/method-payment", paymentHandler.CreatePaymentMethod)
	admin.GET("/method-payment", paymentHandler.GetAllPaymentMethod)
	admin.DELETE("/method-payment/:methodID", paymentHandler.DeletePaymentMethod)
	admin.GET("/list-waiting-payment", paymentHandler.GetListWaitingPayment)
	admin.GET("/payments", paymentHandler.GetPayments)
	admin.POST("/payment/:code/approve", paymentHandler.ApprovePayment)
	admin.POST("/payment/:code/reject", paymentHandler.RejectPayment)

	// --- User (login): /api/v1/i ---
	user.GET("/my-payment", paymentHandler.GetMyPayment)
	user.GET("/code-payment/:code", paymentHandler.GetFullPaymentCodeByCode)
	user.GET("/payment/:orderID", paymentHandler.GetPaymentByOrderID)
	user.POST("/checkout", paymentHandler.Checkout)
	user.POST("/payment/:code", paymentHandler.Payment)
	user.GET("/cancelled/payment/:code", paymentHandler.CancelledPaymentAndOrder)
	user.POST("/payment-proof/:code", paymentHandler.ProofPayment)
}
