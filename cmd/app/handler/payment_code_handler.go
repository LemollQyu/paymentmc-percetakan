package handler

import (
	"net/http"
	"paymentmc/infrastructure/log"
	"paymentmc/utils"

	"github.com/gin-gonic/gin"
)

func (h *PaymentHandler) GetFullPaymentCodeByCode(c *gin.Context) {

	code := c.Param("code")
	if code == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "code tidak boleh kosong",
		})
		return
	}

	data, err := h.PaymentUsecase.GetFullPaymentCodeByCode(c.Request.Context(), code)
	if err != nil {
		switch err {
		case utils.ErrPaymentCodeNotFound:
			c.JSON(http.StatusBadRequest, gin.H{
				"error_message": err.Error(),
			})
		default:
			log.Logger.Error("paymentHandler: etFullPaymentCodeByCode ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"error_message": "internal server error",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Success",
		"data":    data,
	})

}
