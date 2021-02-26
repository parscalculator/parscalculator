package main

import (
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "net/http"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
    name := "World"
    
    return &events.APIGatewayProxyResponse{
        StatusCode:        302,
        Headers:           map[string]string{Location: "https://google.com",'Cache-Control': 'no-cache',},
        Body:       "Hi " + name,
    }, nil
}

func main() {
    // Make the handler available for Remote Procedure Call by AWS Lambda
    lambda.Start(handler)
}