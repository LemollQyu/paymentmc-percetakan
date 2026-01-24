package repository

import (
	"gorm.io/gorm"
)

type PaymentRepository struct {
	Database *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{
		Database: db,
	}
}
