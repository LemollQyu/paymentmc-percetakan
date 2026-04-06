package handler

import (
	"net/http"
	"paymentmc/models"
	"paymentmc/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// handler pengajuan refund dari admin
func (h *PaymentHandler) SubmitRefund(c *gin.Context) {

	code := c.Param("code")
	if code == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "code tidak boleh kosong",
		})
		return
	}

	// param

	var param models.RequestRejectedPayment

	err := c.ShouldBind(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Invalid param",
		})

		return
	}

	// ambil data payment
	dataPayment, err := h.PaymentUsecase.GetFullPaymentCodeByCode(c.Request.Context(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_message": "Kesalahan internal dari system",
		})
		return
	}

	// data code ada tidak
	if dataPayment == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "code tidak ditemukan",
		})
		return
	}

	// validasi status payment harus cancelled
	if dataPayment.Payment.Status != utils.StatusPaymentCancelled {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "status payment harus cancelled",
		})

		return
	}

	request := models.RequestRejectedPayment{
		PaymentID:   dataPayment.PaymentID,
		UserID:      dataPayment.Payment.UserID,
		OrderCode:   param.OrderCode,
		Amount:      int64(dataPayment.Payment.Amount),
		PaymentCode: dataPayment.Code,
		OrderName:   param.OrderName,
		AdminNote:   param.AdminNote,
	}

	err = h.PaymentUsecase.SubmitRefund(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_message": "Kesalahan internal dari system",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

// my refund
func (h *PaymentHandler) GetMyRefund(c *gin.Context) {
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

	data, err := h.PaymentUsecase.GetMyRefund(c.Request.Context(), userID)
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

// all nrefund
func (h *PaymentHandler) GetAllRefunds(c *gin.Context) {

	data, err := h.PaymentUsecase.GetAllRefunds(c.Request.Context())
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

// get detail refund
func (h *PaymentHandler) GetRefundByID(c *gin.Context) {
	rejectID := c.Param("rejectID")
	id, err := strconv.ParseInt(rejectID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "invalid service id",
		})
		return
	}

	data, err := h.PaymentUsecase.GetRefundDetailByID(c.Request.Context(), id)
	if err != nil {
		switch err {
		case utils.ErrRejectPaymentNotFound:
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
		"data":    data,
	})
}

func (h *PaymentHandler) SendNumberRakening(c *gin.Context) {
	rejectID := c.Param("rejectID")
	id, err := strconv.ParseInt(rejectID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "invalid service id",
		})
		return
	}

	var param models.RequestRefund

	err = c.ShouldBind(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Invalid param",
		})

		return
	}

	err = h.PaymentUsecase.SendNumberRekening(c.Request.Context(), id, param)
	if err != nil {
		switch err {
		case utils.ErrRejectPaymentNotFound, utils.ErrDataRefundIsFound:
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

func (h *PaymentHandler) ProofRefund(c *gin.Context) {

	refundID := c.Param("refundID")
	id, err := strconv.ParseInt(refundID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "invalid refund id",
		})
		return
	}

	var param models.RequestRefundProof
	if err = c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Invalid param",
		})
		return
	}

	proofFile, err := c.FormFile("bukti")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Bukti transfer wajib diupload",
		})
		return
	}

	err = h.PaymentUsecase.ProofRefund(c.Request.Context(), proofFile, param, id)
	if err != nil {
		if err.Error() == "refund tidak ditemukan" {
			c.JSON(http.StatusNotFound, gin.H{
				"error_message": err.Error(),
			})
			return
		}
		if err.Error() == "refund sudah diproses sebelumnya" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error_message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_message": "Gagal mengupload bukti refund",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Bukti refund berhasil diupload",
	})
}

func (h *PaymentHandler) ApproveRefund(c *gin.Context) {
	refundID := c.Param("refundID")
	id, err := strconv.ParseInt(refundID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "invalid service id",
		})
		return
	}

	data, err := h.PaymentUsecase.GetRefundDetailByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_message": "Kesalahan internal system",
		})
		return
	}

	if data == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "data refund tidak ditemukan",
		})
		return
	}

	err = h.PaymentUsecase.ApproveRefund(c.Request.Context(), id)
	if err != nil {
		switch err {
		case utils.ErrStatusNotCompatible:
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
		"message": "Berhasil diterima",
	})
}
