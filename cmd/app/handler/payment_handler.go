package handler

import (
	"net/http"
	"paymentmc/infrastructure/log"
	"paymentmc/models"
	"paymentmc/utils"

	"github.com/gin-gonic/gin"
)

func (h *PaymentHandler) Checkout(c *gin.Context) {

	// param code order
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "order code is required",
		})
		return
	}

	// field dari request checkout
	var param models.RequestPayment

	err := c.ShouldBind(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Invalid param",
		})

		return
	}

	// ambil data order
	dataOrder, err := h.PaymentUsecase.GetOrderByCode(c.Request.Context(), code)
	if err != nil {
		switch err {
		case utils.OrderExpiredOrNotFound:
			c.JSON(http.StatusBadRequest, gin.H{
				"error_message": err.Error(),
			})
		default:
			log.Logger.Error("paymentHandler: GetOrderByCode ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"error_message": "internal server error",
			})
		}
		return
	}

	// request payment
	request := &models.RequestPayment{
		OrderID:       dataOrder.ID,
		OrderCode:     dataOrder.OrderCode.Code,
		UserID:        dataOrder.UserID,
		Amount:        float64(dataOrder.TotalPriceSnapshot),
		PaymentMethod: param.PaymentMethod,
	}

	// validasi method ada tidak
	dataMethod, err := h.PaymentUsecase.GetPaymentMethod(c.Request.Context(), param.PaymentMethod)
	if err != nil {
		log.Logger.Error("h.PaymentUsecase.GetPaymentMethod")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_message": "Kesalahan dari system",
		})
		return
	}

	if dataMethod == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Metode pembayaran tidak ada",
		})
		return
	}

	resDataPayment, err := h.PaymentUsecase.Checkout(c.Request.Context(), code, *request)

	if err != nil {
		switch err {
		case utils.OrderNotFound, utils.ErrStatusNotCreated, utils.OrderExpiredOrNotFound, utils.ErrStatusNotCompatible:
			c.JSON(http.StatusBadRequest, gin.H{
				"error_message": err.Error(),
			})
		default:
			log.Logger.Error("paymentHandler: Checkout", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"error_message": "internal server error",
			})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": resDataPayment,
	})
}

// handler cancelled payment and order
func (h *PaymentHandler) CancelledPaymentAndOrder(c *gin.Context) {
	// param code payment
	codePayment := c.Param("code")
	if codePayment == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "code tidak boleh kosong",
		})
		return
	}

	err := h.PaymentUsecase.CancelledPaymentAndOrder(c.Request.Context(), codePayment)

	if err != nil {
		switch err {
		case utils.ErrStatusPaymentShouldPending, utils.ErrPaymentCodeNotFound, utils.ErrStatusOrderShouldWaitingPayment, utils.ErrStatusNotCompatible:
			c.JSON(http.StatusBadRequest, gin.H{
				"error_message": err.Error(),
			})
		default:
			log.Logger.Error("paymentHandler: CancelledPaymentAndOrder", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"error_message": "internal server error",
			})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})

}

// handler bayar order
// flownya, kirim bukti pembayaran, dan kirim ke admin buat di acc pembayarannya
func (h *PaymentHandler) Payment(c *gin.Context) {
	codePayment := c.Param("code")
	if codePayment == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "code tidak boleh kosong",
		})
		return
	}

	// usecase flow payment

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}
