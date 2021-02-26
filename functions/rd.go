package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
    req, err := http.NewRequest("GET", "https://www.google.com", nil)
    if err != nil {
        panic(err)
    }

    client := new(http.Client)
    response, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    fmt.Println(ioutil.ReadAll(response.Body))

    return &events.APIGatewayProxyResponse{
        StatusCode:        302,
        Headers:           map[string]string{Location: redirectUrl,'Cache-Control': 'no-cache',},
    }, nil
}

func main() {
    // Make the handler available for Remote Procedure Call by AWS Lambda
    lambda.Start(handler)
}