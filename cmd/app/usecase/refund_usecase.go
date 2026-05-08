package usecase

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"paymentmc/infrastructure/log"
	"paymentmc/models"
	"paymentmc/utils"
	"time"

	"github.com/sirupsen/logrus"
)

func (uc *PaymentUsecase) SubmitRefund(ctx context.Context, param models.RequestRejectedPayment) error {
	err := uc.PaymentService.CreateSubmitRefund(ctx, param)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("uc.PaymentService.CreateSubmitRefund")
		return err
	}

	// kirim notifikasi ke usernya
	go func() {
		notifCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// ambil data payment by code
		dataPaymentCode, err := uc.PaymentService.GetFullPaymentCodeByCode(notifCtx, param.PaymentCode)
		if err != nil {
			log.Logger.Error("uc.PaymentService.GetFullPaymentCodeByCode for notification: ", err.Error())
			return
		}

		if dataPaymentCode == nil {
			log.Logger.Error("dataPaymentCode is nil for notification")
			return
		}

		// ambil data order by id
		dataOrder, err := uc.PaymentService.GetOrderByID(notifCtx, dataPaymentCode.Payment.OrderID)
		if err != nil {
			log.Logger.Error("uc.PaymentService.GetOrderByID for notification: ", err.Error())
			return
		}

		// ambil data user untuk nama & email
		resUser, err := uc.PaymentService.GetUser(notifCtx, dataOrder.UserID)
		if err != nil {
			log.Logger.Error("uc.PaymentService.GetUser for notification: ", err.Error())
			return
		}

		notifParam := models.NotificationOrderRequest{
			UserID:           dataOrder.UserID,
			Type:             "order_refund",
			TypeNotification: "email",
			Name:             resUser.Name,
			PaymentCode:      param.PaymentCode,
			Email:            resUser.Email,
			Service:          dataOrder.ServiceNameSnapshot,
			Amount:           fmt.Sprintf("Rp %.0f", dataPaymentCode.Payment.Amount), // sesuaikan field amount di model
		}

		err = uc.PaymentService.Notification(notifCtx, notifParam)
		if err != nil {
			log.Logger.Error("uc.PaymentService.Notification refund: ", err.Error())
		}
	}()

	return nil
}

func (uc *PaymentUsecase) GetMyRefund(ctx context.Context, userID int64, status string, page int, limit int) (*[]models.RejectedPayment, error) {
	data, err := uc.PaymentService.GetMyRefund(ctx, userID, status, page, limit)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (uc *PaymentUsecase) GetAllRefunds(ctx context.Context, status string, page int, limit int) (*[]models.RejectedPayment, error) {
	data, err := uc.PaymentService.GetAllRefunds(ctx, status, page, limit)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (uc *PaymentUsecase) GetRefundDetailByID(ctx context.Context, rejectID int64) (*models.RejectedPayment, error) {
	dataReject, err := uc.PaymentService.GetRejectPaymentByID(ctx, rejectID)
	if err != nil {

		log.Logger.Error("usecase.GetRejectPaymentByID")
		return nil, err
	}

	if dataReject == nil {
		return nil, utils.ErrRejectPaymentNotFound
	}

	data, err := uc.PaymentService.GetFullDetailRejectedPaymentByID(ctx, rejectID)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"id":    rejectID,
			"error": err.Error(),
		}).Error("usecase.GetFullDetailRejectedPaymentByID")
		return nil, err
	}

	return data, nil
}

func (uc *PaymentUsecase) SendNumberRekening(ctx context.Context, rejectID int64, param models.RequestRefund) error {
	data, err := uc.PaymentService.GetRejectPaymentByID(ctx, rejectID)
	if err != nil {
		log.Logger.Error("usecase.GetRejectPaymentByID")
		return err
	}

	if data == nil {
		return utils.ErrRejectPaymentNotFound
	}

	refund := &models.Refund{
		RejectedID:    rejectID,
		BankName:      param.BankName,
		AccountNumber: param.AccountNumber,
		AccountName:   param.AccountName,
		Status:        "requested",
	}

	// cek data di situ apakah sudah di crete
	dataRefund, err := uc.PaymentService.GetRefundByRejectedID(ctx, rejectID)
	if err != nil {
		log.Logger.Error("usecase.GetRefundByRejectedID")
		return err
	}

	if dataRefund != nil {
		return utils.ErrDataRefundIsFound
	}

	err = uc.PaymentService.CreateRefund(ctx, refund)
	if err != nil {
		log.Logger.Error("usecase.CreateRefund")
		return err
	}

	return nil
}

func (uc *PaymentUsecase) ProofRefund(ctx context.Context, fileProof *multipart.FileHeader, param models.RequestRefundProof, refundID int64) error {

	// cek refund id, ada tidak
	refund, err := uc.PaymentService.GetRefundByID(ctx, refundID)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"refund_id": refundID,
			"error":     err.Error(),
		}).Error("uc.PaymentService.GetRefundByID")
		return err
	}
	if refund == nil {
		return errors.New("refund tidak ditemukan")
	}

	if refund.Status != "requested" {
		return errors.New("refund sudah diproses sebelumnya")
	}

	// upload bukti ke storage
	url, err := uc.StorageService.UploadProofRefund(ctx, fileProof)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"refund_id": refundID,
			"error":     err.Error(),
		}).Error("uc.StorageService.UploadProofRefund")
		return err
	}

	// simpan ke db
	proof := &models.RefundProof{
		RefundID: refundID,
		FileURL:  url,
		Note:     param.Note,
	}

	err = uc.PaymentService.CreateRefundProof(ctx, proof)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"refund_id": refundID,
			"error":     err.Error(),
		}).Error("uc.PaymentService.CreateRefundProof")
		return err
	}

	// update status refund jadi transferred
	err = uc.PaymentService.UpdateRefundStatus(ctx, refundID, "transferred")
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"refund_id": refundID,
			"error":     err.Error(),
		}).Error("uc.PaymentService.UpdateRefundStatus")
		return err
	}

	return nil
}

// usecase user approve refund
func (uc *PaymentUsecase) ApproveRefund(ctx context.Context, refundID int64) error {

	data, err := uc.PaymentService.GetRefundByID(ctx, refundID)
	if err != nil {
		log.Logger.Error("usecase.GetRefundByID")
		return err
	}

	if data.Status != "transferred" {
		return utils.ErrStatusNotCompatible
	}

	err = uc.PaymentService.UpdateRefundStatus(ctx, refundID, "accepted")
	if err != nil {
		log.Logger.Error("usecase.UpdateRefundStatus")
		return err
	}

	return nil
}
