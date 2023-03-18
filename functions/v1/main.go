package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"strconv"
)

var artworkData Artworks = ReadJSON()

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	urlPath := request.Path
	var response string
	switch urlPath {
	case "/api/v1/recent/":
		pageNumberParam := request.QueryStringParameters["page"]
		pageNumber, err := strconv.Atoi(pageNumberParam)
		if err != nil {
			return &events.APIGatewayProxyResponse{
				StatusCode: 400,
				Headers:    map[string]string{"Content-Type": "application/json"},
				Body:       `{"error": "Invalid page number parameter"}`,
			}, nil
		}
		response = browseImages(pageNumber)
		// case "image":
		// 	getImageDetails(request.QueryStringParameters["id"])
	}
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       response,
	}, nil
}

// Should return map I think?
func browseImages(pageNumber int) string {
	jsonBytes, _ := json.Marshal(artworkData[pageNumber])
	return string(jsonBytes)
}

// // Should return map I think?
// func getImageDetails(imageId string) {

// }

func main() {
	// ReadJSON()
	lambda.Start(handler)
}
