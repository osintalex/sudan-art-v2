package main

import (
  "fmt"
  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

  urlPath := request.Path
  return &events.APIGatewayProxyResponse{
    StatusCode:        200,
    Body:              fmt.Sprintf("URL path should be %s", urlPath),
  }, nil
}

func main() {
  lambda.Start(handler)
}
