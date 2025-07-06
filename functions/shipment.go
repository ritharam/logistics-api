package functions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/ritharam/logistics-api/database"
	"github.com/ritharam/logistics-api/models"
	"github.com/ritharam/logistics-api/scraper"
)

func ScoreOption(opt *models.ShippingOption, urgency string) {
	factor := 1.0
	if urgency == "high" {
		factor = 1.5
	}
	opt.Score = factor*float64(100-opt.EstimatedTime) - opt.Cost
}

func RecommendHandler(w http.ResponseWriter, r *http.Request) {
	var req models.Shipment
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	options, _ := HandleShipment(req, "")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(options)
}

func HandleShipment(req models.Shipment, email string) ([]models.ShippingOption, error) {
	eta, err := scraper.GetTravelTime(req.Origin, req.Destination)
	if err != nil {
		return nil, err
	}
	options := []models.ShippingOption{
		{"DHL", eta + 10, 150.0, 0},
		{"FedEx", eta + 5, 180.0, 0},
		{"UPS", eta + 15, 130.0, 0},
	}
	for i := range options {
		ScoreOption(&options[i], req.Urgency)
	}
	sort.Slice(options, func(i, j int) bool { return options[i].Score > options[j].Score })
	database.InsertShipment(req)

	if email != "" {
		link := fmt.Sprintf("https://www.google.com/maps/dir/?api=1&origin=%s&destination=%s", req.Origin, req.Destination)
		data := map[string]string{
			"origin":      req.Origin,
			"destination": req.Destination,
			"urgency":     req.Urgency,
			"link":        link,
		}
		err := SendMail(email, "Shipment Route Recommendation", data)
		if err != nil {
			log.Println("Failed to send mail:", err)
		} else {
			log.Println("Email sent to", email)
		}
		fmt.Println("📍 Google Maps Route:", link)
	}

	fmt.Printf("🚚 Best Option: %s | ETA: %d mins | Cost: ₹%.2f\n",
		options[0].Provider, options[0].EstimatedTime, options[0].Cost)

	return options, nil
}