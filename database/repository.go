package database

import "github.com/ritharam/logistics-api/models"

func SaveAndFetchShipment(s models.Shipment) ([]models.ShippingOption, error) {
	_, err := InsertShipment(s)
	if err != nil {
		return nil, err
	}
	return []models.ShippingOption{
		{"DHL", 12, 150.00, 0},
		{"FedEx", 9, 180.00, 0},
		{"UPS", 15, 130.00, 0},
	}, nil
}
