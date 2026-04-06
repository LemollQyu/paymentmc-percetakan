package usecase

import (
	"context"
	"fmt"
	"mime/multipart"
	"paymentmc/infrastructure/log"
	"paymentmc/models"
	"paymentmc/utils"
	"time"

	"github.com/sirupsen/logrus"
)

func (uc *PaymentUsecase) Checkout(ctx context.Context, code string, param models.RequestPayment) (*models.ResponsePayment, error) {
	// ambil data order degan code
	resOrder, err := uc.PaymentService.GetOrderByCode(ctx, code)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"code":  code,
			"error": err.Error(),
		}).Error("uc.PaymentService.GetOrderByCode")

		return nil, err
	}

	if resOrder == nil {
		return nil, utils.OrderNotFound
	}

	// validasi status harus Created
	if resOrder.Status != utils.StatusOrderCreated {
		return nil, utils.OrderExpiredOrNotFound
	}

	// set status order jadi waiting_payment
	ok, err := uc.PaymentService.SetStatusOrder(ctx, code, utils.StatusOrderWaitingPayment)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"code":   code,
			"status": utils.StatusOrderWaitingPayment,
			"error":  err.Error(),
		}).Error("uc.PaymentService.SetStatusOrder")
		return nil, err
	}

	if !ok {
		return nil, err
	}

	reqPayment := &models.Payment{
		OrderID:       param.OrderID,
		UserID:        param.UserID,
		Amount:        param.Amount,
		PaymentMethod: param.PaymentMethod,
		Status:        utils.StatusPaymentPending,
	}

	// insert data paymentnya, status pending
	paymentId, err := uc.PaymentService.CreatePayment(ctx, reqPayment)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"param": reqPayment,
			"error": err.Error(),
		}).Error("uc.PaymentService.CreatePayment")
		return nil, err
	}

	// insert code payment, dan expirednya

	codePayment, err := utils.GeneratePaymentCode()
	if err != nil {
		return nil, err
	}

	reqPaymentCode := &models.PaymentCode{
		Code:      codePayment,
		PaymentID: paymentId,
		ExpiredAt: utils.Expire12Hour(),
	}
	idPaymentCode, err := uc.PaymentService.CreatePaymentCode(ctx, *reqPaymentCode)

	if err != nil {
		log.Logger.Error("uc.PaymentService.CreatePaymentCode")
		return nil, err
	}

	dataPaymentCode, err := uc.PaymentService.GetPaymentCodeByID(ctx, idPaymentCode)
	if err != nil {
		log.Logger.Error("uc.PaymentService.GetPaymentCodeByID")
		return nil, err
	}

	dataPaymentMethod, err := uc.PaymentService.GetPaymentMethod(ctx, param.PaymentMethod)
	if err != nil {
		log.Logger.Error("uc.PaymentSerivce.GetPaymentMethod")
		return nil, err
	}

	// create list_payment_waiting
	reqListWaitingPayment := &models.ListWaitingPayment{
		PaymentID:         dataPaymentCode.PaymentID,
		OrderID:           resOrder.ID,
		UserID:            resOrder.UserID,
		Amount:            param.Amount,
		OrderCode:         dataPaymentCode.Code,
		IconMethodPayment: dataPaymentMethod.UrlIcon,
		NumberPayment:     dataPaymentMethod.NumberPayment,
		CodeQris:          dataPaymentMethod.UrlCode,
		CheckoutAt:        time.Now(),
		ExpiredAt:         dataPaymentCode.ExpiredAt,
		CreatedAt:         time.Now(),
	}

	_, err = uc.PaymentService.InsertListWaitingPayment(ctx, reqListWaitingPayment)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"param": reqListWaitingPayment,
			"error": err.Error(),
		}).Error("InsertListWaitingPayment")
		return nil, err
	}

	// kirim notifikasi order checkout
	go func() {
		notifCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// ambil data user untuk nama & email
		resUser, err := uc.PaymentService.GetUser(notifCtx, resOrder.UserID)
		if err != nil {
			log.Logger.Error("uc.PaymentService.GetUser for notification: ", err.Error())
			return
		}

		notifParamUser := models.NotificationOrderRequest{
			UserID:           resOrder.UserID,
			Type:             "order_created",
			TypeNotification: "email",
			Name:             resUser.Name,
			OrderCode:        param.OrderCode,
			ExpiredAt:        dataPaymentCode.ExpiredAt.Format("02 Jan 2006 15:04"),
			Email:            resUser.Email,
			Service:          resOrder.ServiceNameSnapshot,
		}

		err = uc.PaymentService.Notification(notifCtx, notifParamUser)
		if err != nil {
			log.Logger.Error("uc.PaymentService.Notification checkout: ", err.Error())
		}

		notifParamAdmin := models.NotificationOrderRequest{
			UserID:           resOrder.UserID,
			Type:             "order_comes_in",
			TypeNotification: "email",
			Name:             resUser.Name,
			OrderCode:        param.OrderCode,
			ExpiredAt:        dataPaymentCode.ExpiredAt.Format("02 Jan 2006 15:04"),
			Email:            "nabilafotocopy09@gmail.com",
			Service:          resOrder.ServiceNameSnapshot,
			Amount:           utils.FormatRupiah(float64(resOrder.TotalPriceSnapshot)),
		}

		err = uc.PaymentService.Notification(notifCtx, notifParamAdmin)
		if err != nil {
			log.Logger.Error("uc.PaymentService.Notification checkout: ", err.Error())
		}
	}()

	return &models.ResponsePayment{
		OrderID:       resOrder.ID,
		Amount:        float64(resOrder.TotalPriceSnapshot),
		ExpiredAt:     dataPaymentCode.ExpiredAt,
		OrderCode:     dataPaymentCode.Code,
		CodeQris:      dataPaymentMethod.UrlCode,
		NumberPayment: dataPaymentMethod.NumberPayment,
		UserID:        resOrder.UserID,
		Service:       resOrder.ServiceNameSnapshot,
	}, nil
}

// broker expired payment and order
func (uc *PaymentUsecase) StartBrokerExpiredPaymentAndOrder(
	ctx context.Context,
) {

	log.Logger.Info("worker expired payment-order")
	ticker := time.NewTicker(5 * time.Minute)

	go func() {
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:

				log.Logger.Info("worker expired payment-order: mulai proses")

				// ambil data code payment
				paymentIDs, total, err := uc.PaymentService.GetExpiredPendingPaymentCodes(ctx)
				if err != nil {
					log.Logger.Error("uc.PaymentService.GetExpiredPendingPaymentCodes")
					continue
				}

				if total == 0 {
					log.Logger.Info("worker expired payment-order: tidak ada payment expired")
					continue
				}

				// ambil data order id nya
				orderIDs, err := uc.PaymentService.GetOrderIDsByPaymentIDs(ctx, paymentIDs)
				if err != nil {
					log.Logger.Error("uc.PaymentService.GetOrderIDsByPaymentIDs")
					continue
				}

				// ambil data order sebelum di-update (untuk notifikasi)
				var orders []*models.Order
				for _, orderID := range orderIDs {
					dataOrder, err := uc.PaymentService.GetOrderByID(ctx, orderID)
					if err != nil {
						log.Logger.WithFields(logrus.Fields{
							"order_id": orderID,
							"error":    err.Error(),
						}).Error("uc.PaymentService.GetOrderByID")
						continue
					}
					if dataOrder != nil {
						orders = append(orders, dataOrder)
					}
				}

				// set status order jadi Expired
				_, totalExpiredOrder, err := uc.PaymentService.UpdateExpiredOrder(ctx, orderIDs)
				if err != nil {
					log.Logger.WithFields(logrus.Fields{
						"ids":   orderIDs,
						"error": err.Error(),
					}).Error("uc.PaymentService.UpdateExpiredOrder")
					continue
				}

				if totalExpiredOrder == 0 {
					log.Logger.Info("worker expired payment-order: tidak ada order yang expired")
					continue
				}

				// set status payment jadi Expired
				err = uc.PaymentService.UpdateExpiredPayment(ctx, paymentIDs)
				if err != nil {
					log.Logger.WithFields(logrus.Fields{
						"ids":   paymentIDs,
						"error": err.Error(),
					}).Error("uc.PaymentService.UpdateExpiredPayment")
					continue
				}

				// kirim notifikasi per order
				for _, dataOrder := range orders {
					go func(dataOrder *models.Order) {
						notifCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
						defer cancel()

						resUser, err := uc.PaymentService.GetUser(notifCtx, dataOrder.UserID)
						if err != nil {
							log.Logger.Error("uc.PaymentService.GetUser for notification: ", err.Error())
							return
						}

						notifParam := models.NotificationOrderRequest{
							UserID:           dataOrder.UserID,
							Type:             "order_expired",
							TypeNotification: "email",
							Name:             resUser.Name,
							OrderCode:        dataOrder.OrderCode.Code,
							Email:            resUser.Email,
							Service:          dataOrder.ServiceNameSnapshot,
						}

						err = uc.PaymentService.Notification(notifCtx, notifParam)
						if err != nil {
							log.Logger.Error("uc.PaymentService.Notification order_expired: ", err.Error())
						}
					}(dataOrder)
				}

				log.Logger.WithFields(logrus.Fields{
					"total_payment": total,
					"total_order":   totalExpiredOrder,
				}).Info("worker expired payment-order: proses selesai")

			case <-ctx.Done():
				log.Logger.Info("worker expired payment-order: dihentikan (context done)")
				return
			}
		}
	}()
}

// uc cancelled pembayaran
func (uc *PaymentUsecase) CancelledPaymentAndOrder(ctx context.Context, code string) error {
	// ambil data payment code
	dataPaymentCode, err := uc.PaymentService.GetFullPaymentCodeByCode(ctx, code)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"payment_code": code,
			"error":        err.Error(),
		}).Error("uc.PaymentService.GetFullPaymentCodeByCode")
		return err
	}

	// check payment ada tidak
	if dataPaymentCode == nil {
		return utils.ErrPaymentCodeNotFound
	}

	// check payment status harus pending
	if dataPaymentCode.Payment.Status != utils.StatusPaymentPending {
		log.Logger.WithFields(logrus.Fields{
			"status_payment": dataPaymentCode.Payment.Status,
		}).Info("uc.PaymentService.GetFullPaymentCodeByCode")
		return utils.ErrStatusPaymentShouldPending
	}

	// set status payment jadi cancelled
	err = uc.PaymentService.UpdatePaymentStatus(ctx, dataPaymentCode.PaymentID, utils.StatusPaymentCancelled)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"payment_id": dataPaymentCode.PaymentID,
			"new_status": utils.StatusPaymentCancelled,
			"error":      err.Error(),
		}).Error("uc.PaymentService.UpdatePaymentStatus")
		return err
	}

	// ambil data order by id
	dataOrder, err := uc.PaymentService.GetOrderByID(ctx, dataPaymentCode.Payment.OrderID)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"order_id": dataPaymentCode.Payment.OrderID,
			"error":    err.Error(),
		}).Error("uc.PaymentService.GetOrderByID")

		return err
	}

	// validasi status order harus yang Waiting_payment
	if dataOrder.Status != utils.StatusOrderWaitingPayment {
		return utils.ErrStatusOrderShouldWaitingPayment
	}

	// set status order jadi cancelled
	_, err = uc.PaymentService.SetStatusOrder(ctx, dataOrder.OrderCode.Code, utils.StatusOrderCancelled)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"order_code": dataOrder.OrderCode.Code,
			"new_status": utils.StatusOrderCancelled,
			"error":      err.Error(),
		}).Error("uc.PaymentService.SetStatusOrder")

		return err
	}

	// kirim notifikasi order di cancelled
	go func() {
		notifCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// ambil data user untuk nama & email
		resUser, err := uc.PaymentService.GetUser(notifCtx, dataOrder.UserID)
		if err != nil {
			log.Logger.Error("uc.PaymentService.GetUser for notification: ", err.Error())
			return
		}

		notifParam := models.NotificationOrderRequest{
			UserID:           dataOrder.UserID,
			Type:             "order_cancelled",
			TypeNotification: "email",
			Name:             resUser.Name,
			OrderCode:        dataOrder.OrderCode.Code,
			Email:            resUser.Email,
			Service:          dataOrder.ServiceNameSnapshot,
		}

		err = uc.PaymentService.Notification(notifCtx, notifParam)
		if err != nil {
			log.Logger.Error("uc.PaymentService.Notification checkout: ", err.Error())
		}
	}()

	return nil
}

// uc bayar ordernya
func (uc *PaymentUsecase) ProofPayment(ctx context.Context, code string, fileProof *multipart.FileHeader, note string) (*models.ResponseUploadBuktiPembayaran, error) {
	// ambil data payment

	dataPaymentCode, err := uc.PaymentService.GetFullPaymentCodeByCode(ctx, code)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"payment_code": code,
			"error":        err.Error(),
		}).Error("uc.PaymentService.GetFullPaymentCodeByCode")
		return nil, err
	}

	// validasi payment ada

	if dataPaymentCode == nil {
		return nil, utils.ErrPaymentCodeNotFound
	}

	// validasi payment status harus Pending dan belum expired
	if dataPaymentCode.Payment.Status != utils.StatusPaymentPending {
		log.Logger.WithFields(logrus.Fields{
			"status_payment": dataPaymentCode.Payment.Status,
		}).Info("uc.PaymentService.GetFullPaymentCodeByCode")
		return nil, utils.ErrStatusPaymentShouldPending
	}

	// ambil data order
	dataOrder, err := uc.PaymentService.GetOrderByID(ctx, dataPaymentCode.Payment.OrderID)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"order_id": dataPaymentCode.Payment.OrderID,
			"error":    err.Error(),
		}).Error("uc.PaymentService.GetOrderByID")

		return nil, err
	}

	// validasi order harus Waiting_payment
	if dataOrder.Status != utils.StatusOrderWaitingPayment {
		return nil, utils.ErrStatusOrderShouldWaitingPayment
	}

	// kirim bukti pembyaran ke storage
	url, err := uc.StorageService.UploadProofPayment(ctx, fileProof)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"file":  fileProof,
			"error": err.Error(),
		}).Error("uc.PaymentService.GetOrderByID")
		return nil, err
	}

	reqProof := models.RequestBuktiPembayaran{
		PaymentID: dataPaymentCode.Payment.ID,
		Note:      note,
		ProofURL:  url,
	}

	// panggil service simpan ke db
	uploadAt, err := uc.PaymentService.PaymentProof(ctx, reqProof)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"url":   url,
			"param": reqProof,
			"error": err.Error(),
		}).Error("uc.PaymentService.CreatePaymentProof")
		return nil, err
	}

	// cek paymet sudah paid atau belum
	if dataPaymentCode.Payment.PaidAt != nil {
		return nil, utils.ErrPaymentPaid
	}

	// set payment di paid at nya dan status nya paid
	err = uc.PaymentService.UpdatedPaidAt(ctx, dataPaymentCode.PaymentID)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("uc.PaymentService.UpdatedPaidAt")
		return nil, err
	}

	if dataOrder.Status != utils.StatusOrderWaitingPayment {
		return nil, utils.ErrStatusOrderShouldWaitingPayment
	}

	// set status order paid
	_, err = uc.PaymentService.SetStatusOrder(ctx, dataOrder.OrderCode.Code, utils.StatusOrderPaid)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"order_code": dataOrder.OrderCode.Code,
			"new_status": utils.StatusOrderPaid,
			"error":      err.Error(),
		}).Error("uc.PaymentService.SetStatusOrder")

		return nil, err
	}

	// kirim notifikasi pembayaran dikonfirmasi
	go func() {
		notifCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		resUser, err := uc.PaymentService.GetUser(notifCtx, dataOrder.UserID)
		if err != nil {
			log.Logger.Error("uc.PaymentService.GetUser for notification: ", err.Error())
			return
		}

		amount := fmt.Sprintf("Rp %s", utils.FormatRupiah(dataPaymentCode.Payment.Amount))

		notifParam := models.NotificationOrderRequest{
			UserID:           dataOrder.UserID,
			Type:             "order_payment_success",
			TypeNotification: "email",
			Name:             resUser.Name,
			OrderCode:        dataOrder.OrderCode.Code,
			Email:            resUser.Email,
			Service:          dataOrder.ServiceNameSnapshot,
			Amount:           amount,
		}

		err = uc.PaymentService.Notification(notifCtx, notifParam)
		if err != nil {
			log.Logger.Error("uc.PaymentService.Notification payment confirmed: ", err.Error())
		}
	}()

	return uploadAt, nil

}

func (uc *PaymentUsecase) GetPayments(
	ctx context.Context,
	status string,
	page int,
	limit int,
) ([]*models.Payment, error) {

	// default
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	data, err := uc.PaymentService.GetPayments(
		ctx,
		status,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (uc *PaymentUsecase) ApprovePayment(ctx context.Context, codePayment string) error {
	// ambil data payment
	dataPaymentCode, err := uc.PaymentService.GetFullPaymentCodeByCode(ctx, codePayment)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"payment_code": codePayment,
			"error":        err.Error(),
		}).Error("uc.PaymentService.GetFullPaymentCodeByCode")
		return err
	}

	// validasi payment ada

	if dataPaymentCode == nil {
		return utils.ErrPaymentCodeNotFound
	}

	// validasi status payment sudah success
	if dataPaymentCode.Payment.Status != utils.StatusPaymentSuccess {
		return utils.ErrPaymentShouldSuccess
	}

	// ambil data order

	dataOrder, err := uc.PaymentService.GetOrderByID(ctx, dataPaymentCode.Payment.OrderID)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"order_id": dataPaymentCode.Payment.OrderID,
			"error":    err.Error(),
		}).Error("uc.PaymentService.GetOrderByID")

		return err
	}

	// cek ada tidak
	if dataOrder == nil {
		return utils.OrderNotFound
	}

	// cek status order

	if dataOrder.Status != utils.StatusOrderPaid {
		return utils.ErrOrderShouldPaid
	}

	// set approve payment
	err = uc.PaymentService.UpdateApproved(ctx, dataPaymentCode.Payment.ID)
	if err != nil {
		log.Logger.Error("uc.PaymentService.UpdateApproved")
		return err
	}

	// set status order on_progress
	_, err = uc.PaymentService.SetStatusOrder(ctx, dataOrder.OrderCode.Code, utils.StatusOrderOnProgress)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"order_code": dataOrder.OrderCode.Code,
			"new_status": utils.StatusOrderOnProgress,
			"error":      err.Error(),
		}).Error("uc.PaymentService.SetStatusOrder")

		return err
	}

	// kirim notidikasi
	go func() {
		notifCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// ambil data user untuk nama & email
		resUser, err := uc.PaymentService.GetUser(notifCtx, dataOrder.UserID)
		if err != nil {
			log.Logger.Error("uc.PaymentService.GetUser for notification: ", err.Error())
			return
		}

		notifParam := models.NotificationOrderRequest{
			UserID:           dataOrder.UserID,
			Type:             "order_confirmed",
			TypeNotification: "email",
			Name:             resUser.Name,
			OrderCode:        dataOrder.OrderCode.Code,
			Email:            resUser.Email,
			Service:          dataOrder.ServiceNameSnapshot,
		}

		err = uc.PaymentService.Notification(notifCtx, notifParam)
		if err != nil {
			log.Logger.Error("uc.PaymentService.Notification checkout: ", err.Error())
		}
	}()

	return nil
}

func (uc *PaymentUsecase) RejectPayment(ctx context.Context, codePayment string) error {
	// ambil data payment
	dataPaymentCode, err := uc.PaymentService.GetFullPaymentCodeByCode(ctx, codePayment)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"payment_code": codePayment,
			"error":        err.Error(),
		}).Error("uc.PaymentService.GetFullPaymentCodeByCode")
		return err
	}

	// validasi payment ada

	if dataPaymentCode == nil {
		return utils.ErrPaymentCodeNotFound
	}

	// validasi status payment sudah success
	if dataPaymentCode.Payment.Status != utils.StatusPaymentSuccess {
		return utils.ErrPaymentShouldSuccess
	}

	// ambil data order

	dataOrder, err := uc.PaymentService.GetOrderByID(ctx, dataPaymentCode.Payment.OrderID)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"order_id": dataPaymentCode.Payment.OrderID,
			"error":    err.Error(),
		}).Error("uc.PaymentService.GetOrderByID")

		return err
	}

	// cek ada tidak
	if dataOrder == nil {
		return utils.OrderNotFound
	}

	// cek status order

	if dataOrder.Status != utils.StatusOrderPaid {
		return utils.ErrOrderShouldPaid
	}

	// set reject/cancelled payment
	err = uc.PaymentService.UpdateRejected(ctx, dataPaymentCode.Payment.ID) //cancelled
	if err != nil {
		log.Logger.Error("uc.PaymentService.UpdateRejected")
		return err
	}

	// set status order on_progress
	_, err = uc.PaymentService.SetStatusOrder(ctx, dataOrder.OrderCode.Code, utils.StatusOrderCancelled)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"order_code": dataOrder.OrderCode.Code,
			"new_status": utils.StatusOrderCancelled,
			"error":      err.Error(),
		}).Error("uc.PaymentService.SetStatusOrder")

		return err
	}

	// kirim notifikasi order di cancelled
	go func() {
		notifCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// ambil data user untuk nama & email
		resUser, err := uc.PaymentService.GetUser(notifCtx, dataOrder.UserID)
		if err != nil {
			log.Logger.Error("uc.PaymentService.GetUser for notification: ", err.Error())
			return
		}

		notifParam := models.NotificationOrderRequest{
			UserID:           dataOrder.UserID,
			Type:             "order_cancelled",
			TypeNotification: "email",
			Name:             resUser.Name,
			OrderCode:        dataOrder.OrderCode.Code,
			Email:            resUser.Email,
			Service:          dataOrder.ServiceNameSnapshot,
		}

		err = uc.PaymentService.Notification(notifCtx, notifParam)
		if err != nil {
			log.Logger.Error("uc.PaymentService.Notification checkout: ", err.Error())
		}
	}()

	return nil
}

func (uc *PaymentUsecase) GetPaymentByOrderID(ctx context.Context, orderID int64) (*models.Payment, error) {

	data, err := uc.PaymentService.GetPaymentByOrderID(ctx, orderID)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"order_id": orderID,
			"error":    err.Error(),
		}).Error("uc.PaymentService.GetPaymentByOrderID")
		return nil, err
	}

	return data, nil

}

func (uc *PaymentUsecase) DeletePaymentByOrderID(ctx context.Context, orderID int64) error {
	err := uc.PaymentService.DeletePaymentByOrderID(ctx, orderID)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"order_id": orderID,
			"error":    err.Error(),
		}).Error("uc.PaymentService.GetPaymentByOrderID")
		return err
	}

	return nil
}
