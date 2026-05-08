package handler

import (
	"context"
	"net/http"
	"paymentmc/infrastructure/log"
	"paymentmc/models"
	"paymentmc/utils"
	"strconv"
	"time"

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

// handle bukti pembayaran
func (h *PaymentHandler) ProofPayment(c *gin.Context) {
	code := c.Param("code")

	if code == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "code tidak boleh kosong",
		})
		return
	}

	var param models.RequestBuktiPembayaran
	if err := c.ShouldBind(&param); err != nil {
		log.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Invalid param",
		})
		return
	}

	proofFile, err := c.FormFile("bukti")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Bukti pembayaran wajib diupload",
		})
		return
	}

	// sesudah
	uploadCtx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	uploadAt, err := h.PaymentUsecase.ProofPayment(uploadCtx, code, proofFile, param.Note)
	if err != nil {
		switch err {
		case utils.FileExtInvalid, utils.FileMaxSize, utils.FileRequired, utils.ErrPaymentCodeNotFound, utils.ErrStatusPaymentShouldPending, utils.ErrStatusOrderShouldWaitingPayment, utils.ErrPaymentPaid:
			c.JSON(http.StatusBadRequest, gin.H{
				"error_message": err.Error(),
			})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error_message": err.Error(),
			})

		}
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "Menunggu dikonfirmasi pihak percetakan",
		"upload":  uploadAt,
	})
}

// management payments
func (h *PaymentHandler) GetPayments(c *gin.Context) {

	status := c.Query("status")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	data, err := h.PaymentUsecase.GetPayments(
		c.Request.Context(),
		status,
		page,
		limit,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    data,
		"meta": gin.H{
			"page":  page,
			"limit": limit,
		},
	})
}

func (h *PaymentHandler) ApprovePayment(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "code tidak boleh kosong",
		})
		return
	}

	err := h.PaymentUsecase.ApprovePayment(c.Request.Context(), code)
	if err != nil {
		switch err {
		case utils.ErrPaymentCodeNotFound, utils.ErrPaymentShouldSuccess, utils.OrderNotFound, utils.ErrOrderShouldPaid:
			c.JSON(http.StatusBadRequest, gin.H{
				"error_message": err.Error(),
			})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error_message": err.Error(),
			})

		}
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *PaymentHandler) RejectPayment(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "code tidak boleh kosong",
		})
		return
	}

	err := h.PaymentUsecase.RejectPayment(c.Request.Context(), code)
	if err != nil {
		switch err {
		case utils.ErrPaymentCodeNotFound, utils.ErrPaymentShouldSuccess, utils.OrderNotFound, utils.ErrOrderShouldPaid:
			c.JSON(http.StatusBadRequest, gin.H{
				"error_message": err.Error(),
			})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error_message": err.Error(),
			})

		}
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *PaymentHandler) GetPaymentByOrderID(c *gin.Context) {
	id := c.Param("orderID")
	orderID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "invalid service id",
		})
		return
	}

	data, err := h.PaymentUsecase.GetPaymentByOrderID(c.Request.Context(), orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_message": "Kesalahan internal system",
		})
		return
	}

	if data == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "payment tidak ada",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    data,
	})
}

func (h *PaymentHandler) GetPaymentProof(c *gin.Context) {
	id := c.Param("paymentID")
	paymentID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "invalid service id",
		})
		return
	}

	data, err := h.PaymentUsecase.GetPaymentProofByID(c.Request.Context(), paymentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Kesalahan internal system",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    data,
	})
}
