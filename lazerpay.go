package lazerpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

var (
	initTrnx = "https://api.lazerpay.engineering/api/v1"
)

// LazerpayClient adds the necessary ``http.header``
//
//	lazerpay.LazerpayClient(publicKey, secretKey)
func LazerpayClient(publicKey, secretKey string) (http.Header, error) {
	return http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{secretKey},
		"x-api-key":     []string{publicKey},
	}, nil
}

// InitTransaction function helps initialize new transactions.
//
// The ``clientHeader`` is gotten from the LazerpayClient
//
// 		lazerpay.InitTransaction(clientHeader ,"4tytytreytrey65756u5u66", "1000", "Abdulfatai Suleiman", "staticdev20046@gmail.com", "USD", "DAI", true)
func InitTransaction(clientHeader http.Header, reference, amount, customerName, customerEmail, currencyType, coin string, acceptPartialPayment bool) (string, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	data := InitializeTransaction{
		Reference:            reference,
		Amount:               amount,
		CustomerName:         customerName,
		CustomerEmail:        customerEmail,
		Currency:             currencyType,
		Coin:                 coin,
		AcceptPartialPayment: acceptPartialPayment,
	}

	jsonData, _ := json.Marshal(data)

	client := &http.Client{}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/transaction/initialize", initTrnx), bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("error making a new request")
	}
	req.Header = clientHeader
	if err != nil {
		log.Println(err.Error())
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("error getting response")
	}
	defer res.Body.Close()

	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
	}

	return string(response), nil
}

// VerifyTransactions helps verify a payment.
//
// The ``clientHeader`` is gotten from the LazerpayClient
//
//		lazerpay.VerifyTransaction(clientHeader, "0xf2345527195C3bdc6C5f07576a3C860281926841")
func VerifyTransaction(clientHeader http.Header, reference string) (string, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/transaction/verify/%s", initTrnx, reference), nil)
	if err != nil {
		log.Println("error making a new request")
	}

	req.Header = clientHeader
	if err != nil {
		log.Println(err.Error())
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("error getting response")
	}
	defer res.Body.Close()

	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
	}

	return string(response), nil
}

// GetAcceptedCoins helps retrive all accepted coins.
//
// The ``clientHeader`` is gotten from the LazerpayClient
//
//		lazerpay.GetAcceptedCoins(clientHeader)
func GetAcceptedCoins(clientHeader http.Header) (string, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/coins", initTrnx), nil)
	if err != nil {
		log.Println("error making a new request")
	}

	req.Header = clientHeader
	if err != nil {
		log.Println(err.Error())
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("error getting response")
	}
	defer res.Body.Close()

	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	return string(response), nil
}

// Transfer helps send funds to another recipient.
//  amount must not be less than 1 and also it should not be empty.
//  recipient should not be empty, and also the recipient must be an Ethereum address.
//  coin must be a string, coin should not be empty.
//  blockchain must be a string, and blockchain must not be empty.
//
// The ``clientHeader`` is gotten from the LazerpayClient
//
//  lazerpay.Transfer(clientHeader,1, "0xF378c952d5266eF8e1783521a1395Fe40cDCe55B", "USDT", "Binance Smart Chain")
func Transfer(clientHeader http.Header, amount uint, recipient, coin, blockchain string) (string, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	data := TransferFunds{
		Amount:     amount,
		Recipient:  recipient,
		Coin:       coin,
		BlockChain: blockchain,
	}

	jsonData, _ := json.Marshal(data)

	client := &http.Client{}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/transfer", initTrnx), bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("error making a new request")
	}
	req.Header = clientHeader
	if err != nil {
		log.Println(err.Error())
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("error getting response")
	}
	defer res.Body.Close()

	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
	}

	return string(response), nil
}
