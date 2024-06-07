package ports

import (
	"context"

	"github.com/3Davydov/ms-payment/internal/application/core/domain"
)

type DBPort interface {
	Save(ctx context.Context, payment *domain.Payment) error
}
