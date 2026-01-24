package handler

import (
	"mime/multipart"
	"net/http"
	"paymentmc/infrastructure/log"
	"paymentmc/models"
	"paymentmc/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *PaymentHandler) CreatePaymentMethod(c *gin.Context) {
	var param models.PaymentMethodRequest
	if err := c.ShouldBind(&param); err != nil {
		log.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Invalid param",
		})
		return
	}

	// ambil icon (WAJIB)
	iconFile, err := c.FormFile("icon")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "icon wajib diupload",
		})
		return
	}

	files := map[string]*multipart.FileHeader{
		"icon": iconFile,
	}

	// jika QRIS, code WAJIB
	if strings.EqualFold(param.PaymentMethod, "qris") {
		codeFile, err := c.FormFile("qris")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error_message": "QRIS code wajib diupload",
			})
			return
		}
		files["qris"] = codeFile
	}

	err = h.PaymentUsecase.CreatePaymentMethod(c.Request.Context(), param, files)
	if err != nil {
		switch err {
		case utils.FileExtInvalid, utils.FileMaxSize, utils.FileRequired, utils.PaymentMethodExist:
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

	c.JSON(http.StatusCreated, gin.H{
		"message": "payment method berhasil dibuat",
	})
}

func (h *PaymentHandler) GetAllPaymentMethod(c *gin.Context) {
	dataPaymentMethod, err := h.PaymentUsecase.GetAllPaymentMethod(c.Request.Context())
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error_message": "Kesalahan dari system",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    dataPaymentMethod,
		"message": "success",
	})
}

func (h *PaymentHandler) DeletePaymentMethod(c *gin.Context) {

	methodID := c.Param("methodID")
	id, err := strconv.ParseInt(methodID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "invalid service id",
		})
		return
	}

	err = h.PaymentUsecase.DeletePaymentMethod(c.Request.Context(), id)
	if err != nil {
		switch err {
		case utils.PaymentMethodNotExist:
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
		"message": "berhasil menghapus payment method",
	})
}
