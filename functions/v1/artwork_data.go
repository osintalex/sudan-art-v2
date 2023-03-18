package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Artworks []struct {
	Artist         string `json:"artist"`
	Tags           string `json:"tags"`
	Id             string `json:"id"`
	Thumbnail      string `json:"thumbnail"`
	High_res_image string `json:"high_res_image"`
	Date_uploaded  string `json:"date_uploaded"`
	Url            string `json:"url"`
}

func ReadJSON() Artworks {
	artworksList, error := ioutil.ReadFile("./sudan_art_database.json")
	if error != nil {
		log.Fatalf("Error when opening file: %v", error.Error())
	}
	log.Printf("All good in ze hood")
	var artworkData Artworks
	error = json.Unmarshal(artworksList, &artworkData)
	if error != nil {
		log.Fatalf("Error during Unmarshal(): %v", error.Error())
	}

	// Check it's working init
	log.Printf("The artist is: %s\n", artworkData[5].Artist)
	return artworkData
}
