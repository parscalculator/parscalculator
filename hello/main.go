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
   firebaseurl := "https://calculator-a43b1.firebaseio.com/rest/saving-data/fireblog/posts.json"
    fmt.Printf("id: %s, link: ", request)



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
req2, _ := http.NewRequest("POST", firebaseurl, bytes.NewBuffer(payload))


req.Header.Set("Content-Type", "application/json")
req.Header.Set("X-API-KEY", "8b25c85f-0675-4c17-a28a-013aa8e5ecaa")
req.Header.Set("X-SANDBOX", "1")

res, _ := http.DefaultClient.Do(req)
res2, _ := http.DefaultClient.Do(req2)


defer res.Body.Close()
defer res2.Body.Close()

body, _ := ioutil.ReadAll(res.Body)
body2, _ := ioutil.ReadAll(res2.Body)


strbody:= string(body)
type nova struct {
    Id    string
    Link  string
}
strbody2:= string(body)
type nova2 struct {
    Name    string
}

var ret nova
var ret2 nova2

json.Unmarshal([]byte(strbody), &ret)
json.Unmarshal([]byte(strbody2), &ret2)

fmt.Printf("id: %s, link: %s", ret.Id, ret.Link)
fmt.Printf("name: %s", ret.Name)
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