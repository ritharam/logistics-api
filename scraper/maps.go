package scraper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type MapsResponse struct {
	Routes []struct {
		Legs []struct {
			Duration struct {
				Text  string `json:"text"`
				Value int    `json:"value"`
			} `json:"duration"`
		} `json:"legs"`
	} `json:"routes"`
}

func GetTravelTime(origin, destination string) (int, error) {
	key := os.Getenv("MAPS_API_KEY")
	urlStr := fmt.Sprintf("https://maps.googleapis.com/maps/api/directions/json?origin=%s&destination=%s&key=%s",
		url.QueryEscape(origin), url.QueryEscape(destination), key)

	// Print shareable Google Maps URL
	fmt.Printf("📍 Google Maps Route: https://www.google.com/maps/dir/?api=1&origin=%s&destination=%s\n", origin, destination)

	r, err := http.Get(urlStr)
	if err != nil {
		return 0, err
	}
	defer r.Body.Close()

	var res MapsResponse
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		return 0, err
	}
	if len(res.Routes) == 0 {
		return 0, fmt.Errorf("No routes found")
	}
	return res.Routes[0].Legs[0].Duration.Value / 60, nil // travel time in minutes
}
