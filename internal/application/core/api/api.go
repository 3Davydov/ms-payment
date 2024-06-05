package api

import "github.com/3Davydov/ms-payment/internal/application/core/domain"

type API interface {
	Charge(payment domain.Payment)
}
