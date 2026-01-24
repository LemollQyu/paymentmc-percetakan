package usecase

import (
	"context"
	"mime/multipart"
	"paymentmc/infrastructure/log"
	"paymentmc/models"
	"paymentmc/utils"
	"strings"

	"github.com/sirupsen/logrus"
)

func (uc *PaymentUsecase) CreatePaymentMethod(ctx context.Context, param models.PaymentMethodRequest, file map[string]*multipart.FileHeader) error {
	// payment method apakah sudah ada
	dataPaymentMethod, err := uc.PaymentService.GetPaymentMethod(ctx, param.PaymentMethod)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"method": param.PaymentMethod,
			"error":  err.Error(),
		}).Info("uc.PaymentService.GetPaymentMethod")
		return err
	}

	// kalo ada
	if dataPaymentMethod != nil {
		return utils.PaymentMethodExist
	}

	// upload icon
	urlIcon, err := uc.StorageService.UploadIconPaymentMethod(ctx, file["icon"])
	if err != nil {
		log.Logger.Error("uc.StorageService.UploadIconPaymentMethod")
		return err
	}

	if strings.EqualFold(param.PaymentMethod, "qris") {

		urlCode, err := uc.StorageService.UploadCodeQris(ctx, file["qris"])
		if err != nil {
			log.Logger.Error("uc.StorageService.UploadCodeQris")
			return err
		}

		req := models.PaymentMethodRequest{
			PaymentMethod: param.PaymentMethod,
			UrlIcon:       urlIcon,
			UrlCode:       urlCode,
			NumberPayment: "",
		}

		err = uc.PaymentService.CreatePaymentMethod(ctx, req)
		if err != nil {
			log.Logger.WithFields(logrus.Fields{
				"req":   req,
				"error": err.Error(),
			}).Info("uc.PaymentService.CreatePaymentMethod")
			return err
		}

		return nil

	}

	req := models.PaymentMethodRequest{
		PaymentMethod: param.PaymentMethod,
		UrlIcon:       urlIcon,
		UrlCode:       "",
		NumberPayment: param.NumberPayment,
	}

	err = uc.PaymentService.CreatePaymentMethod(ctx, req)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"req":   req,
			"error": err.Error(),
		}).Info("uc.PaymentService.CreatePaymentMethod")
		return err
	}

	return nil
}

func (uc *PaymentUsecase) GetAllPaymentMethod(ctx context.Context) (*[]models.PaymentMethod, error) {
	dataPaymentMethod, err := uc.PaymentService.GetAllPaymentMethod(ctx)
	if err != nil {
		log.Logger.Error("uc.PaymentService.GetAllPaymentMethod")
		return nil, err
	}

	return dataPaymentMethod, nil
}

func (uc *PaymentUsecase) DeletePaymentMethod(ctx context.Context, methodID int64) error {
	// get payment method
	dataPaymentMethod, err := uc.PaymentService.GetPaymentMethodByID(ctx, methodID)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"id":    methodID,
			"error": err.Error(),
		}).Error("uc.PaymentService.GetPaymentMethodByID")
		return err
	}

	if dataPaymentMethod == nil {
		return utils.PaymentMethodNotExist
	}

	// delete file icon storage
	err = uc.StorageService.DeleteFile(ctx, dataPaymentMethod.UrlIcon)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"url":   dataPaymentMethod.UrlIcon,
			"error": err.Error(),
		}).Error("uc.StorageService.DeleteFileIcon")
	}

	// kondisi jika yang didelete payment method qris
	if dataPaymentMethod.PaymentMethod == "Qris" {
		// delete file code qris storage
		err = uc.StorageService.DeleteFile(ctx, dataPaymentMethod.UrlCode)
		if err != nil {
			log.Logger.WithFields(logrus.Fields{
				"url":   dataPaymentMethod.UrlCode,
				"error": err.Error(),
			}).Error("uc.StorageService.DeleteFileQris")

		}
	}

	// delete di database
	err = uc.PaymentService.DeletePaymentMethod(ctx, methodID)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"id":    methodID,
			"error": err.Error(),
		}).Error("uc.PaymentService.DeletePaymentMethod")
		return err
	}

	return nil

}

// get payment method by name
func (uc *PaymentUsecase) GetPaymentMethod(ctx context.Context, method string) (*models.PaymentMethod, error) {
	dataPaymentMethod, err := uc.PaymentService.GetPaymentMethod(ctx, method)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"method": method,
			"error":  err.Error(),
		}).Error("uc.PaymentService.GetPaymentMethod")
		return nil, err
	}

	return dataPaymentMethod, err
}
