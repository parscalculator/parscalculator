package main

import (
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "net/http"
    "encoding/json"
    "bytes"
    "io/ioutil"
    "fmt"
    "context"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
    params := request.QueryStringParameters
   url := "https://api.idpay.ir/v1.1/payment"



data := map[string]string{
  "order_id": params["order_id"],
  "amount":   params["amount"],
  "name":     params["name"],
  "phone":    params["phone"],
  "mail":     params["mail"],
  "desc":     params["desc"],
  "callback": "https://angry-albattani-4d4bde.netlify.app/.netlify/functions/callback",
}

payload, _ := json.Marshal(data)

req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))

req.Header.Set("Content-Type", "application/json")
req.Header.Set("X-API-KEY", "8b25c85f-0675-4c17-a28a-013aa8e5ecaa")
req.Header.Set("X-SANDBOX", "1")

res, _ := http.DefaultClient.Do(req)

defer res.Body.Close()
body, _ := ioutil.ReadAll(res.Body)

strbody:= string(body)
type nova struct {
    Id    string
    Link  string
}
var ret nova
json.Unmarshal([]byte(strbody), &ret)
fmt.Printf("id: %s, link: %s", ret.Id, ret.Link)

return &events.APIGatewayProxyResponse{
        StatusCode:        302,
        Headers:           map[string]string{"Location": "https://google.com" , "Content-Type": "text/plain"},
        MultiValueHeaders: http.Header{"Set-Cookie": {"Ding", "Ping"}},
        Body:              "Hello, World!",
        IsBase64Encoded:   false,
    }, nil
}

func main() {
    // Make the handler available for Remote Procedure Call by AWS Lambda
    lambda.Start(handler)
}