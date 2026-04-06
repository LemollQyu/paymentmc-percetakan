package handler

import (
	"net/http"
	"paymentmc/models"

	"github.com/gin-gonic/gin"
)

func (h *PaymentHandler) SendNotificationOrder(c *gin.Context) {
	var param models.NotificationOrderRequest

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := h.PaymentUsecase.SendNotificationOrder(c.Request.Context(), param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "notification sent"})
}
