package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/yourusername/logistics-api/database"
	"github.com/yourusername/logistics-api/functions"
	"github.com/yourusername/logistics-api/models"
	"github.com/yourusername/logistics-api/scraper"
)

func recommendHandler(w http.ResponseWriter, r *http.Request) {
	var req models.Shipment
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	eta, err := scraper.GetTravelTime(req.Origin, req.Destination)
	if err != nil {
		http.Error(w, "Maps API failed", http.StatusInternalServerError)
		return
	}

	options := []models.ShippingOption{
		{"DHL", eta + 10, 150.0, 0},
		{"FedEx", eta + 5, 180.0, 0},
		{"UPS", eta + 15, 130.0, 0},
	}

	for i := range options {
		functions.ScoreOption(&options[i], req.Urgency)
	}
	sort.Slice(options, func(i, j int) bool { return options[i].Score > options[j].Score })
	database.InsertShipment(req)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(options)
}

func main() {
	godotenv.Load()
	database.ConnectDB()
	r := mux.NewRouter()
	r.HandleFunc("/api/recommend", recommendHandler).Methods("POST")

	log.Println("🚀 Server running at :8080")
	http.ListenAndServe(":8080", r)
}
