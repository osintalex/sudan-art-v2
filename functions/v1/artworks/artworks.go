package artworks

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Artwork struct {
	Artist         string `json:"artist"`
	Tags           string `json:"tags"`
	Id             string `json:"id"`
	Thumbnail      string `json:"thumbnail"`
	High_res_image string `json:"high_res_image"`
	Date_uploaded  string `json:"date_uploaded"`
	Url            string `json:"url"`
}

type ArtworksWithImageData struct {
	Artwork
	Imagedata string `json:"image_data"`
}

func ReadArtworksJSON(filePath string) []Artwork {
	artworksList, error := ioutil.ReadFile(filePath)
	if error != nil {
		log.Fatalf("Error when opening file: %v", error.Error())
	}
	var artworkData []Artwork
	error = json.Unmarshal(artworksList, &artworkData)
	if error != nil {
		log.Fatalf("Error during unmarshal(): %v", error.Error())
	}
	return artworkData
}

func GenerateBase64Images(artworkData []Artwork) []ArtworksWithImageData {
	var artworksWithImageData []ArtworksWithImageData
	for _, artwork := range artworkData {
		imageFileName := artwork.High_res_image
		imageData, _, _ := readImage(imageFileName)
		artworkWithImageData := ArtworksWithImageData{Artwork: artwork, Imagedata: imageData}
		artworksWithImageData = append(artworksWithImageData, artworkWithImageData)
	}
	return artworksWithImageData
}

func readImage(imageFilePath string) (string, string, error) {
	imageFullPath := fmt.Sprintf("./sudan_art_images/%s", imageFilePath)
	bytes, error := ioutil.ReadFile(imageFullPath)
	if error != nil {
		log.Fatalf("Error reading in image with filepath %s. Error: %v",
			imageFilePath, error.Error())
		return "", "", error
	}
	var base64Encoding string
	mimeType := http.DetectContentType(bytes)

	switch mimeType {
	case "image/jpeg":
		base64Encoding = "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding = "data:image/png;base64,"
	}
	base64Encoding += base64.StdEncoding.EncodeToString(bytes)
	return base64Encoding, mimeType, nil
}
