package collpay_go_sdk

import "time"

type Config struct {
	PublicKey string //Merchant's public key
	Env uint //Environment: sandbox or production
	Version string //Api version, ex:v1
	baseUrl string
}

type CommonData struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type ExchangeRate struct {
	Success bool
	Message string
	RateStr string `json:"rate"`
	Rate float64
}

type Transaction struct {
	ID string `json:"transaction_id"`
	Type string `json:"type"`
	OrderCurrency string `json:"order_currency"`
	OrderAmountStr string `json:"order_amount"`
	OrderAmount float64
	PaymentCurrency string `json:"payment_currency"`
	PaymentAmountStr string `json:"payment_amount"`
	PaymentAmount float64
	PayerName string `json:"payer_name"`
	PayerEmail string `json:"payer_email"`
	PayerPhone string `json:"payer_phone"`
	PayerAddress string `json:"payer_address"`
	CryptoAddress string `json:"crypto_address"`
	ExchangeRateStr string `json:"exchange_rate"`
	ExchangeRate float64
	ExpiryDate time.Time `json:"expiry_date"`
	HostedUrl string `json:"hosted_url"`
	IpnUrl string `json:"ipn_url"`
	SuccessUrl string `json:"success_url"`
	CancelUrl string `json:"cancel_url"`
	IpnSecret string `json:"ipn_secret"`
	Status string `json:"status"`
	Success bool
	Message string
	WebhookEvent string `json:"event"`
	Cart string `json:"cart"`
	WebhookData string `json:"webhook_data"`
}