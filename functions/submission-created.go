package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	url := "https://api.idpay.ir/v1.1/payment"

	data := map[string]string{
		"order_id": "101",
		"amount":   "10000",
		"name":     "قاسم رادمان",
		"phone":    "09382198592",
		"mail":     "my@site.com",
		"desc":     "توضیحات پرداخت کننده",
		"callback": "https://example.com/callback",
	}

	payload, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", "6a7f99eb-7c20-4412-a972-6dfb7cd253a4")
	req.Header.Set("X-SANDBOX", "1")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
