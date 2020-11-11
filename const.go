package collpay_go_sdk

const (
	ENV_PRODUCTION = 1
	ENV_SANDBOX = 2
	V1 = "v1"
	TRANSACTION_PROCESSING = "Processing"
	TRANSACTION_CONFIRMED = "Confirmed"
	TRANSACTION_NOTIFIED = "Notified"
	TRANSACTION_COMPLETED = "Completed"
	TRANSACTION_FAILED = "Failed"
	TRANSACTION_REJECTED = "Rejected"
	TRANSACTION_EXPIRED = "Expired"
	TRANSACTION_BLOCKED = "Blocked"
	TRANSACTION_REFUNDED = "Refunded"
	TRANSACTION_VOIDED = "Voided"
)

var (
	productionBaseUrl = "http://localhost:8000/api/"
	sandBoxBaseUrl = "https://collpay-dev.dev03.squaredbyte.com/api/"
	configData *Config
)
