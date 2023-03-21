package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"strconv"
	"sudan-art/artworks"
)

const offset = 5

var artworkData []artworks.Artwork = artworks.ReadArtworksJSON("./sudan_art_database.json")
var totalArtworks int = len(artworkData)

type BrowseResults struct {
	Results []artworks.ArtworksWithImageData `json:"results"`
	Next    bool                             `json:"next"`
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	urlPath := request.Path
	var response string
	var statusCode int = 200
	var headers map[string]string = map[string]string{"Content-Type": "application/json"}
	switch urlPath {
	case "/api/v1/recent":
		pageNumberParam := request.QueryStringParameters["page"]
		pageNumber, error := strconv.Atoi(pageNumberParam)
		if error != nil || !(pageNumber >= 1 && pageNumber <= 100) {
			statusCode = 400
			response = `{"error": "Page number must be between 1 and 100"}`
			break
		}
		response = browseImages(pageNumber)
	// case "image":
	// 	getImageDetails(request.QueryStringParameters["id"])
	// case "/api/v1/thumbnail":
	// 	imageFilenameParam := request.QueryStringParameters["fileName"]
	// 	// validate me with regex gah that won't be fun
	// 	imageData, _, error := getImage(imageFilenameParam)
	// 	if error != nil {
	// 		statusCode = 500
	// 		response = `{"error": "Could not read requested image"}`
	// 		break
	// 	}
	// 	headers = map[string]string{"Content-Type": "text/plain"}

	// 	// b64 = true
	// 	response = imageData
	default:
		statusCode = 404
		response = `{"error": "No such route exists"}`
	}
	fmt.Println(statusCode)
	fmt.Println(headers)
	return &events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers:    headers,
		Body:       response,
	}, nil
}

func browseImages(pageNumber int) string {
	startIndex := offset * pageNumber
	endIndex := startIndex + offset
	var hasMore bool = true
	if endIndex > totalArtworks {
		endIndex = totalArtworks
		hasMore = false
	}
	artworksSection := artworkData[startIndex:endIndex]
	artworkTings := artworks.GenerateBase64Images(artworksSection)
	browseResults := BrowseResults{artworkTings, hasMore}
	artworksBytes, error := json.Marshal(browseResults)
	if error != nil {
		log.Fatalf("Something went hellabad so idk, panic for help %v", error.Error())
	}
	return string(artworksBytes)
}

// func getImageDetails(imageId string) {

// }

func main() {
	lambda.Start(handler)
}
