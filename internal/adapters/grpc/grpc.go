package grpc

import (
	"context"
	"fmt"

	"github.com/3Davydov/ms-payment/internal/application/core/domain"
	"github.com/3Davydov/ms-proto/golang/payment"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) Create(ctx context.Context, request *payment.CreatePaymentRequest) (*payment.CreatePaymentResponse, error) {
	newPayment := domain.NewPayment(request.UserId, request.OrderId, request.TotalPrice)
	result, err := a.api.Charge(ctx, newPayment)
	fmt.Println("RESULT :")
	fmt.Println(result.ID)
	if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failed to charge. %v ", err)).Err()
	}
	return &payment.CreatePaymentResponse{PaymentId: result.ID}, nil
}
