package artworks

import (
	"encoding/json"
	"log"
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

type BrowseResults struct {
	Results []Artwork `json:"results"`
	Next    bool      `json:"next"`
}

const offset = 5

var totalArtworks int = len(artworkData)
var artworkData []Artwork = readArtworksJSON("./sudan_art_database.json")

func BrowseImages(pageNumber int) (string, error) {
	startIndex := offset * pageNumber
	endIndex := startIndex + offset
	var hasMore bool = true
	if endIndex > totalArtworks {
		endIndex = totalArtworks
		hasMore = false
	}
	artworksSection := artworkData[startIndex:endIndex]
	browseResults := BrowseResults{artworksSection, hasMore}
	artworksBytes, error := json.Marshal(browseResults)
	if error != nil {
		log.Fatalf("Error unmarshalling images %v", error.Error())
		return "", error
	}
	return string(artworksBytes), nil
}

func readArtworksJSON(filePath string) []Artwork {
	var artworkData []Artwork
	error := json.Unmarshal([]byte(ArtworksJSON), &artworkData)
	if error != nil {
		log.Fatalf("Error unpacking JSON string %v", error.Error())
	}
	return artworkData
}
