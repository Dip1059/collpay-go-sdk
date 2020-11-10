package collpay_go_sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	ENV_PRODUCTION = 1;
	ENV_SANDBOX = 2
	V1 = "v1"
)

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
	OrderCurrency string `json:"order_currency"`
	OrderAmountStr string `json:"order_amount"`
	OrderAmount float64
	PaymentCurrency string `json:"payment_currency"`
	PaymentAmountStr string `json:"payment_amount"`
	PaymentAmount float64
	CryptoAddress string `json:"crypto_address"`
	ExchangeRateStr string `json:"exchange_rate"`
	ExchangeRate float64
	ExpiryDate time.Time `json:"expiry_date"`
	HostedUrl string `json:"hosted_url"`
	IpnSecret string `json:"ipn_secret"`
	Status string `json:"status"`
	Success bool
	Message string
	Event string `json:"event"`
	Cart string `json:"cart"`
	WebhookData string `json:"webhook_data"`
	StatusDate time.Time `json:"status_date"`
}

var (
	productionBaseUrl = "http://localhost:8000/api/";
	sandBoxBaseUrl = "https://collpay-dev.dev03.squaredbyte.com/api/";
	configData *Config
)

func recoverPanic() {
	if r := recover(); r != nil {
		var ok bool
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("pkg: %v", r)
		}
		log.Println(err.Error())
	}
}

func ConfigureEnv(config *Config) {
	configData = config
	switch configData.Env {
		case ENV_SANDBOX:
			configData.Env = ENV_SANDBOX
			configData.baseUrl = sandBoxBaseUrl
			break

		default: configData.Env = ENV_PRODUCTION
			configData.baseUrl = productionBaseUrl
			break
	}
	if configData.Version == "" {
		configData.Version = V1
	}
	configData.baseUrl += configData.Version
}

func setHeaders(req *http.Request) {
	req.Header.Add("x-auth", configData.PublicKey)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
}

func doRequestAndGetResponse(req *http.Request) ([]byte, error){
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	resBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return resBytes, nil
}

func GetExchangeRate(fromCurrency, toCurrency string) (*ExchangeRate, error){
	defer recoverPanic()
	data := url.Values{}
	data.Set("from", fromCurrency)
	data.Set("to", toCurrency)

	req, err := http.NewRequest("POST", configData.baseUrl+"/exchange-rate", strings.NewReader(data.Encode()))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	setHeaders(req)
	resBytes, err := doRequestAndGetResponse(req)

	var respData CommonData
	err = json.Unmarshal(resBytes, &respData)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	dataBytes, err := json.Marshal(respData.Data)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var exch ExchangeRate
	err = json.Unmarshal(dataBytes, &exch)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	exch.Success = respData.Success
	exch.Message = respData.Message

	if exch.RateStr != "" {
		exch.Rate, err = strconv.ParseFloat(exch.RateStr, 64)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}

	return &exch, nil
}

/*func CreateTransaction(fromCurrency, toCurrency string) (*ExchangeRate, error){
	defer recoverPanic()

	data := url.Values{}
	data.Set("from", fromCurrency)
	data.Set("to", toCurrency)

	req, err := http.NewRequest("POST", configData.baseUrl+"/exchange-rate", strings.NewReader(data.Encode()))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	setHeaders(req)
	resBytes, err := doRequestAndGetResponse(req)

	var respData ExchangeRate
	err = json.Unmarshal(resBytes, &respData)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	if respData.Data.Rate != "" {
		respData.Rate, err = strconv.ParseFloat(respData.Data.Rate, 64)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}

	return &respData, nil
}*/