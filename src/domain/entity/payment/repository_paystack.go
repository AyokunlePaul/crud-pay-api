package payment

import (
	crudPayError "github.com/AyokunlePaul/crud-pay-api/src/pkg/error_service"
	"github.com/rpip/paystack-go"
	"os"
)

var (
	client         *paystack.Client
	paystackEnvKey = "PAYMENT_GATEWAY_SECRET_KEY"
)

type paystackClient struct {
	errorService crudPayError.Service
}

func InitPaystack() {
	apiKey := os.Getenv(paystackEnvKey)
	client = paystack.NewClient(apiKey, nil)
}
