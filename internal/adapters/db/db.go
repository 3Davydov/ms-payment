package db

import (
	"context"
	"fmt"

	"github.com/3Davydov/ms-payment/internal/application/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderID    int64
	TotalPrice float32
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, err := gorm.Open(postgres.Open(dataSourceUrl), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("db connection error: %v", err)
	}

	migrateErr := db.AutoMigrate(&Payment{})
	if migrateErr != nil {
		return nil, fmt.Errorf("db migration error: %v", migrateErr)
	}

	return &Adapter{db: db}, nil
}

func (a Adapter) Save(ctx context.Context, payment *domain.Payment) error {
	orderModel := Payment{
		CustomerID: payment.CustomerID,
		Status:     payment.Status,
		OrderID:    payment.OrderId,
		TotalPrice: payment.TotalPrice,
	}
	res := a.db.WithContext(ctx).Create(&orderModel)
	if res.Error == nil {
		payment.ID = int64(orderModel.ID)
	}
	return res.Error
}

func (a Adapter) Get(ctx context.Context, id string) (domain.Payment, error) {
	var paymentEntity Payment
	res := a.db.WithContext(ctx).First(&paymentEntity, id)
	payment := domain.Payment{
		ID:         int64(paymentEntity.ID),
		CustomerID: paymentEntity.CustomerID,
		Status:     paymentEntity.Status,
		OrderId:    paymentEntity.OrderID,
		TotalPrice: paymentEntity.TotalPrice,
		CreatedAt:  paymentEntity.CreatedAt.UnixNano(),
	}
	return payment, res.Error
}
