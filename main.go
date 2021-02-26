package main

import (
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "net/http"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
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