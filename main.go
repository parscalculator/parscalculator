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
   url := "https://api.idpay.ir/v1.1/payment"

lc, ok := lambdacontext.FromContext(ctx)
cc := lc.ClientContext

data := map[string]string{
  "order_id": "101",
  "amount":   "10000",
  "name":     "cc.Client.AppTitle",
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
        Headers:           map[string]string{"Location": ret.Link , "Content-Type": "text/plain"},
        MultiValueHeaders: http.Header{"Set-Cookie": {"Ding", "Ping"}},
        Body:              "Hello, World!",
        IsBase64Encoded:   false,
    }, nil
}

func main() {
    // Make the handler available for Remote Procedure Call by AWS Lambda
    lambda.Start(handler)
}