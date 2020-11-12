package collpay_go_sdk

const (
	//env types
	ENV_PRODUCTION = 1
	ENV_SANDBOX = 2

	//versions
	V1 = "v1"

	//payment transaction statuses
	TRANSACTION_PROCESSING = "Processing"
	TRANSACTION_NOTIFIED = "Notified"
	TRANSACTION_EXPIRED = "Expired"
	TRANSACTION_CONFIRMED = "Confirmed"
	/*TRANSACTION_COMPLETED = "Completed"
	TRANSACTION_FAILED = "Failed"
	TRANSACTION_REJECTED = "Rejected"
	TRANSACTION_BLOCKED = "Blocked"
	TRANSACTION_REFUNDED = "Refunded"
	TRANSACTION_VOIDED = "Voided"*/

	//ipn or webhook events
	PAYMENT_EVENT = "payment"
)

var (
	productionBaseUrl = "http://localhost:8000/api/"
	sandBoxBaseUrl = "https://collpay-dev.dev03.squaredbyte.com/api/"
	configData *Config
)
