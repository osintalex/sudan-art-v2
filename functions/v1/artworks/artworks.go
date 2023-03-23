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

var artworkData []Artwork = marshalArtworksJSON()
var totalArtworks int = len(artworkData)

// Paginate the JSON back to client
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

// Convert the god awful string in `data.go` into a nice slice of structs
func marshalArtworksJSON() []Artwork {
	var artworkData []Artwork
	json.Unmarshal([]byte(ArtworksJSON), &artworkData)
	return artworkData
}
