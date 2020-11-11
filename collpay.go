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
	if err != nil {
		return nil, err
	}

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

func makeTransactionRequestData(transaction *Transaction) (url.Values) {
	data := url.Values{}
	data.Set("type", "collpay")
	data.Set("payment_currency", transaction.PaymentCurrency)
	data.Set("order_currency", transaction.OrderCurrency)
	if transaction.OrderAmount > 0 {
		data.Set("order_amount", fmt.Sprint(transaction.OrderAmount))
	} else {
		data.Set("order_amount", transaction.OrderAmountStr)
	}
	data.Set("payer_name", transaction.PayerName)
	data.Set("payer_email", transaction.PayerEmail)
	data.Set("payer_phone", transaction.PayerPhone)
	data.Set("payer_address", transaction.PayerAddress)
	data.Set("ipn_url", transaction.IpnUrl)
	data.Set("ipn_secret", transaction.IpnSecret)
	data.Set("success_url", transaction.SuccessUrl)
	data.Set("cancel_url", transaction.CancelUrl)
	data.Set("cart", transaction.Cart)
	data.Set("webhook_data", transaction.WebhookData)
	return data
}

func CreateTransaction(tr *Transaction) (*Transaction, error){
	defer recoverPanic()
	data := makeTransactionRequestData(tr)

	req, err := http.NewRequest("POST", configData.baseUrl+"/transactions", strings.NewReader(data.Encode()))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	setHeaders(req)
	resBytes, err := doRequestAndGetResponse(req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

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

	//var tr Transaction
	err = json.Unmarshal(dataBytes, tr)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	tr.Success = respData.Success
	tr.Message = respData.Message

	if tr.Success {
		err = processTransactionFloatFields(tr)
		if err != nil {
			return nil, err
		}
	}

	return tr, nil
}

func GetTransaction(transactionId string) (*Transaction, error){
	defer recoverPanic()

	req, err := http.NewRequest("GET", configData.baseUrl+"/transactions/"+transactionId, nil)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	setHeaders(req)
	resBytes, err := doRequestAndGetResponse(req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

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

	var tr Transaction
	err = json.Unmarshal(dataBytes, &tr)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	tr.Success = respData.Success
	tr.Message = respData.Message

	if tr.Success {
		err = processTransactionFloatFields(&tr)
		if err != nil {
			return nil, err
		}
	}

	return &tr, nil
}

func processTransactionFloatFields(tr *Transaction) error {
	var err error
	tr.OrderAmount, err = strconv.ParseFloat(tr.OrderAmountStr, 64)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	tr.PaymentAmount, err = strconv.ParseFloat(tr.PaymentAmountStr, 64)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	tr.ExchangeRate, err = strconv.ParseFloat(tr.ExchangeRateStr, 64)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}