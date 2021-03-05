package main

import (
    "github.com/golang/protobuf/ptypes/timestamp"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "net/http"
    "encoding/json"
    "bytes"
    "io/ioutil"
    "fmt"
    "context"
)


type Form struct {
    Status    string `json:"status"`
    Track_id    number `json:"track_id"`
    Id string `json:"id"`
    Order_id string `json:"order_id"`
    Amount  string `json:"amount"`
    Card_no    string `json:"card_no"`
    Hashed_card_no    string `json:"hashed_card_no"`
    Date   string `json:"date"`
}

type Payload struct {
    Form Form `json:"data"`
}

type Body struct {
    Payload Payload `json:"payload"`
}


func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
    params := request.QueryStringParameters
  url := "https://api.idpay.ir/v1.1/payment/verify"
    var ret Body
    json.Unmarshal([]byte(request.Body), &ret)
    fmt.Printf("id: %s, link: %s", ret.Payload.Form.Id, ret.Payload.Form.Order_id)

data := map[string]string{
  "id":       params["id"],
  "order_id": params["order_id"],
}

payload, _ := json.Marshal(data)

req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))

req.Header.Set("Content-Type", "application/json")
req.Header.Set("X-API-KEY", "8b25c85f-0675-4c17-a28a-013aa8e5ecaa")
req.Header.Set("X-SANDBOX", "1")

res, _ := http.DefaultClient.Do(req)

defer res.Body.Close()
body, _ := ioutil.ReadAll(res.Body)

fmt.Println(string(body))
return &events.APIGatewayProxyResponse{
        StatusCode:        302,
        Headers:           map[string]string{"Location": "https://angry-albattani-4d4bde.netlify.app/success/" , "Content-Type": "text/plain"},
        MultiValueHeaders: http.Header{"Set-Cookie": {"Ding", "Ping"}},
        Body:              "Hello, World!",
        IsBase64Encoded:   false,
    }, nil
}

func main() {
    // Make the handler available for Remote Procedure Call by AWS Lambda
    lambda.Start(handler)
}