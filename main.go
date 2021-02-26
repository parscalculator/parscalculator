package main

import (
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "net/http"
    "encoding/json"
    "bytes"
    "io/ioutil"
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

}

func main() {
    // Make the handler available for Remote Procedure Call by AWS Lambda
    lambda.Start(handler)
}