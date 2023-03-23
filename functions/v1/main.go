package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"strconv"
	"sudan-art/artworks"
)

// Entrypoint for AWS Lamba function which is what runs this under the hood in production
func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	urlPath := request.Path
	var response string
	var statusCode int = 200
	switch urlPath {
	case "/api/v1/recent":
		pageNumberParam := request.QueryStringParameters["page"]
		pageNumber, error := strconv.Atoi(pageNumberParam)
		if error != nil || !(pageNumber >= 1 && pageNumber <= 100) {
			statusCode = 400
			response = `{"error": "Page number must be between 1 and 100"}`
			break
		}
		response, error = artworks.BrowseImages(pageNumber)
		if error != nil {
			statusCode = 500
			response = `{"error": "Internal error reading images :-("}`

		}
	default:
		statusCode = 404
		response = `{"error": "No such route exists"}`
	}
	return &events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       response,
	}, nil
}

func main() {
	lambda.Start(handler)
}
