package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *PaymentHandler) GetListWaitingPayment(c *gin.Context) {
	data, err := h.PaymentUsecase.GetListWaitingPayment(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_message": "Kesalahan internal system",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    data,
	})
}

// GetMyPayment untuk user yang login: menampilkan data checkout (payment Pending, order Waiting_payment).
// Data sama dengan list-waiting-payment, hanya milik user ini.
func (h *PaymentHandler) GetMyPayment(c *gin.Context) {
	userIDAny, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error_message": "unauthorized",
		})
		return
	}

	userID, ok := userIDAny.(int64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error_message": "invalid token",
		})
		return
	}

	data, err := h.PaymentUsecase.GetMyPayment(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_message": "Kesalahan internal system",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    data,
	})
}
