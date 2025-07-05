package database

import (
	"log"
	"github.com/yourusername/logistics-api/models"
)

func InsertShipment(s models.Shipment) (int64, error) {
	res, err := DB.Exec(`INSERT INTO shipments (origin, destination, weight, urgency) VALUES (?, ?, ?, ?)`,
		s.Origin, s.Destination, s.Weight, s.Urgency)
	if err != nil {
		log.Println("Insert error:", err)
		return 0, err
	}
	return res.LastInsertId()
}

func GetShipments() ([]models.Shipment, error) {
	rows, err := DB.Query("SELECT * FROM shipments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shipments []models.Shipment
	for rows.Next() {
		var s models.Shipment
		_ = rows.Scan(&s.ID, &s.Origin, &s.Destination, &s.Weight, &s.Urgency, &s.CreatedAt)
		shipments = append(shipments, s)
	}
	return shipments, nil
}